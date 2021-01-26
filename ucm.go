package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"ucm/commands"
)

func main() {
	if _, err := commands.Parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
