package config

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

func getPath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".ucmconfig.yaml"), nil
}
