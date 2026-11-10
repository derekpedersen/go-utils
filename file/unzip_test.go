package file_test

import (
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"

	"github.com/derekpedersen/go-utils/file"
)

func TestUnZipGzCsvTo(t *testing.T) {
	directory := t.TempDir()
	input := filepath.Join(directory, "input.csv.gz")
	output := filepath.Join(directory, "nested", "output.csv")

	gzipWriterFile, err := os.Create(input)
	if err != nil {
		t.Fatal(err)
	}
	gzipWriter := gzip.NewWriter(gzipWriterFile)
	if _, err := gzipWriter.Write([]byte("a,b\n1,2\n")); err != nil {
		t.Fatal(err)
	}
	if err := gzipWriter.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gzipWriterFile.Close(); err != nil {
		t.Fatal(err)
	}

	if err := file.UnZipGzCsvTo(input, output); err != nil {
		t.Fatal(err)
	}
	contents, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "a,b\n1,2\n" {
		t.Fatalf("got %q", contents)
	}
}
