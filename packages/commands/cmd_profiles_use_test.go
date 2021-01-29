package commands

import (
	"strings"
	"testing"
)

func TestUseCommand_Execute(t *testing.T) {
	t.Run("returns_error", func(t *testing.T) {
		t.Run("if_no_arguments", func(t *testing.T) {
			sut := &useCommand{}
			err := sut.Execute([]string{})
			if err == nil {
				t.Fatal("expected to get error but got nil")
			}
			if !strings.Contains(err.Error(), "name") {
				t.Errorf("should mention 'name' parameter")
			}
			if !strings.Contains(err.Error(), "file") {
				t.Errorf("should mention 'file' parameter")
			}
		})
	})
}
