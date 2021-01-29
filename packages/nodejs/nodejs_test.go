package nodejs

import "testing"

func TestGetRegistryByName(t *testing.T) {
	for k, v := range registries {
		registry := GetRegistryByName(k)
		if registry != v {
			t.Errorf("wrong registry returned for [%v]", k)
		}
	}
}
