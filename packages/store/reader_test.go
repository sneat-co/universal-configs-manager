package store

import (
	"errors"
	"os"
	"testing"
)

func TestReadProfileFromFile(t *testing.T) {
	const fileName = "testFile"
	const filePath = "~/test_file"
	t.Run("returns_error", func(t *testing.T) {
		t.Run("if_failed_to_open_file", func(t *testing.T) {
			_, err := ReadProfileFromFile(fileName, filePath)
			if err == nil {
				t.Fatal("error expected but nil received")
			}
			if !errors.Is(err, os.ErrNotExist) {
				// TODO: (wording) What would be correct wording if returned error does not wraps os.ErrNotExist?
				t.Fatalf("expected to hold os.ErrNotExist got %T: %v", err, err)
			}
		})
	})
}
