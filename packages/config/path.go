package config

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

var homedirDir = homedir.Dir

func getPath() (string, error) {
	homeDir, err := homedirDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".ucmconfig.yaml"), nil
}
