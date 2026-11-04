//go:build windows

package main

// maybeRunControlDispatchSigtermChild is a no-op on Windows: the SIGTERM
// re-exec child it serves exists only for the !windows test that signals it.
func maybeRunControlDispatchSigtermChild() {}
