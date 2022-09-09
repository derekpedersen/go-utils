package file

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

func UnZipGzCsv(filename string) error {
	gzipFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	fileparts := strings.Split(filename, "/")
	outputfileuncompressed := fileparts[len(fileparts)-1]
	op := strings.Split(outputfileuncompressed, ".")
	outputfile := op[0] + ".csv"
	outfileWriter, err := os.Create(outputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()

	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
