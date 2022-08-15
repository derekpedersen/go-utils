package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func DownloadFile(key, datasource string) (string, error) {
	data, err := HttpRequest(datasource, http.MethodGet, nil)
	if err != nil {
		return *data, err
	}

	// save data points so we can re-run if needed for debugging or re-processing due to an encountered error.
	err = ioutil.WriteFile("./bin/downloads/"+key+"_"+time.Now().Format(time.ANSIC)+".json", []byte(*data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return *data, nil
}
