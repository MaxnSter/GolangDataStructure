package try

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

type Runner interface {
	Do() error
	Catch(func(error) error) Runner
	Final(func(error) error) Runner
}

var _ Runner = (*runner)(nil)

type PanicErr struct {
	err error
}

func (p *PanicErr) Error() string {
	return p.err.Error()
}

type runner struct {
	try   func() error
	catch func(error) error
	final func(error) error
}

func Try(f func() error) Runner {
	return &runner{try: f}
}

func (r *runner) Do() (err error) {
	defer func() {
		if rc := recover(); rc != nil {
			buf := make([]byte, 2048)
			n := runtime.Stack(buf, false)
			fmt.Fprintf(os.Stderr, "runner panic: %v", string(buf[:n]))

			// type assert helper
			err = &PanicErr{err: errors.New(fmt.Sprintf("%v", rc))}
		}

		if err != nil && r.catch != nil {
			err = r.catch(err)
		}

		if r.final != nil {
			err = r.final(err)
		}
	}()

	err = r.try()
	return
}

func (r *runner) Catch(f func(error) error) Runner {
	r.catch = f
	return r
}

func (r *runner) Final(f func(error) error) Runner {
	r.final = f
	return r
}
