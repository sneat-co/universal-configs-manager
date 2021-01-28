package store

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"ucm/packages/config"
	"ucm/packages/profiles"
)

// ReadFile reads YAML configuration file
func ReadFile(name string) (err error, configSet profiles.Profile) {
	var file *os.File

	if file, err = os.Open(name); err != nil {
		return err, configSet
	}
	return decodeProfile(file)
}

func decodeProfile(r io.Reader) (err error, configSet profiles.Profile) {
	decoder := yaml.NewDecoder(r)
	err = decoder.Decode(&configSet)
	return
}

func ReadConfig() (ucmConfig config.Ucm, err error) {
	return
}
