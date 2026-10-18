package commands

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"ucm/packages/profiles"
)

func writeProfileFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "profile.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("failed to write temp profile file: %v", err)
	}
	return path
}

func TestShowCommand_Execute(t *testing.T) {
	t.Run("returns_error_if_no_arguments", func(t *testing.T) {
		sut := &showCommand{}
		err := sut.Execute([]string{})
		if err == nil {
			t.Fatal("expected to get error but got nil")
		}
		if !strings.Contains(err.Error(), "name") || !strings.Contains(err.Error(), "file") {
			t.Errorf("error should mention 'name' and 'file' parameters, got: %v", err)
		}
	})

	t.Run("returns_error_if_file_does_not_exist", func(t *testing.T) {
		sut := &showCommand{}
		sut.Files = []string{"/no/such/path/missing.yaml"}
		if err := sut.Execute([]string{}); err == nil {
			t.Fatal("expected error for missing file")
		}
	})

	t.Run("succeeds_for_valid_profile_file", func(t *testing.T) {
		content := `environment_variables:
  set:
    FOO: bar
  remove:
    - BAZ
nodejs:
  version: "15.6.1"
  registry:
    id: npm
git:
  user:
    name: John Doe
    email: john@example.com
  author:
    name: Jane
  committer:
    email: c@example.com
hosts:
  entries:
    - ip: 127.0.0.1
      name: local.test
      comment: dev host
    - ip: 10.0.0.1
      name: bare.test
`
		path := writeProfileFile(t, content)
		sut := &showCommand{}
		sut.Files = []string{path}
		if err := sut.Execute([]string{}); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	})

	t.Run("succeeds_with_verbose", func(t *testing.T) {
		path := writeProfileFile(t, "git:\n  user:\n    name: Alex\n")
		orig := options.Verbose
		t.Cleanup(func() { options.Verbose = orig })
		options.Verbose = true
		sut := &showCommand{}
		sut.Files = []string{path}
		if err := sut.Execute([]string{"arg"}); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	})

	t.Run("returns_error_when_getTargets_fails", func(t *testing.T) {
		orig := filepathAbs
		t.Cleanup(func() { filepathAbs = orig })
		filepathAbs = func(string) (string, error) { return "", errors.New("abs boom") }
		sut := &showCommand{}
		sut.Files = []string{"any.yaml"}
		if err := sut.Execute(nil); err == nil {
			t.Fatal("expected error from getTargets")
		}
	})
}

func TestShowEnvVars(t *testing.T) {
	t.Run("nil_when_empty", func(t *testing.T) {
		showEnvVars(&profiles.EnvVars{})
	})
	t.Run("renders_set_and_remove", func(t *testing.T) {
		ev := &profiles.EnvVars{
			Remove: []string{"OLD"},
			Set:    map[string]string{"NEW": "1"},
		}
		showEnvVars(ev)
	})
}

func TestShowNode(t *testing.T) {
	t.Run("nil_node", func(t *testing.T) {
		showNode(nil)
	})
	t.Run("latest_version_and_registry_with_url", func(t *testing.T) {
		node := &profiles.NodeJS{
			Version:  "15.6.1",
			Registry: &profiles.NodeRegistry{ID: "npm", URL: "https://custom/"},
		}
		showNode(node)
	})
	t.Run("registry_without_url_resolves_by_name", func(t *testing.T) {
		node := &profiles.NodeJS{
			Version:  "12.0.0",
			Registry: &profiles.NodeRegistry{ID: "npm"},
		}
		showNode(node)
	})
}

func TestShowGit(t *testing.T) {
	t.Run("nil_git", func(t *testing.T) {
		showGit(nil)
	})
	t.Run("renders_all_users", func(t *testing.T) {
		git := &profiles.GitConfig{
			User:      &profiles.GitUser{Name: "U", Email: "u@e"},
			Author:    &profiles.GitUser{Name: "A"},
			Committer: &profiles.GitUser{Email: "c@e"},
		}
		showGit(git)
	})
}

func TestShowGitUser(t *testing.T) {
	t.Run("nil_user", func(t *testing.T) {
		showGitUser("User", nil)
	})
	t.Run("name_and_email", func(t *testing.T) {
		showGitUser("User", &profiles.GitUser{Name: "N", Email: "e@e"})
	})
}

func TestShowConfigSet(t *testing.T) {
	p := profiles.Profile{
		EnvVars: &profiles.EnvVars{Set: map[string]string{"K": "V"}},
		NodeJS:  &profiles.NodeJS{Version: "1.0.0"},
		Git:     &profiles.GitConfig{User: &profiles.GitUser{Name: "N"}},
		Hosts:   &profiles.Hosts{Entries: []profiles.Host{{IP: "1.1.1.1", Name: "h"}}},
	}
	showConfigSet("source", p)
}
