package math_test

import (
	"errors"
	"testing"

	umath "github.com/derekpedersen/go-utils/math"
)

func TestRandomString(t *testing.T) {
	value, err := umath.RandomString(24)
	if err != nil {
		t.Fatal(err)
	}
	if len(value) != 24 {
		t.Fatalf("got length %d", len(value))
	}
}

func TestRandomRejectsNegativeLength(t *testing.T) {
	_, err := umath.RandomBytes(-1)
	if !errors.Is(err, umath.ErrInvalidLength) {
		t.Fatalf("got %v", err)
	}
}
