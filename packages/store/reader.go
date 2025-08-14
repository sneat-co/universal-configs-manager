package store

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"ucm/packages/profiles"
)

// This is for unit tests
var osOpen = os.Open

// ReadProfileFromFile reads YAML configuration profile file
func ReadProfileFromFile(name, path string) (profile profiles.Profile, err error) {
	var file *os.File
	if file, err = osOpen(path); err != nil {
		if name != "" {
			err = fmt.Errorf("failed to open file for profle [%v]: %w", name, err)
		}
		return profile, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("failed to close profile file %v: %v", path, err))
		}
	}()
	return decodeProfile(file)
}

func decodeProfile(r io.Reader) (profile profiles.Profile, err error) {
	decoder := yaml.NewDecoder(r)
	return profile, decoder.Decode(&profile)
}
