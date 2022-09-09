package csv

import (
	"encoding/csv"
	"os"

	"github.com/sirupsen/logrus"
)

func Read(filepath string, seperator *rune) (records [][]string, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		logrus.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	if seperator != nil {
		csvReader.Comma = *seperator
	}
	csvReader.LazyQuotes = true
	return csvReader.ReadAll()
}
