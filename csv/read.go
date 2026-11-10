package csv

import (
	"encoding/csv"
	"io"
	"os"
)

type ReadOptions struct {
	Separator  rune
	LazyQuotes bool
}

func Read(filepath string, separator *rune) (records [][]string, err error) {
	options := &ReadOptions{LazyQuotes: true}
	if separator != nil {
		options.Separator = *separator
	}
	return ReadWithOptions(filepath, options)
}

func ReadWithOptions(filepath string, options *ReadOptions) (records [][]string, err error) {
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
	if options != nil {
		if options.Separator != 0 {
			csvReader.Comma = options.Separator
		}
		csvReader.LazyQuotes = options.LazyQuotes
	} else {
		csvReader.LazyQuotes = true
	}
	return csvReader.ReadAll()
}

func ReadWithHeader(filepath string, separator *rune) (header []string, rows [][]string, err error) {
	records, err := Read(filepath, separator)
	if err != nil {
		return nil, nil, err
	}
	if len(records) == 0 {
		return nil, nil, nil
	}
	return records[0], records[1:], nil
}

func ReadEach(filepath string, separator *rune, handler func([]string) error) (err error) {
	if handler == nil {
		return os.ErrInvalid
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
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
	for {
		record, readErr := csvReader.Read()
		if readErr == io.EOF {
			return nil
		}
		if readErr != nil {
			return readErr
		}
		if handlerErr := handler(record); handlerErr != nil {
			return handlerErr
		}
	}
}
