package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSymlinkPreflight(t *testing.T) {
	for _, denied := range []bool{false, true} {
		name := "supported"
		if denied {
			name = "denied"
		}
		t.Run(name, func(t *testing.T) {
			var probeDir string
			err := symlinkPreflight("windows", func(target, link string) error {
				probeDir = filepath.Dir(link)
				if target != probeDir {
					t.Fatalf("target = %q, want probe directory %q", target, probeDir)
				}
				if denied {
					return os.ErrPermission
				}
				return os.Symlink(target, link)
			})
			if denied {
				if !errors.Is(err, os.ErrPermission) {
					t.Fatalf("error = %v, want permission failure", err)
				}
				for _, hint := range []string{"Developer Mode", "elevated", "formulas", "skills"} {
					if !strings.Contains(err.Error(), hint) {
						t.Errorf("error %q is missing %q", err, hint)
					}
				}
			} else if err != nil {
				t.Fatalf("preflight: %v", err)
			}
			if probeDir == "" {
				t.Fatal("preflight did not attempt a symlink")
			}
			if _, err := os.Stat(probeDir); !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("probe directory was not removed: %v", err)
			}
		})
	}
}

func TestSymlinkPreflightDoesNothingOnUnix(t *testing.T) {
	for _, goos := range []string{"linux", "darwin"} {
		if err := symlinkPreflight(goos, func(string, string) error {
			t.Fatal("Unix preflight attempted a symlink")
			return os.ErrPermission
		}); err != nil {
			t.Fatalf("%s preflight: %v", goos, err)
		}
	}
}
