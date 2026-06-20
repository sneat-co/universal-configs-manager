package commands

import (
	"path/filepath"
	"testing"
)

func TestGetTargets(t *testing.T) {
	t.Run("resolves_names_to_home_dir_paths", func(t *testing.T) {
		c := profilesBaseCommand{Names: []string{"dev", "prod"}}
		targets, err := c.getTargets()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(targets) != 2 {
			t.Fatalf("expected 2 targets, got %d", len(targets))
		}
		for i, name := range []string{"dev", "prod"} {
			if targets[i].name != name {
				t.Errorf("target %d: expected name %q, got %q", i, name, targets[i].name)
			}
			if !filepath.IsAbs(targets[i].path) {
				t.Errorf("target %d: expected absolute path, got %q", i, targets[i].path)
			}
		}
	})

	t.Run("resolves_files_to_absolute_paths_with_empty_name", func(t *testing.T) {
		c := profilesBaseCommand{Files: []string{"some/relative.yaml"}}
		targets, err := c.getTargets()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(targets) != 1 {
			t.Fatalf("expected 1 target, got %d", len(targets))
		}
		if targets[0].name != "" {
			t.Errorf("expected empty name for file target, got %q", targets[0].name)
		}
		if !filepath.IsAbs(targets[0].path) {
			t.Errorf("expected absolute path, got %q", targets[0].path)
		}
	})

	t.Run("empty_when_nothing_provided", func(t *testing.T) {
		c := profilesBaseCommand{}
		targets, err := c.getTargets()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(targets) != 0 {
			t.Fatalf("expected no targets, got %d", len(targets))
		}
	})
}
