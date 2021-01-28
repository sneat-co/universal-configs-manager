package commands

import (
	"fmt"
	"os"
	"ucm/packages/config"
)

// profilesCommand implements "list" command
type configCommand struct {
	Raw bool `long:"raw" description:"Show raw content of the ~/.ucmconfig.yaml file"`
}

// Execute implements "list" command
func (x *configCommand) Execute([]string) error {
	ucmConfig, err := config.Read()
	if err != nil {
		return err
	}
	if x.Raw {
		if _, err = os.Stdout.Write(ucmConfig.Raw); err != nil {
			panic(err)
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
