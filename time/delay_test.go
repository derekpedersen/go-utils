package time_test

import (
	"testing"
	"time"

	ut "github.com/derekpedersen/go-utils/time"
)

func TestDelay(t *testing.T) {
	// Arrange
	starttime := time.Now()

	// Acts
	ut.Delay()
	endtime := time.Now()

	// Assert
	if endtime.Sub(starttime) < time.Duration(time.Second*1) {
		t.Errorf("There was no delay!")
	}
	if endtime.Sub(starttime) > time.Duration(time.Second*11) {
		t.Errorf("The delay was too long!")
	}
}
