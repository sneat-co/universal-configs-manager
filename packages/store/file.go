package store

import (
	"fmt"
	"path"
)

func GetProfileFileName(profileName string) string {
	if profileName == "" {
		panic("not path for empty profile name")
	}
	return fmt.Sprintf(".ucm.%v.yaml", profileName)
}

func GetProfileFilePath(profileName, homeDir string) (filePath string) {
	return path.Join(homeDir, GetProfileFileName(profileName))
}
