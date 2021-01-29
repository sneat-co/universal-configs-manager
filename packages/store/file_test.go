package store

import "testing"

func TestGetProfileFilePath(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		t.Run("if_empty_profile_name", func(t *testing.T) {
			defer func() {
				err := recover()
				if err == nil {
					t.Fatal("expected to panic")
				}
			}()
			GetProfileFilePath("", "~/")
		})
	})
	t.Run("succeed", func(t *testing.T) {
		path := GetProfileFilePath("test_profile_name", "~/")
		const expectedPath = "~/.ucm.test_profile_name.yaml"
		if path != expectedPath {
			t.Fatalf("expected %v got %v", expectedPath, path)
		}
	})
}
