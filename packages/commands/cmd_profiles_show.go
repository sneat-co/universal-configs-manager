package commands

import (
	"errors"
	"fmt"
	"ucm/packages/files"
	"ucm/packages/nodejs"
	"ucm/packages/profiles"
	"ucm/packages/store"
)

type showCommand struct {
	profilesBaseCommand
	//Names  string   `short:"n" long:"name" description:"Names of configs set"`
	//Files []string `short:"f" long:"file" description:"Path to config file"`
}

// Execute implements "show" command
func (x *showCommand) Execute(args []string) error {
	if len(x.Names) == 0 && len(x.Files) == 0 {
		return errors.New("either --name or --file parameter is required with 'use' command")
	}
	profileTargets, err := x.getTargets()
	if err != nil {
		return err
	}
	for _, profileTarget := range profileTargets { // TODO: process in parallel and show in alphabetical order
		profile, err := store.ReadProfileFromFile("", profileTarget.path)
		if err != nil {
			return err
		}
		showConfigSet("Files: "+profileTarget.path, profile)
	}

	if options.Verbose {
		_, _ = fmt.Printf("Args: %+v\n", args)
	}
	return nil
}

func showConfigSet(source string, configSet profiles.Profile) { // TODO: implement colors
	_, _ = fmt.Println(source)
	showEnvVars(configSet.EnvVars)
	showNode(configSet.NodeJS)
	showGit(configSet.Git)
	_ = files.ShowHosts(configSet.Hosts)
}

func showEnvVars(envVars *profiles.EnvVars) {
	if envVars.IsEmpty() {
		return
	}
	_, _ = fmt.Println("\tEnvironment variables:")
	if len(envVars.Remove) > 0 {
		fmt.Println("\t\tRemove:")
		for _, name := range envVars.Remove { // TODO: output in alphabetical order
			fmt.Println("\t\t\t", name)
		}
	}
	if len(envVars.Set) > 0 {
		fmt.Println("\t\tSet:")
		for name, value := range envVars.Set { // TODO: output in alphabetical order
			fmt.Printf("\t\t\t%v: %v\n", name, value)
		}
	}
}

func showNode(node *profiles.NodeJS) {
	if node == nil {
		return
	}
	_, _ = fmt.Println("\tNodeJS:")
	if node.Version != "" {
		_, _ = fmt.Printf("\t\tVersion:    %v", node.Version)
		if node.Version == "15.6.1" { // TODO: get actual latest version
			_, _ = fmt.Print(" (latest)") // TODO: grey color
		}
		_, _ = fmt.Println("")
	}
	if node.Registry != nil {
		_, _ = fmt.Printf("\t\tRepository: %v", node.Registry.ID)
		// TODO: retrieve repository URL from `nrm` or node config and show or compare for equality
		url := node.Registry.URL
		if url == "" {
			url = nodejs.GetRegistryByName(node.Registry.ID).Url
			fmt.Printf(" @ %v", url)
		}
		_, _ = fmt.Print("\n")
	}
}

func showGit(git *profiles.GitConfig) {
	if git == nil {
		return
	}
	_, _ = fmt.Println("\tGit:")
	showGitUser("User", git.User)
	showGitUser("Author", git.Author)
	showGitUser("Committer", git.Committer)
}

func showGitUser(propName string, user *profiles.GitUser) {
	if user == nil {
		return
	}
	_, _ = fmt.Printf("\t\t%v:\n", propName)
	if user.Name != "" {
		_, _ = fmt.Printf("\t\t\tName:  %v\n", user.Name)
	}
	if user.Email != "" {
		_, _ = fmt.Printf("\t\t\tEmail: %v\n", user.Email)
	}
}
