package commands

import (
	"github.com/jessevdk/go-flags"
	"ucm/packages/config"
)

// readConfig is a seam for unit tests.
var readConfig = config.Read

// Options define & holds common additional arguments for the UCM application
type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Verbose output"`
}

var options Options

// Parser holds list of commands and provide a method to parse CLI arguments
var Parser = flags.NewParser(&options, flags.Default)

type commandAdder interface {
	AddCommand(name, shortDescription, longDescription string, data interface{}) (*flags.Command, error)
}

func mustAddCommand(adder commandAdder, name, short, long string, data interface{}) *flags.Command {
	cmd, err := adder.AddCommand(name, short, long, data)
	if err != nil {
		panic(err)
	}
	return cmd
}

func init() {
	mustAddCommand(Parser, "config",
		"Shows/manages UCM config",
		"Run `ucm config --raw`` to see raw content of the config file located at ~/.ucmconfig.yaml",
		&configCommand{})

	mustAddCommand(Parser, "modules",
		"Manages UCM modules",
		"List, add, removes UCM modules",
		&modulesCommand{})

	addProfileCommands(Parser)

	mustAddCommand(Parser, "use",
		"Switch to named configs set",
		"Shortcut for `ucm profiles use` - switches/updates configs as defined in a named configuration profile.",
		&useCommand{})

	mustAddCommand(Parser, "create",
		"Creates a named configs profile or writes it to a file",
		"Asks for all required values as configured per .ucmconfig.yaml and either creates a new named config or writes to a file.",
		&createProfileCommand{})
}

func addProfileCommands(adder commandAdder) {
	profilesCmd := mustAddCommand(adder, "profiles",
		"Manages config profiles",
		"Manages configuration profiles",
		&profilesCommand{})

	profilesCmd.SubcommandsOptional = false

	mustAddCommand(profilesCmd, "list",
		"Shows known profiles",
		"Outputs a list of known configuration profiles.",
		&listProfilesCommand{})

	mustAddCommand(profilesCmd, "create",
		"Creates new profile",
		"Creates either a named configuration profile or writes it to a file.",
		&createProfileCommand{})

	mustAddCommand(profilesCmd, "show",
		"Shows a configs set",
		"Shows values defined by a configurations set",
		&showCommand{})

	mustAddCommand(profilesCmd, "use",
		"Switch to named configs set",
		"Switches all configs as defined in a named configs set",
		&useCommand{})
}
