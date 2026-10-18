package commands

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

// profilesCommand implements "list" command
type configCommand struct {
	Raw bool `long:"raw" description:"Show raw content of the ~/.ucmconfig.yaml file"`
}

// Execute implements "list" command
func (x *configCommand) Execute([]string) error {
	ucmConfig, err := readConfig()
	if err != nil {
		return err
	}
	if x.Raw {
		if _, err = stdout.Write(ucmConfig.Raw); err != nil {
			return err
		}
	} else {
		fmt.Printf("Active profile: %v\n", ucmConfig.ActiveProfile)
		fmt.Println("Named profiles:")
		for _, profile := range ucmConfig.Profiles {
			fmt.Printf("\t%v @ %v\n", profile.Name, profile.DisplayPath())
		}
	}
	return nil
}
