package store

import (
	"fmt"
	"path/filepath"
)

func GetNamedFile(name string) (file string, err error) {
	fileName := fmt.Sprintf(".ucm.%v.yaml", name)
	return filepath.Abs(fileName)
}
