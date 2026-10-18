package store

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadProfileFromFile(t *testing.T) {
	const fileName = "testFile"
	const filePath = "~/test_file"
	t.Run("returns_error", func(t *testing.T) {
		t.Run("if_failed_to_open_file", func(t *testing.T) {
			_, err := ReadProfileFromFile(fileName, filePath)
			if err == nil {
				t.Fatal("error expected but nil received")
			}
			if !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("expected to hold os.ErrNotExist got %T: %v", err, err)
			}
		})

		t.Run("if_failed_to_open_file_empty_name", func(t *testing.T) {
			_, err := ReadProfileFromFile("", filePath)
			if err == nil {
				t.Fatal("error expected but nil received")
			}
			if !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("expected to hold os.ErrNotExist got %T: %v", err, err)
			}
		})
	})

	t.Run("success", func(t *testing.T) {
		tmpDir := t.TempDir()
		profilePath := filepath.Join(tmpDir, "dev.yaml")
		content := `environment_variables:
  set:
    FOO: BAR
`
		if err := os.WriteFile(profilePath, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
		p, err := ReadProfileFromFile("dev", profilePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if p.EnvVars == nil || p.EnvVars.Set["FOO"] != "BAR" {
			t.Errorf("unexpected profile: %+v", p)
		}
	})

	t.Run("panic_on_close_error", func(t *testing.T) {
		tmpDir := t.TempDir()
		profilePath := filepath.Join(tmpDir, "dev.yaml")
		if err := os.WriteFile(profilePath, []byte("{}"), 0600); err != nil {
			t.Fatal(err)
		}
		oldOpen := osOpen
		defer func() { osOpen = oldOpen }()
		osOpen = func(name string) (*os.File, error) {
			f, err := os.Open(name)
			if err != nil {
				return nil, err
			}
			_ = f.Close()
			return f, nil
		}
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic on close error")
			}
		}()
		_, _ = ReadProfileFromFile("dev", profilePath)
	})
}
