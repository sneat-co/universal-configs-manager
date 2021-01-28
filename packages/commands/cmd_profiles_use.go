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
	Name string   `short:"n" long:"name" description:"Name of configs set"`
	File []string `short:"f" long:"file" description:"Path to config file"`
}

// var useCommand useCommand

// Execute implements "use" command
func (x *useCommand) Execute(args []string) error {
	if x.Name == "" && len(x.File) == 0 {
		return errors.New("either --name or --file parameter is required with 'use' command")
	}
	if x.Name != "" {
		file, err := store.GetNamedFile(x.Name)
		if err != nil {
			return err
		}
		err, configSet := store.ReadFile(file)
		if err != nil {
			return err
		}
		if configSet.NodeJS != nil {
			if err := nodejs.SwitchVersion(configSet.NodeJS.Version); err != nil {
				return err
			}
			if configSet.NodeJS.Registry != nil {
				if _, err := nodejs.SetRegistry(configSet.NodeJS.Registry.ID); err != nil {
					return err
				}
			}
		}
		if err = showConfigSet(fmt.Sprintf("\nActivated named configs set: %v @ %v", x.Name, file), configSet); err != nil {
			return err
		}
		if err = files.SwitchHosts(configSet.Hosts); err != nil {
			return err
		}
		if err = showScriptToChangeEnvVars(configSet.EnvVars); err != nil {
			return err
		}
	}
	if len(x.File) > 0 {
		_, _ = fmt.Printf("Configured files: %+v:\n", x.File)
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
