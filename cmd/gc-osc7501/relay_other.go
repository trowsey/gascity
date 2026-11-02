//go:build !linux && !darwin

package main

import (
	"errors"
	"flag"
	"io"
	"runtime"
)

// relayVersion is 0: this build has no relay, so gc never launches a program
// under it.
const relayVersion = 0

// runRelay refuses: the relay needs a pseudo-terminal, which package relay
// provides only on Linux and macOS.
func runRelay(args []string, stderr io.Writer, _ []string) int {
	if _, err := parseRelayArgs(args, stderr); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			diag(stderr, "gc-osc7501 relay: %v\n%s", err, usage)
		}
		return exitUsage
	}
	diag(stderr, "gc-osc7501 relay: not supported on %s\n", runtime.GOOS)
	return exitStart
}
