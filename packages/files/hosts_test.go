package files

import "testing"

func TestSwitchHosts(t *testing.T) {
	t.Run("returns_error", func(t *testing.T) {
		err := SwitchHosts(nil)
		if err == nil {
			t.Fatal("expected to return an error")
		}
		if err.Error() != "not implemented yet" {
			t.Fatalf("expected to get 'not implemented yet', got %T: %v", err, err)
		}
	})
}
