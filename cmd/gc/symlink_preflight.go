package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func symlinkPreflight(goos string, symlink func(string, string) error) (err error) {
	if goos != "windows" {
		return nil
	}
	dir, err := os.MkdirTemp("", "gc-symlink-preflight-")
	if err != nil {
		return fmt.Errorf("creating symlink preflight directory: %w", err)
	}
	defer func() {
		if cleanupErr := os.RemoveAll(dir); cleanupErr != nil {
			err = errors.Join(err, fmt.Errorf("removing symlink preflight directory: %w", cleanupErr))
		}
	}()
	if err := symlink(dir, filepath.Join(dir, "link")); err != nil {
		return fmt.Errorf("symbolic links are required for formulas and skills on Windows; enable Developer Mode or run gc in an elevated shell: %w", err)
	}
	return nil
}
