package commands

import (
	"errors"
	"strings"
	"testing"

	"ucm/packages/profiles"
)

func TestUseCommand_Execute_ReadAndSwitch(t *testing.T) {
	origSwitch := nodeSwitchVersion
	origRegistry := nodeSetRegistry
	origHosts := filesSwitchHosts
	origAbs := filepathAbs
	origVerbose := options.Verbose
	t.Cleanup(func() {
		nodeSwitchVersion = origSwitch
		nodeSetRegistry = origRegistry
		filesSwitchHosts = origHosts
		filepathAbs = origAbs
		options.Verbose = origVerbose
	})

	t.Run("returns_error_when_getTargets_fails", func(t *testing.T) {
		t.Cleanup(func() { filepathAbs = origAbs })
		filepathAbs = func(string) (string, error) { return "", errors.New("abs boom") }
		sut := &useCommand{}
		sut.Files = []string{"any.yaml"}
		if err := sut.Execute(nil); err == nil {
			t.Fatal("expected error from getTargets")
		}
	})

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

	t.Run("fails_when_nodeSwitchVersion_fails", func(t *testing.T) {
		t.Cleanup(func() { nodeSwitchVersion = origSwitch })
		content := `nodejs:
  version: "18.0.0"
`
		path := writeProfileFile(t, content)
		nodeSwitchVersion = func(string) error { return errors.New("nvm failed") }
		sut := &useCommand{}
		sut.Files = []string{path}
		err := sut.Execute([]string{})
		if err == nil {
			t.Fatal("expected error from nodeSwitchVersion")
		}
		if !strings.Contains(err.Error(), "failed to switch NodeJS versoin") {
			t.Errorf("expected node switch failure, got: %v", err)
		}
	})

	t.Run("fails_when_nodeSetRegistry_fails", func(t *testing.T) {
		t.Cleanup(func() {
			nodeSwitchVersion = origSwitch
			nodeSetRegistry = origRegistry
		})
		content := `nodejs:
  version: "18.0.0"
  registry:
    id: npm
`
		path := writeProfileFile(t, content)
		nodeSwitchVersion = func(string) error { return nil }
		nodeSetRegistry = func(string) (string, error) { return "", errors.New("registry failed") }
		sut := &useCommand{}
		sut.Files = []string{path}
		err := sut.Execute([]string{})
		if err == nil {
			t.Fatal("expected error from nodeSetRegistry")
		}
		if !strings.Contains(err.Error(), "failed to set NodeJS registry") {
			t.Errorf("expected registry failure, got: %v", err)
		}
	})

	t.Run("fails_on_switch_hosts_for_valid_profile_without_nodejs", func(t *testing.T) {
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

	t.Run("succeeds_end_to_end_with_verbose", func(t *testing.T) {
		t.Cleanup(func() {
			nodeSwitchVersion = origSwitch
			nodeSetRegistry = origRegistry
			filesSwitchHosts = origHosts
			options.Verbose = origVerbose
		})
		content := `nodejs:
  version: "18.0.0"
  registry:
    id: npm
environment_variables:
  set:
    FOO: bar
`
		path := writeProfileFile(t, content)
		nodeSwitchVersion = func(string) error { return nil }
		nodeSetRegistry = func(string) (string, error) { return "ok", nil }
		filesSwitchHosts = func(*profiles.Hosts) error { return nil }
		options.Verbose = true

		sut := &useCommand{}
		sut.Files = []string{path}
		if err := sut.Execute([]string{"arg"}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestShowScriptToChangeEnvVars(t *testing.T) {
	t.Run("nil_is_noop", func(t *testing.T) {
		showScriptToChangeEnvVars(nil)
	})
	t.Run("empty_is_noop", func(t *testing.T) {
		showScriptToChangeEnvVars(&profiles.EnvVars{})
	})
	t.Run("renders_set_vars", func(t *testing.T) {
		ev := &profiles.EnvVars{Set: map[string]string{"FOO": "bar"}}
		showScriptToChangeEnvVars(ev)
	})
}
