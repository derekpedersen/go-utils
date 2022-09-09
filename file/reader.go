package file

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		logrus.Fatal(err)
	}

	return f, nil
}

func OpenCsvReader(f *os.File) *csv.Reader {
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.Comma = rune('|')
	return r
}

func LineCounter(r io.Reader) (int64, error) {
	buf := make([]byte, 32*1024)
	count := int64(0)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += int64(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
