package limit_listener

import (
	"net"
	"sync"
)

// copy from
type limitListener struct {
	net.Listener
	sem       chan struct{}
	closeOnce sync.Once
	done      chan struct{}
}

func (l *limitListener) acquire() bool {
	select {
	case <-l.done:
		return false
	case l.sem <- struct{}{}:
		return true
	}
}

func (l *limitListener) release() {
	<-l.sem
}

func LimitListener(l net.Listener, n int) net.Listener {
	return &limitListener{
		Listener:  l,
		sem:       make(chan struct{}, n),
		closeOnce: sync.Once{},
		done:      make(chan struct{}),
	}
}

func (l *limitListener) Accept() (net.Conn, error) {
	acquired := l.acquire()

	// listener 关闭时, 直接返回 error
	c, err := l.Listener.Accept()
	if err != nil {
		if acquired {
			l.release()
		}
		return nil, err
	}
	return &limitListenerConn{
		Conn:        c,
		releaseOnce: sync.Once{},
		release:     l.release,
	}, nil
}

func (l *limitListener) Close() error {
	err := l.Listener.Close()
	l.closeOnce.Do(func() {
		close(l.done)
	})
	return err
}

type limitListenerConn struct {
	net.Conn
	releaseOnce sync.Once
	release     func()
}

func (l *limitListenerConn) Close() error {
	err := l.Conn.Close()
	l.releaseOnce.Do(func() {
		l.release()
	})
	return err
}
