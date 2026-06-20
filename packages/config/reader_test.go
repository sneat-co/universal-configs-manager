package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	orig := osOpen
	t.Cleanup(func() { osOpen = orig })

	t.Run("reads_and_decodes_config", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, ".ucmconfig.yaml")
		if err := os.WriteFile(path, []byte("active_profile: dev\n"), 0o600); err != nil {
			t.Fatalf("failed to write file: %v", err)
		}
		osOpen = func(string) (*os.File, error) { return os.Open(path) }
		cfg, err := Read()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.ActiveProfile != "dev" {
			t.Errorf("expected active profile 'dev', got %q", cfg.ActiveProfile)
		}
	})

	t.Run("wraps_not_found_error", func(t *testing.T) {
		osOpen = func(string) (*os.File, error) { return nil, os.ErrNotExist }
		_, err := Read()
		if err == nil {
			t.Fatal("expected error")
		}
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected wrapped os.ErrNotExist, got: %v", err)
		}
		if !strings.Contains(err.Error(), ".ucmconfig.yaml") {
			t.Errorf("expected message to mention config file, got: %v", err)
		}
	})
}

func TestDecodeConfig(t *testing.T) {
	t.Run("parses_valid_yaml", func(t *testing.T) {
		raw := `active_profile: dev
profiles:
  - name: dev
    path: /tmp/dev.yaml
  - name: prod
`
		cfg, err := decodeConfig(strings.NewReader(raw))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.ActiveProfile != "dev" {
			t.Errorf("expected active profile 'dev', got %q", cfg.ActiveProfile)
		}
		if len(cfg.Profiles) != 2 {
			t.Fatalf("expected 2 profiles, got %d", len(cfg.Profiles))
		}
		if string(cfg.Raw) != raw {
			t.Errorf("expected Raw to hold original bytes")
		}
	})

	t.Run("returns_error_for_invalid_yaml", func(t *testing.T) {
		_, err := decodeConfig(strings.NewReader("\tnot: : valid"))
		if err == nil {
			t.Fatal("expected error for invalid yaml")
		}
	})
}

func TestProfileSettings_DisplayPath(t *testing.T) {
	t.Run("returns_path_when_set", func(t *testing.T) {
		ps := ProfileSettings{Name: "dev", Path: "/custom/path.yaml"}
		if got := ps.DisplayPath(); got != "/custom/path.yaml" {
			t.Errorf("expected explicit path, got %q", got)
		}
	})
	t.Run("derives_from_name_when_no_path", func(t *testing.T) {
		ps := ProfileSettings{Name: "dev"}
		if got := ps.DisplayPath(); got != "~/.ucm.dev.yaml" {
			t.Errorf("expected derived path, got %q", got)
		}
	})
}

func TestGetPath(t *testing.T) {
	path, err := getPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasSuffix(path, ".ucmconfig.yaml") {
		t.Errorf("expected path to end with .ucmconfig.yaml, got %q", path)
	}
}
