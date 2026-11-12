package time_test

import (
	"context"
	"errors"
	"testing"
	stdtime "time"

	ut "github.com/derekpedersen/go-utils/time"
)

func TestSleepContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := ut.SleepContext(ctx, stdtime.Second); !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v", err)
	}
}

func TestSleepContextZeroDelay(t *testing.T) {
	if err := ut.SleepContext(context.Background(), 0); err != nil {
		t.Fatal(err)
	}
}

func TestSleepWithJitterRejectsInvalidRange(t *testing.T) {
	err := ut.SleepWithJitter(context.Background(), 2*stdtime.Second, stdtime.Second)
	if !errors.Is(err, ut.ErrInvalidDelayRange) {
		t.Fatalf("got %v", err)
	}
}

func TestSleepWithJitterExactDelay(t *testing.T) {
	start := stdtime.Now()
	if err := ut.SleepWithJitter(context.Background(), 0, 0); err != nil {
		t.Fatal(err)
	}
	if stdtime.Since(start) > 100*stdtime.Millisecond {
		t.Fatal("expected near-immediate return")
	}
}
