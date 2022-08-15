package bufio_test

import (
	"bufio"
	"strings"
	"testing"

	ub "github.com/derekpedersen/go-utils/bufio"
)

func TestReadline(t *testing.T) {
	// Arrange
	input := "Derek\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// Act
	line := ub.ReadLine(reader)

	// Assert
	if line == nil {
		t.Fatalf("No input was read!")
	}
	if *line != "Derek" {
		t.Errorf("Unable to read line correctly! Expected %v got %v", "Derek", line)
	}
}
