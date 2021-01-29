package commands

import (
	"errors"
	"fmt"
	"ucm/packages/files"
	"ucm/packages/nodejs"
	"ucm/packages/profiles"
	"ucm/packages/store"
)

// useCommand implements "use" command
type useCommand struct {
	profilesBaseCommand
	//Name  string   `short:"n" long:"name" description:"Name of configs set"`
	//Files []string `short:"f" long:"file" description:"Path to config file"`
}

// var useCommand useCommand

// Execute implements "use" command
func (x *useCommand) Execute(args []string) error {
	if len(x.Names) == 0 && len(x.Files) == 0 {
		return errors.New("either --name or --file parameter is required with 'use' command")
	}
	profileTargets, err := x.getTargets()
	if err != nil {
		return err
	}
	for _, profileTarget := range profileTargets {
		profile, err := store.ReadProfileFromFile(profileTarget.name, profileTarget.path)
		if err != nil {
			return fmt.Errorf("failed to read profile from file: %w", err)
		}
		if profile.NodeJS != nil {
			if err := nodejs.SwitchVersion(profile.NodeJS.Version); err != nil {
				return fmt.Errorf("failed to switch NodeJS versoin: %w", err)
			}
			if profile.NodeJS.Registry != nil {
				if _, err := nodejs.SetRegistry(profile.NodeJS.Registry.ID); err != nil {
					return fmt.Errorf("failed to set NodeJS registry: %w", err)
				}
			}
		}
		if err = showConfigSet(fmt.Sprintf("\nActivated named configs set: %v @ %v", profileTarget.name, profileTarget.path), profile); err != nil {
			return fmt.Errorf("failed to display profile: %w", err)
		}
		if err = files.SwitchHosts(profile.Hosts); err != nil {
			return fmt.Errorf("failed to switch hosts: %w", err)
		}
		if err = showScriptToChangeEnvVars(profile.EnvVars); err != nil {
			return fmt.Errorf("failed to show script for changing environment variables: %w", err)
		}
	}
	if options.Verbose {
		_, _ = fmt.Printf("Args: %+v\n", args)
	}
	return nil
}

func showScriptToChangeEnvVars(envVars *profiles.EnvVars) error {
	if envVars == nil || len(envVars.Set) == 0 && len(envVars.Remove) == 0 {
		return nil
	}
	_, _ = fmt.Println(`
===========================================================
Execute next in your shell to change environment variables:`)

	_, _ = fmt.Println("===== Beginning of shell script ====>")
	for key, value := range envVars.Set {
		_, _ = fmt.Printf("set %v=%v\n", key, value)
		//if err := os.Setenv(key, value); err != nil {
		//	return fmt.Errorf("failed to set environment variable [%v]: %w", key, err)
		//}
	}
	fmt.Println(`<==== End of shell script ==========
===========================================================`)
	return nil
}
