package singleflight

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroup_Do(t *testing.T) {
	g := new(Group)
	v, err := g.Do("do", func() (interface{}, error) {
		return "bar", nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "bar", v)
}

func TestGroup_Do2(t *testing.T) {
	g := new(Group)
	v, err := g.Do("do", func() (interface{}, error) {
		return "do", errors.New("some error")
	})

	assert.NotNil(t, err)
	assert.Equal(t, "do", v)
}

func TestGroup_Do3(t *testing.T) {
	var count int32
	g := new(Group)
	c := make(chan struct{})

	fn := func() (interface{}, error) {
		atomic.AddInt32(&count, 1)
		return <-c, nil
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			g.Do("do", fn)
			wg.Done()
		}()
	}

	time.Sleep(time.Millisecond * 100)
	c <- struct{}{}
	wg.Wait()

	assert.Equal(t, int32(1), atomic.LoadInt32(&count))
}
