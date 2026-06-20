package store

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadProfileFromFile_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "profile.yaml")
	content := `nodejs:
  version: "18.0.0"
git:
  user:
    name: John
`
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	profile, err := ReadProfileFromFile("dev", path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if profile.NodeJS == nil || profile.NodeJS.Version != "18.0.0" {
		t.Errorf("expected nodejs version 18.0.0, got %+v", profile.NodeJS)
	}
	if profile.Git == nil || profile.Git.User == nil || profile.Git.User.Name != "John" {
		t.Errorf("expected git user John, got %+v", profile.Git)
	}
}

func TestDecodeProfile(t *testing.T) {
	t.Run("parses_valid_yaml", func(t *testing.T) {
		profile, err := decodeProfile(strings.NewReader("nodejs:\n  version: \"20\"\n"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if profile.NodeJS == nil || profile.NodeJS.Version != "20" {
			t.Errorf("expected version 20, got %+v", profile.NodeJS)
		}
	})
	t.Run("returns_error_for_invalid_yaml", func(t *testing.T) {
		_, err := decodeProfile(strings.NewReader("\tbad: : yaml"))
		if err == nil {
			t.Fatal("expected error for invalid yaml")
		}
	})
}
