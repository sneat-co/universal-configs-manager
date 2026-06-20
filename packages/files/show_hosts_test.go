package files

import (
	"testing"

	"ucm/packages/profiles"
)

func TestShowHosts(t *testing.T) {
	t.Run("nil_is_noop", func(t *testing.T) {
		if err := ShowHosts(nil); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("empty_entries_is_noop", func(t *testing.T) {
		if err := ShowHosts(&profiles.Hosts{}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("renders_entries_with_and_without_comment", func(t *testing.T) {
		hosts := &profiles.Hosts{Entries: []profiles.Host{
			{IP: "127.0.0.1", Name: "local.test", Comment: "dev"},
			{IP: "10.0.0.1", Name: "bare.test", Comment: "  "},
		}}
		if err := ShowHosts(hosts); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
