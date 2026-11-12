// Package retry provides context-aware operation retries.
package retry

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

type RetryPredicate func(error) bool

type JitterStrategy func(baseDelay time.Duration, attempt int) time.Duration

func NoJitter(baseDelay time.Duration, _ int) time.Duration {
	return baseDelay
}

func FullJitter(baseDelay time.Duration, _ int) time.Duration {
	if baseDelay <= 0 {
		return 0
	}
	return time.Duration(rand.Int63n(int64(baseDelay) + 1))
}

func DecorrelatedJitter(baseDelay time.Duration, attempt int) time.Duration {
	if baseDelay <= 0 {
		return 0
	}
	if attempt <= 0 {
		return baseDelay
	}
	maxDelay := time.Duration((1 << min(attempt, 10)) * int(baseDelay))
	if maxDelay < baseDelay {
		maxDelay = baseDelay
	}
	return baseDelay + time.Duration(rand.Int63n(int64(maxDelay-baseDelay)+1))
}

type Config struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
	Sleep        func(context.Context, time.Duration) error
	ShouldRetry  RetryPredicate
	Jitter       JitterStrategy
}

func Immediate(maxAttempts int) Config {
	return Config{MaxAttempts: maxAttempts}
}

func Exponential(maxAttempts int, initialDelay, maxDelay time.Duration) Config {
	return Config{
		MaxAttempts:  maxAttempts,
		InitialDelay: initialDelay,
		MaxDelay:     maxDelay,
		Multiplier:   2,
	}
}

func Do(ctx context.Context, config Config, operation func(context.Context) error) error {
	if operation == nil {
		return errors.New("retry: operation is nil")
	}
	if config.MaxAttempts <= 0 {
		config.MaxAttempts = 1
	}
	if config.Multiplier <= 0 {
		config.Multiplier = 2
	}
	if config.Sleep == nil {
		config.Sleep = sleep
	}
	if config.ShouldRetry == nil {
		config.ShouldRetry = func(error) bool { return true }
	}
	if config.Jitter == nil {
		config.Jitter = NoJitter
	}

	var err error
	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err = operation(ctx)
		if err == nil {
			return nil
		}
		if !config.ShouldRetry(err) {
			return err
		}
		if attempt == config.MaxAttempts-1 {
			return err
		}

		delay := config.InitialDelay
		for index := 0; index < attempt; index++ {
			delay = time.Duration(float64(delay) * config.Multiplier)
		}
		if config.MaxDelay > 0 && delay > config.MaxDelay {
			delay = config.MaxDelay
		}
		delay = config.Jitter(delay, attempt)
		if config.MaxDelay > 0 && delay > config.MaxDelay {
			delay = config.MaxDelay
		}
		if delay < 0 {
			delay = 0
		}
		if delay > 0 {
			if err := config.Sleep(ctx, delay); err != nil {
				return err
			}
		}
	}
	return err
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sleep(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
