//go:build !unix

package main

import (
	"fmt"
	"runtime"
	"syscall"
)

// raise reports that this platform cannot send a process a signal.
func raise(sig syscall.Signal) error {
	return fmt.Errorf("cannot raise %v on %s", sig, runtime.GOOS)
}
