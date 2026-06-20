package commands

import (
	"strings"
	"testing"

	"ucm/packages/profiles"
)

func TestUseCommand_Execute_ReadAndSwitch(t *testing.T) {
	t.Run("returns_error_if_file_missing", func(t *testing.T) {
		sut := &useCommand{}
		sut.Files = []string{"/no/such/profile.yaml"}
		err := sut.Execute([]string{})
		if err == nil {
			t.Fatal("expected error for missing file")
		}
		if !strings.Contains(err.Error(), "failed to read profile") {
			t.Errorf("expected read failure, got: %v", err)
		}
	})

	t.Run("fails_on_switch_hosts_for_valid_profile_without_nodejs", func(t *testing.T) {
		// SwitchHosts is not implemented yet, so a valid profile without nodejs
		// progresses through reading + showConfigSet and then errors on hosts switch.
		content := `git:
  user:
    name: John
`
		path := writeProfileFile(t, content)
		sut := &useCommand{}
		sut.Files = []string{path}
		err := sut.Execute([]string{})
		if err == nil {
			t.Fatal("expected error from SwitchHosts")
		}
		if !strings.Contains(err.Error(), "switch hosts") {
			t.Errorf("expected hosts switch failure, got: %v", err)
		}
	})
}

func TestShowScriptToChangeEnvVars(t *testing.T) {
	t.Run("nil_is_noop", func(t *testing.T) {
		if err := showScriptToChangeEnvVars(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("empty_is_noop", func(t *testing.T) {
		if err := showScriptToChangeEnvVars(&profiles.EnvVars{}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("renders_set_vars", func(t *testing.T) {
		ev := &profiles.EnvVars{Set: map[string]string{"FOO": "bar"}}
		if err := showScriptToChangeEnvVars(ev); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
