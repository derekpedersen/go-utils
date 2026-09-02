package file

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnZipGzCsv(filename string) error {
	gzipFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer gzipFile.Close()

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	outputfile := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".csv"
	outfileWriter, err := os.Create(outputfile)
	if err != nil {
		return err
	}
	defer outfileWriter.Close()

	_, err = io.Copy(outfileWriter, gzipReader)
	return err
}
