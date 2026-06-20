package nodejs

import (
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func TestSwitchVersion(t *testing.T) {
	origExec := execCommand
	t.Cleanup(func() { execCommand = origExec })

	t.Run("noop_for_empty_version", func(t *testing.T) {
		if err := SwitchVersion(""); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("noop_when_already_active", func(t *testing.T) {
		execCommand = func(*exec.Cmd) (string, error) { return "v18.0.0\n", nil }
		if err := SwitchVersion("18.0.0"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("switches_when_different", func(t *testing.T) {
		calls := 0
		execCommand = func(*exec.Cmd) (string, error) {
			calls++
			if calls == 1 {
				return "v16.0.0\n", nil // current version
			}
			return "Now using node v18.0.0", nil // nvm use output
		}
		if err := SwitchVersion("18.0.0"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if calls != 2 {
			t.Errorf("expected node lookup + nvm use, got %d calls", calls)
		}
	})

	t.Run("returns_error_if_current_version_lookup_fails", func(t *testing.T) {
		execCommand = func(*exec.Cmd) (string, error) { return "", errors.New("node not found") }
		if err := SwitchVersion("18.0.0"); err == nil {
			t.Fatal("expected error when current version lookup fails")
		}
	})

	t.Run("returns_error_if_nvm_use_fails", func(t *testing.T) {
		calls := 0
		execCommand = func(*exec.Cmd) (string, error) {
			calls++
			if calls == 1 {
				return "v16.0.0\n", nil
			}
			return "", errors.New("nvm failed")
		}
		if err := SwitchVersion("18.0.0"); err == nil {
			t.Fatal("expected error when nvm use fails")
		}
	})
}

func TestGetCurrentNodeVersion(t *testing.T) {
	origExec := execCommand
	t.Cleanup(func() { execCommand = origExec })

	t.Run("trims_v_prefix_and_whitespace", func(t *testing.T) {
		execCommand = func(*exec.Cmd) (string, error) { return "  v20.1.0\n", nil }
		v, err := getCurrentNodeVersion()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if v != "20.1.0" {
			t.Errorf("expected 20.1.0, got %q", v)
		}
	})

	t.Run("returns_error", func(t *testing.T) {
		execCommand = func(*exec.Cmd) (string, error) { return "", errors.New("boom") }
		if _, err := getCurrentNodeVersion(); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestCallNpmConfigSet(t *testing.T) {
	origRun := runCmd
	t.Cleanup(func() { runCmd = origRun })

	t.Run("returns_output_on_success", func(t *testing.T) {
		runCmd = func(*exec.Cmd) error { return nil }
		out, err := callNpmConfigSet("registry", `"https://x/"`)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_ = out
	})

	t.Run("wraps_error_on_failure", func(t *testing.T) {
		runCmd = func(*exec.Cmd) error { return errors.New("npm boom") }
		if _, err := callNpmConfigSet("registry", "x"); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestSetRegistry(t *testing.T) {
	t.Run("noop_for_empty_registry", func(t *testing.T) {
		out, err := SetRegistry("")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out != "" {
			t.Errorf("expected empty output, got %q", out)
		}
	})
	t.Run("sets_known_registry", func(t *testing.T) {
		origRun := runCmd
		t.Cleanup(func() { runCmd = origRun })
		runCmd = func(*exec.Cmd) error { return nil }
		if _, err := SetRegistry("npm"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestSetNodeRegistryUsingNpmConfigSet(t *testing.T) {
	t.Run("noop_for_empty_registry", func(t *testing.T) {
		out, err := setNodeRegistryUsingNpmConfigSet("")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if out != "" {
			t.Errorf("expected empty output, got %q", out)
		}
	})
	t.Run("errors_for_unknown_registry_name", func(t *testing.T) {
		_, err := setNodeRegistryUsingNpmConfigSet("does-not-exist")
		if err == nil {
			t.Fatal("expected error for unknown registry name")
		}
		if !strings.Contains(err.Error(), "empty string") {
			t.Errorf("expected empty-url error, got: %v", err)
		}
	})
	t.Run("resolves_known_registry_name", func(t *testing.T) {
		origRun := runCmd
		t.Cleanup(func() { runCmd = origRun })
		runCmd = func(*exec.Cmd) error { return nil }
		if _, err := setNodeRegistryUsingNpmConfigSet("npm"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("accepts_explicit_https_url", func(t *testing.T) {
		origRun := runCmd
		t.Cleanup(func() { runCmd = origRun })
		runCmd = func(*exec.Cmd) error { return nil }
		if _, err := setNodeRegistryUsingNpmConfigSet("https://custom.example/"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
