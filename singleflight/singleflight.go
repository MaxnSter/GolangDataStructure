// Package singleflight provides a duplicate function call suppression
// mechanism.
package singleflight

import "sync"

// call is an in-flight or completed Do call
type call struct {
	//wg  sync.WaitGroup
	done chan struct{}
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

type DoFunc func() (interface{}, error)

// Do executes and returns the results of the give function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for then
// original to complete and receive the same results
func (g *Group) Do(key string, fn DoFunc) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	//方法在调用中
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()

		//等待调用完成
		<-c.done
		return c.val, c.err
	}

	c := newCall()
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	close(c.done)

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

func newCall() *call {
	return &call{
		done:make(chan struct{}),
	}
}
