//go:build linux || darwin

package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"os/signal"

	"github.com/gastownhall/gascity/internal/runtime/programstatus"
	"github.com/gastownhall/gascity/internal/runtime/programstatus/relay"
)

// relayVersion is the relay protocol this build implements; 0: no relay.
const relayVersion = programstatus.RelayVersion

// runRelay relays the process's own terminal to the program and ends the
// process as the program ended; it returns only on a usage or start failure.
func runRelay(args []string, stderr io.Writer, environ []string) int {
	argv, err := parseRelayArgs(args, stderr)
	if err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			diag(stderr, "gc-osc7501 relay: %v\n%s", err, usage)
		}
		return exitUsage
	}
	// Subscribe before the program starts, so no signal meant for it is
	// lost or takes the relay down instead.
	sigs := make(chan os.Signal, 16)
	signal.Notify(sigs, relay.Signals...)
	st, err := relay.Run(relay.Options{Argv: argv, Env: environ, Stdin: os.Stdin, Stdout: os.Stdout, Signals: sigs})
	if errors.Is(err, relay.ErrNotTerminal) {
		signal.Reset(relay.Signals...)
		err = relay.Exec(argv, environ)
	}
	if err != nil {
		diag(stderr, "gc-osc7501 relay: %v\n", err)
		return exitStart
	}
	relay.Exit(st)
	return exitFailure // unreachable: Exit ends the process
}
