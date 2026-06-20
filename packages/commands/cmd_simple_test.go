package commands

import (
	"strings"
	"testing"
)

func TestModulesCommand_Execute(t *testing.T) {
	sut := &modulesCommand{}
	if err := sut.Execute([]string{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProfilesCommand_Execute(t *testing.T) {
	sut := &profilesCommand{}
	if err := sut.Execute([]string{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreateProfileCommand_Execute(t *testing.T) {
	t.Run("returns_error_if_no_arguments", func(t *testing.T) {
		sut := &createProfileCommand{}
		err := sut.Execute([]string{})
		if err == nil {
			t.Fatal("expected error but got nil")
		}
		if !strings.Contains(err.Error(), "name") || !strings.Contains(err.Error(), "file") {
			t.Errorf("error should mention 'name' and 'file', got: %v", err)
		}
	})
	t.Run("succeeds_with_name", func(t *testing.T) {
		sut := &createProfileCommand{Name: "dev"}
		if err := sut.Execute([]string{}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("succeeds_with_files", func(t *testing.T) {
		sut := &createProfileCommand{File: []string{"a.yaml"}}
		if err := sut.Execute([]string{}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
