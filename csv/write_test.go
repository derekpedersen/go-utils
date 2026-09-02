package csv_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/derekpedersen/go-utils/csv"
)

func TestWriteFileCreatesParentAndWritesRecords(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "nested", "records.csv")
	if err := csv.WriteFile(filename, [][]string{{"name", "value"}, {"Ada", "42"}}); err != nil {
		t.Fatal(err)
	}
	contents, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "name,value\nAda,42\n" {
		t.Fatalf("got %q", contents)
	}
}
