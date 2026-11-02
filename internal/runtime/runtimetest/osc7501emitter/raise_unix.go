//go:build unix

package main

import (
	"os"
	"syscall"
)

// raise sends sig to this process.
func raise(sig syscall.Signal) error {
	return syscall.Kill(os.Getpid(), sig)
}
