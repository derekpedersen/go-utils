package csv

import (
	"encoding/csv"
	"os"
	filepathpkg "path/filepath"
)

type WriteOptions struct {
	Separator rune
}

func Write(key string, data [][]string) error {
	return WriteFile("./bin/output/"+key+"_data.csv", data)
}

func WriteFile(filepath string, data [][]string) (err error) {
	return WriteFileWithOptions(filepath, data, nil)
}

func WriteFileWithHeader(filepath string, header []string, rows [][]string) error {
	records := make([][]string, 0, len(rows)+1)
	records = append(records, header)
	records = append(records, rows...)
	return WriteFile(filepath, records)
}

func WriteFileWithOptions(filepath string, data [][]string, options *WriteOptions) (err error) {
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
	if options != nil && options.Separator != 0 {
		writer.Comma = options.Separator
	}
	writer.WriteAll(data)
	return writer.Error()
}
