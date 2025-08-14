package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

//var ErrNotExists = os.ErrNotExist

func Read() (config Ucm, err error) {
	var path string

	if path, err = getPath(); err != nil {
		return
	}
	var file *os.File
	if file, err = os.Open(path); err != nil {
		if err == os.ErrNotExist {
			err = fmt.Errorf(".ucmconfig.yaml file not found in user's home directory: %w", err)
		}
		return
	}
	return decodeConfig(file)
}

func decodeConfig(r io.Reader) (config Ucm, err error) {
	if config.Raw, err = io.ReadAll(r); err != nil {
		return
	}
	err = yaml.Unmarshal(config.Raw, &config)
	return
}
