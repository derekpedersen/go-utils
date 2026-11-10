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

func TestDoStopsWhenPredicateRejectsRetry(t *testing.T) {
	attempts := 0
	err := retry.Do(context.Background(), retry.Config{
		MaxAttempts:  5,
		InitialDelay: time.Second,
		ShouldRetry: func(err error) bool {
			return err.Error() == "retryable"
		},
		Sleep: func(context.Context, time.Duration) error { return nil },
	}, func(context.Context) error {
		attempts++
		return errors.New("fatal")
	})
	if err == nil {
		t.Fatal("expected an error")
	}
	if attempts != 1 {
		t.Fatalf("got %d attempts, want 1", attempts)
	}
}

func TestDoUsesJitterStrategy(t *testing.T) {
	var delays []time.Duration
	err := retry.Do(context.Background(), retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Second,
		Jitter: func(base time.Duration, _ int) time.Duration {
			return base + 500*time.Millisecond
		},
		Sleep: func(_ context.Context, delay time.Duration) error {
			delays = append(delays, delay)
			return nil
		},
	}, func(context.Context) error {
		return errors.New("temporary")
	})
	if err == nil {
		t.Fatal("expected operation to fail after retries")
	}
	if len(delays) != 2 {
		t.Fatalf("unexpected delays length: %v", delays)
	}
	if delays[0] != 1500*time.Millisecond || delays[1] != 2500*time.Millisecond {
		t.Fatalf("unexpected jittered delays: %v", delays)
	}
}

func TestPresets(t *testing.T) {
	immediate := retry.Immediate(3)
	if immediate.MaxAttempts != 3 {
		t.Fatalf("unexpected immediate attempts: %d", immediate.MaxAttempts)
	}
	exponential := retry.Exponential(4, time.Second, 5*time.Second)
	if exponential.MaxAttempts != 4 || exponential.InitialDelay != time.Second || exponential.MaxDelay != 5*time.Second {
		t.Fatalf("unexpected exponential config: %+v", exponential)
	}
}
