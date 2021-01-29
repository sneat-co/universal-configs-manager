package profiles

import "testing"

func TestEnvVars_IsEmpty(t *testing.T) {

	t.Run("returns false", func(t *testing.T) {
		var sut *EnvVars
		//goland:noinspection GoNilness
		if !sut.IsEmpty() {
			t.Fatal("expected to return true for nil got false")
		}
		sut = new(EnvVars)
		if !sut.IsEmpty() {
			t.Fatal("expected to return true for empty struct got false")
		}
	})
}
