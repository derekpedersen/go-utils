package csv

import (
	"encoding/csv"
	"os"
	filepathpkg "path/filepath"
)

func Write(key string, data [][]string) error {
	return WriteFile("./bin/output/"+key+"_data.csv", data)
}

func WriteFile(filepath string, data [][]string) (err error) {
	if err := os.MkdirAll(filepathpkg.Dir(filepath), 0755); err != nil {
		return err
	}
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := csvFile.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	writer := csv.NewWriter(csvFile)
	writer.WriteAll(data)
	return writer.Error()
}
