package commands

import "fmt"

// modulesCommand implements "profiles" command
type modulesCommand struct {
}

// Execute implements "modulesCommand" command
func (x *modulesCommand) Execute(args []string) error {
	_, _ = fmt.Printf("Modules:\n")
	return nil
}
