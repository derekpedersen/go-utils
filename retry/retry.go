// Package retry provides context-aware operation retries.
package retry

import (
	"context"
	"errors"
	"time"
)

type Config struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
	Sleep        func(context.Context, time.Duration) error
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

	var err error
	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err = operation(ctx)
		if err == nil {
			return nil
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
		if delay > 0 {
			if err := config.Sleep(ctx, delay); err != nil {
				return err
			}
		}
	}
	return err
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
