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

func TestExists(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "exists.txt")
	if err := os.WriteFile(filename, []byte("ok"), 0644); err != nil {
		t.Fatal(err)
	}
	exists, err := file.Exists(filename)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("expected file to exist")
	}

	missing, err := file.Exists(filepath.Join(t.TempDir(), "missing.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if missing {
		t.Fatal("expected file to be missing")
	}
}

func TestEnsureDir(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "a", "b", "c")
	if err := file.EnsureDir(directory); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(directory)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatal("expected directory")
	}
}
