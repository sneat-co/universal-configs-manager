package commands

import "fmt"

// profilesCommand implements "profiles" command
type profilesCommand struct {
}

// Execute implements "profiles" command
func (x *profilesCommand) Execute(args []string) error {
	_, _ = fmt.Printf("Profiles:\n")
	return nil
}
