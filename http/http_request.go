package http

import (
	"context"
	"fmt"
	"io"
	stdhttp "net/http"

	log "github.com/sirupsen/logrus"
)

type ResponseError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("http request failed with status %s", e.Status)
}

func IsSuccessStatus(statusCode int) bool {
	return statusCode >= stdhttp.StatusOK && statusCode < stdhttp.StatusMultipleChoices
}

func HttpRequest(url, method string, headers map[string][]string) (*string, error) {
	return HttpRequestWithClientAndContext(stdhttp.DefaultClient, context.Background(), url, method, headers, nil)
}

func HttpRequestWithContext(ctx context.Context, url, method string, headers map[string][]string) (*string, error) {
	return HttpRequestWithClientAndContext(stdhttp.DefaultClient, ctx, url, method, headers, nil)
}

func HttpRequestWithBody(url, method string, headers map[string][]string, body io.Reader) (*string, error) {
	return HttpRequestWithClientAndContext(stdhttp.DefaultClient, context.Background(), url, method, headers, body)
}

func HttpRequestWithClient(client *stdhttp.Client, url, method string, headers map[string][]string) (*string, error) {
	return HttpRequestWithClientAndContext(client, context.Background(), url, method, headers, nil)
}

func HttpRequestWithClientAndContext(client *stdhttp.Client, ctx context.Context, url, method string, headers map[string][]string, body io.Reader) (*string, error) {
	if client == nil {
		client = stdhttp.DefaultClient
	}
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := stdhttp.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		log.Errorf("Error creating request:\n %v", err)
		return nil, err
	}
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	log.WithField("request", req)
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("Error making request:\n %v", err)
		return nil, err
	}
	defer res.Body.Close()
	log.WithField("response", res)
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading res.Body:\n %v", err)
		return nil, err
	}
	s := string(responseBody)
	if !IsSuccessStatus(res.StatusCode) {
		return &s, &ResponseError{StatusCode: res.StatusCode, Status: res.Status, Body: s}
	}
	log.WithFields(log.Fields{
		"body": s,
	}).Debug("HttpRequest")
	return &s, nil
}
