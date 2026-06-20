package commands

import (
	"errors"
	"testing"

	"ucm/packages/config"
)

func TestConfigCommand_Execute(t *testing.T) {
	orig := readConfig
	t.Cleanup(func() { readConfig = orig })

	t.Run("returns_error_from_read", func(t *testing.T) {
		readConfig = func() (config.Ucm, error) { return config.Ucm{}, errors.New("boom") }
		sut := &configCommand{}
		if err := sut.Execute(nil); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("prints_raw_when_requested", func(t *testing.T) {
		readConfig = func() (config.Ucm, error) {
			return config.Ucm{Raw: []byte("active_profile: dev\n")}, nil
		}
		sut := &configCommand{Raw: true}
		if err := sut.Execute(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("prints_profiles_summary", func(t *testing.T) {
		readConfig = func() (config.Ucm, error) {
			return config.Ucm{
				ActiveProfile: "dev",
				Profiles: []config.ProfileSettings{
					{Name: "dev", Path: "/tmp/dev.yaml"},
					{Name: "prod"},
				},
			}, nil
		}
		sut := &configCommand{}
		if err := sut.Execute(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestListProfilesCommand_Execute(t *testing.T) {
	orig := readConfig
	t.Cleanup(func() { readConfig = orig })

	t.Run("returns_error_from_read", func(t *testing.T) {
		readConfig = func() (config.Ucm, error) { return config.Ucm{}, errors.New("boom") }
		sut := &listProfilesCommand{}
		if err := sut.Execute(nil); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("lists_names_with_location", func(t *testing.T) {
		readConfig = func() (config.Ucm, error) {
			return config.Ucm{Profiles: []config.ProfileSettings{
				{Name: "dev", Path: "/tmp/dev.yaml"},
				{Name: "prod"},
			}}, nil
		}
		sut := &listProfilesCommand{Location: true}
		if err := sut.Execute(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("lists_names_without_location", func(t *testing.T) {
		readConfig = func() (config.Ucm, error) {
			return config.Ucm{Profiles: []config.ProfileSettings{{Name: "dev"}}}, nil
		}
		sut := &listProfilesCommand{}
		if err := sut.Execute(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
