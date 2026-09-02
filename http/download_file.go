package http

import (
	"log"
	stdhttp "net/http"
	"os"
	"path/filepath"
	"time"
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
	filename := filepath.Join("./bin/downloads", key+"_"+time.Now().Format("20060102T150405.000000000")+".json")
	err = os.WriteFile(filename, []byte(*data), 0644)
	if err != nil {
		log.Printf("failed to save downloaded response: %v", err)
		return *data, err
	}

	return *data, nil
}
