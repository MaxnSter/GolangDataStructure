package retry

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var errCaller = errors.New("caller failed")

type caller struct {
	succeedTimes int
	callTimes int
}

func newCaller(t int) *caller{
	return &caller{succeedTimes:t}
}

func (c *caller) call() error {
	c.callTimes++
	if c.callTimes >= c.succeedTimes {
		return nil
	}
	return errCaller
}

func (c *caller) callWithDelay(d time.Duration) error {
	select {
	case <-time.After(d):
		return c.call()
	}
}


func TestRetrier_Run(t *testing.T) {
	err := New(ConstantBackoff(3, time.Millisecond * 500)).Run(newCaller(3).call)
	assert.Nil(t, err)

	err = New(ExponentialBackoff(3, time.Millisecond * 500)).Run(newCaller(4).call)
	assert.NotNil(t, err)
	assert.EqualValues(t, errCaller, err)
}

func TestRetrier_RunCtx(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	caller := newCaller(3)
	err := New(ExponentialBackoff(3, time.Millisecond * 500)).
		RunCtx(ctx, func(_ context.Context) error {
		return caller.callWithDelay(time.Second)
	})
	assert.NotNil(t, err)
	assert.EqualValues(t, context.DeadlineExceeded, err)

	ctx2, _ := context.WithTimeout(context.Background(), time.Second * 5)
	c2 := newCaller(3)
	err = New(ExponentialBackoff(3, time.Millisecond * 500)).
		RunCtx(ctx2, func(_ context.Context) error {
			return c2.callWithDelay(time.Second)
		})
	assert.Nil(t, err)
}
