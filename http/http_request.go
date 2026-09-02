package http

import (
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

func HttpRequest(url, method string, headers map[string][]string) (*string, error) {
	return HttpRequestWithClient(stdhttp.DefaultClient, url, method, headers)
}

func HttpRequestWithClient(client *stdhttp.Client, url, method string, headers map[string][]string) (*string, error) {
	if client == nil {
		client = stdhttp.DefaultClient
	}
	req, err := stdhttp.NewRequest(method, url, nil)
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading res.Body:\n %v", err)
		return nil, err
	}
	s := string(body)
	if res.StatusCode < stdhttp.StatusOK || res.StatusCode >= stdhttp.StatusMultipleChoices {
		return &s, &ResponseError{StatusCode: res.StatusCode, Status: res.Status, Body: s}
	}
	log.WithFields(log.Fields{
		"body": s,
	}).Debug("HttpRequest")
	return &s, nil
}
