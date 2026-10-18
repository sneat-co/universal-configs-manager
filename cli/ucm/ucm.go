package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"ucm/packages/commands"
)

var (
	parserParse = commands.Parser.Parse
	osExit      = os.Exit
)

func main() {
	if _, err := parserParse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			osExit(0)
		} else {
			osExit(1)
		}
	}
}
