package csv

import (
	"encoding/csv"
	"os"
)

func Read(filepath string, separator *rune) (records [][]string, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := f.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	csvReader := csv.NewReader(f)
	if separator != nil {
		csvReader.Comma = *separator
	}
	csvReader.LazyQuotes = true
	return csvReader.ReadAll()
}
