package time

import (
	"context"
	"errors"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

var ErrInvalidDelayRange = errors.New("invalid delay range")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SleepContext(ctx context.Context, delay time.Duration) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if delay <= 0 {
		return nil
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func SleepWithJitter(ctx context.Context, min, max time.Duration) error {
	if min < 0 || max < 0 || max < min {
		return ErrInvalidDelayRange
	}
	delay := min
	if max > min {
		delta := max - min
		delay += time.Duration(rand.Int63n(int64(delta) + 1))
	}
	log.Debugf("Sleeping %s...", delay)
	return SleepContext(ctx, delay)
}

func Delay() {
	_ = SleepWithJitter(context.Background(), time.Second, 10*time.Second)
}
