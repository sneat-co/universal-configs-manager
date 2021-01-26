package commands

import (
	"errors"
	"fmt"
)

// createProfileCommand implements "create" command
type createProfileCommand struct {
	Name string   `short:"n" long:"name" description:"Name of configs set"`
	File []string `short:"f" long:"file" description:"Path to config file"`
}

// Execute implements "create" command
func (x *createProfileCommand) Execute(args []string) error {
	if x.Name == "" && len(x.File) == 0 {
		return errors.New("either --name or --file parameter or both are required with 'create' command")
	}
	if x.Name != "" {
		_, _ = fmt.Printf("Configured named [%v]:\n", x.Name)
	}
	if len(x.File) > 0 {
		_, _ = fmt.Printf("Configured files: %+v:\n", x.File)
	}
	if options.Verbose {
		_, _ = fmt.Printf("Args: %+v\n", args)
	}
	return nil
}
