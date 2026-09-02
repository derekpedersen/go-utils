package digest_test

import (
	"strings"
	"testing"

	"github.com/derekpedersen/go-utils/digest"
)

func TestSHA256(t *testing.T) {
	got := digest.SHA256([]byte("hello"))
	want := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestSHA512Reader(t *testing.T) {
	got, err := digest.SHA512Reader(strings.NewReader("hello"))
	if err != nil {
		t.Fatal(err)
	}
	want := "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
