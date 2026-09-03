package retry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/derekpedersen/go-utils/retry"
)

func TestDoRetriesAndReturnsSuccess(t *testing.T) {
	attempts := 0
	var delays []time.Duration
	err := retry.Do(context.Background(), retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Second,
		Sleep: func(_ context.Context, delay time.Duration) error {
			delays = append(delays, delay)
			return nil
		},
	}, func(context.Context) error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary")
		}
		return nil
	})
	if err != nil || attempts != 3 {
		t.Fatalf("got err=%v attempts=%d", err, attempts)
	}
	if len(delays) != 2 || delays[0] != time.Second || delays[1] != 2*time.Second {
		t.Fatalf("unexpected delays: %v", delays)
	}
}

func TestDoStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := retry.Do(ctx, retry.Config{MaxAttempts: 2}, func(context.Context) error {
		t.Fatal("operation should not run")
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v, want context cancellation", err)
	}
}
