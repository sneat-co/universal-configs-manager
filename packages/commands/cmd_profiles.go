package commands

// profilesCommand implements "profiles" command
type profilesCommand struct {
}

// Execute implements "profiles" command
func (x *profilesCommand) Execute([]string) error {
	// At the moment is not possible to call this as it has subcommands
	// TODO: consider making subcommands optional and execute "list" subcommand by default
	return nil
}
