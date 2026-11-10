package http

import (
	"context"
	"fmt"
	"io"
	"log"
	stdhttp "net/http"
	"os"
	"path/filepath"
	"time"

	ufile "github.com/derekpedersen/go-utils/file"
)

func DownloadFile(key, datasource string) (string, error) {
	data, err := HttpRequest(datasource, stdhttp.MethodGet, nil)
	if err != nil {
		if data == nil {
			return "", err
		}
		return *data, err
	}

	if err := os.MkdirAll("./bin/downloads", 0755); err != nil {
		return *data, err
	}
	filename := DefaultDownloadPath("./bin/downloads", key)
	err = ufile.AtomicWrite(filename, []byte(*data), 0644)
	if err != nil {
		log.Printf("failed to save downloaded response: %v", err)
		return *data, err
	}

	return *data, nil
}

func DefaultDownloadPath(directory, key string) string {
	return filepath.Join(directory, key+"_"+time.Now().Format("20060102T150405.000000000")+".json")
}

func DownloadToFile(url, destination string, headers map[string][]string) error {
	return DownloadToFileWithClient(context.Background(), stdhttp.DefaultClient, url, destination, headers)
}

func DownloadToFileWithClient(ctx context.Context, client *stdhttp.Client, url, destination string, headers map[string][]string) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if client == nil {
		client = stdhttp.DefaultClient
	}

	req, err := stdhttp.NewRequestWithContext(ctx, stdhttp.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !IsSuccessStatus(res.StatusCode) {
		body, _ := io.ReadAll(res.Body)
		return &ResponseError{StatusCode: res.StatusCode, Status: res.Status, Body: string(body)}
	}

	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		return err
	}

	temporary, err := os.CreateTemp(filepath.Dir(destination), ".download-*")
	if err != nil {
		return err
	}
	temporaryName := temporary.Name()
	cleanup := true
	closed := false
	defer func() {
		if !closed {
			if closeErr := temporary.Close(); err == nil && closeErr != nil {
				err = closeErr
			}
		}
		if cleanup {
			_ = os.Remove(temporaryName)
		}
	}()

	if _, err = io.Copy(temporary, res.Body); err != nil {
		return err
	}
	if err := temporary.Sync(); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	closed = true
	if err := os.Rename(temporaryName, destination); err != nil {
		return fmt.Errorf("move downloaded file into place: %w", err)
	}
	cleanup = false
	return nil
}
