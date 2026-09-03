package file_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/derekpedersen/go-utils/file"
)

func TestAtomicWrite(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "nested", "data.txt")
	if err := file.AtomicWrite(filename, []byte("updated"), 0640); err != nil {
		t.Fatal(err)
	}
	contents, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "updated" {
		t.Fatalf("got %q", contents)
	}
}

func TestCopyFile(t *testing.T) {
	directory := t.TempDir()
	source := filepath.Join(directory, "source.txt")
	destination := filepath.Join(directory, "nested", "copy.txt")
	if err := os.WriteFile(source, []byte("copied"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := file.CopyFile(source, destination, 0640); err != nil {
		t.Fatal(err)
	}
	contents, err := os.ReadFile(destination)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "copied" {
		t.Fatalf("got %q", contents)
	}
}
