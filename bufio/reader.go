package bufio

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Reader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)
	return reader
}

func ReadLine(reader *bufio.Reader) *string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return nil
	}
	input = strings.TrimSuffix(input, "\n")
	log.WithField("input", input).Debug("ReadLine")
	return &input
}
