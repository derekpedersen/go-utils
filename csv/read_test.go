package csv_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/derekpedersen/go-utils/csv"
)

func TestReadWithHeader(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "records.csv")
	if err := os.WriteFile(filename, []byte("name,value\nAda,42\n"), 0644); err != nil {
		t.Fatal(err)
	}
	header, rows, err := csv.ReadWithHeader(filename, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(header) != 2 || header[0] != "name" || header[1] != "value" {
		t.Fatalf("unexpected header: %v", header)
	}
	if len(rows) != 1 || len(rows[0]) != 2 || rows[0][0] != "Ada" {
		t.Fatalf("unexpected rows: %v", rows)
	}
}

func TestReadWithOptions(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "records.psv")
	if err := os.WriteFile(filename, []byte("name|value\nAda|42\n"), 0644); err != nil {
		t.Fatal(err)
	}
	records, err := csv.ReadWithOptions(filename, &csv.ReadOptions{Separator: '|', LazyQuotes: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 2 || records[1][1] != "42" {
		t.Fatalf("unexpected records: %v", records)
	}
}

func TestReadEach(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "records.csv")
	if err := os.WriteFile(filename, []byte("name,value\nAda,42\n"), 0644); err != nil {
		t.Fatal(err)
	}
	count := 0
	err := csv.ReadEach(filename, nil, func(record []string) error {
		count += len(record)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Fatalf("unexpected count: %d", count)
	}
}

func TestReadEachRejectsNilHandler(t *testing.T) {
	err := csv.ReadEach(filepath.Join(t.TempDir(), "missing.csv"), nil, nil)
	if !errors.Is(err, os.ErrInvalid) {
		t.Fatalf("got %v", err)
	}
}
