package recoverF

import (
	"fmt"
	"os"
	"runtime"
)

func Decorate(f func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				fmt.Fprintf(os.Stderr, "runner panic: %v", string(buf[:n]))
			}
		}()

		f()
	}
}