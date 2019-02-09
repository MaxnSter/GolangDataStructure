package retry

import (
	"context"
	"fmt"
	"github.com/MaxnSter/GolangDataStructure/try"
	"time"
)

var defaultBackoff = ExponentialBackoff(1, time.Second)

type retrier struct {
	backoff []time.Duration
}

type retryErr struct {
	retryTime int
}

func (e retryErr) Error() string {
	return fmt.Sprintf("%d time retry", e.retryTime)
}

func (r *retrier) Run(f func() error) error {
	return r.RunCtx(context.Background(), func(_ context.Context) error {
		return f()
	})
}

func (r *retrier) RunCtx(ctx context.Context, f func(context.Context) error) error {
	callTimes := 0
	totalTimes := len(r.backoff)

	tryF := func() error {
		callTimes++
		return f(ctx)
	}
	final := func(err error) error {
		if err == nil {
			return nil
		}

		// last fail
		if callTimes >= totalTimes {
			return err
		}

		// retry
		select {
		case <-time.After(r.backoff[callTimes - 1]):
			return retryErr{retryTime: callTimes}
		case <-ctx.Done():
			// ctx done
			return ctx.Err()
		}
	}

	for {
		err := try.Try(tryF).Final(final).Do()
		if err == nil {
			return nil
		}
		if _, ok := err.(retryErr); ok {
			// ignore and retry
			continue
		}

		return err
	}
}

func New(backoff []time.Duration) *retrier {
	if len(backoff) <= 0 {
		backoff = defaultBackoff
	}

	return &retrier{backoff: backoff}
}
