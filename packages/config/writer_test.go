package config

import "testing"

func TestWrite(t *testing.T) {
	t.Run("returns_error", func(t *testing.T) {
		t.Run("if_no_path", func(t *testing.T) {
			err := Write(Ucm{})
			if err == nil {
				t.Fatal("expected to get error but got nothing")
			}
		})
		t.Run("not_implemented_yet", func(t *testing.T) {
			err := Write(Ucm{Path: "some_path"})
			if err == nil {
				t.Fatal("expected to get errNotImplemented but got nil")
			}
			if err != errNotImplemented {
				t.Fatalf("expected to get errNotImplemented but got %T: %v", err, err)
			}
		})
	})
}
