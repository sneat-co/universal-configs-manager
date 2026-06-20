package commands

import (
	"fmt"
)

// profilesCommand implements "profiles" command
type listProfilesCommand struct {
	Location bool `short:"l" long:"location" description:"Show path or URL to profile"`
}

// Execute implements "profiles" command
func (x *listProfilesCommand) Execute([]string) error {
	ucmConfig, err := readConfig()
	if err != nil {
		return err
	}
	for _, profile := range ucmConfig.Profiles {
		fmt.Print(profile.Name)
		if x.Location {
			fmt.Print(" @ ", profile.DisplayPath())
		}
		fmt.Println("")
	}
	return nil
}
