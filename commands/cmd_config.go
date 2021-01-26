package commands

import (
	"fmt"
	"os"
	"ucm/config"
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
		os.Stdout.Write(ucmConfig.Raw)
	} else {
		fmt.Printf("Active profile: %v\n", ucmConfig.ActiveProfile)
		fmt.Println("Named profiles:")
		for _, profile := range ucmConfig.Profiles {
			fmt.Printf("\t%v @ ~/.ucm.%v.yaml\n", profile.Name, profile.Name)
		}
	}
	return nil
}
