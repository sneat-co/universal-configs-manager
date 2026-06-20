package commands

import (
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
}

func TestShowEnvVars(t *testing.T) {
	t.Run("nil_when_empty", func(t *testing.T) {
		if err := showEnvVars(&profiles.EnvVars{}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("renders_set_and_remove", func(t *testing.T) {
		ev := &profiles.EnvVars{
			Remove: []string{"OLD"},
			Set:    map[string]string{"NEW": "1"},
		}
		if err := showEnvVars(ev); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestShowNode(t *testing.T) {
	t.Run("nil_node", func(t *testing.T) {
		if err := showNode(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("latest_version_and_registry_with_url", func(t *testing.T) {
		node := &profiles.NodeJS{
			Version:  "15.6.1",
			Registry: &profiles.NodeRegistry{ID: "npm", URL: "https://custom/"},
		}
		if err := showNode(node); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("registry_without_url_resolves_by_name", func(t *testing.T) {
		node := &profiles.NodeJS{
			Version:  "12.0.0",
			Registry: &profiles.NodeRegistry{ID: "npm"},
		}
		if err := showNode(node); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestShowGit(t *testing.T) {
	t.Run("nil_git", func(t *testing.T) {
		if err := showGit(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("renders_all_users", func(t *testing.T) {
		git := &profiles.GitConfig{
			User:      &profiles.GitUser{Name: "U", Email: "u@e"},
			Author:    &profiles.GitUser{Name: "A"},
			Committer: &profiles.GitUser{Email: "c@e"},
		}
		if err := showGit(git); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestShowGitUser(t *testing.T) {
	t.Run("nil_user", func(t *testing.T) {
		if err := showGitUser("User", nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("name_and_email", func(t *testing.T) {
		if err := showGitUser("User", &profiles.GitUser{Name: "N", Email: "e@e"}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestShowConfigSet(t *testing.T) {
	p := profiles.Profile{
		EnvVars: &profiles.EnvVars{Set: map[string]string{"K": "V"}},
		NodeJS:  &profiles.NodeJS{Version: "1.0.0"},
		Git:     &profiles.GitConfig{User: &profiles.GitUser{Name: "N"}},
		Hosts:   &profiles.Hosts{Entries: []profiles.Host{{IP: "1.1.1.1", Name: "h"}}},
	}
	if err := showConfigSet("source", p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
