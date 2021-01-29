package commands

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"ucm/packages/store"
)

type profileTarget struct {
	name string
	path string
}

// profilesBaseCommand implements "use" command
type profilesBaseCommand struct {
	Names []string `short:"n" long:"name" description:"Name of configs set"`
	Files []string `short:"f" long:"file" description:"Path to config file"`
}

func (v profilesBaseCommand) getTargets() (profiles []profileTarget, err error) {
	if len(v.Names) > 0 {
		homeDir, err := homedir.Dir()
		if err != nil {
			// Should we fallback to current directory?
			return nil, fmt.Errorf("failed to determine home directory of user: %w", err)
		}
		for _, name := range v.Names {
			filePath := store.GetProfileFilePath(name, homeDir)
			profiles = append(profiles, profileTarget{name: name, path: filePath})
		}
	}
	for _, file := range v.Files {
		absPath, err := filepath.Abs(file)
		if err != nil {
			err = fmt.Errorf("failed to get absulute path for file %v: %w", file, err)
			return nil, err
		}
		profiles = append(profiles, profileTarget{name: "", path: absPath})
	}
	return
}
