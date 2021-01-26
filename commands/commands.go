package commands

import "github.com/jessevdk/go-flags"

// Options define & holds common additional arguments for the UCM application
type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Verbose output"`
}

var options Options

// Parser holds list of commands and provide a method to parse CLI arguments
var Parser = flags.NewParser(&options, flags.Default)

func init() {
	if _, err := Parser.AddCommand("config",
		"Shows/manages UCM config",
		"Run `ucm config --raw`` to see raw content of the config file located at ~/.ucmconfig.yaml",
		&configCommand{}); err != nil {
		panic(err)
	}

	if _, err := Parser.AddCommand("modules",
		"Manages UCM modules",
		"List, add, removes UCM modules",
		&modulesCommand{}); err != nil {
		panic(err)
	}

	addProfileCommands()

	if _, err := Parser.AddCommand("use",
		"Switch to named configs set",
		"Switches all configs as defined in a named configs set",
		&useCommand{}); err != nil {
		panic(err)
	}

	if _, err := Parser.AddCommand("create",
		"Creates a named configs profile or writes it to a file",
		"Asks for all required values as configured per .ucmconfig.yaml and either creates a new named config or writes to a file.",
		&createProfileCommand{}); err != nil {
		panic(err)
	}
}

func addProfileCommands() {
	profilesCmd, err := Parser.AddCommand("profiles",
		"Shows known profiles",
		"Outputs a list of known configuration profiles",
		&profilesCommand{})
	if err != nil {
		panic(err)
	}
	_, err = profilesCmd.AddCommand("create",
		"Creates new profile",
		"Creates either a named configuration profile or writes it to a file.",
		&createProfileCommand{})
	if err != nil {
		panic(err)
	}

	if _, err := profilesCmd.AddCommand("show",
		"Shows a configs set",
		"Shows values defined by a configurations set",
		&showCommand{}); err != nil {
		panic(err)
	}
}
