package byte_test

import (
	"testing"

	ubyte "github.com/derekpedersen/go-utils/byte"
)

type unsupported struct {
	F func()
}

func TestHashWithError(t *testing.T) {
	value, err := ubyte.HashWithError(map[string]int{"one": 1})
	if err != nil {
		t.Fatal(err)
	}
	if len(value) == 0 {
		t.Fatal("expected encoded bytes")
	}
}

func TestHashWithErrorRejectsUnsupportedType(t *testing.T) {
	_, err := ubyte.HashWithError(unsupported{F: func() {}})
	if err == nil {
		t.Fatal("expected gob encoding error")
	}
}

func TestHashReturnsNilForUnsupportedType(t *testing.T) {
	value := ubyte.Hash(unsupported{F: func() {}})
	if value != nil {
		t.Fatalf("expected nil bytes, got %v", value)
	}
}

func TestHashString(t *testing.T) {
	value, err := ubyte.HashString("abc")
	if err != nil {
		t.Fatal(err)
	}
	if value == "" {
		t.Fatal("expected hex string")
	}
}
