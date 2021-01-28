package nodejs

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"ucm/packages/execute"
)

func SwitchVersion(version string) error {
	if version == "" {
		return nil
	}
	// TODO: Get current version of Node and do not switch if matching.
	currentNodeVersion, err := getCurrentNodeVersion()
	if err != nil {
		return err
	}
	if currentNodeVersion == version {
		_, _ = fmt.Printf("Target NodeJS version %v is active, no switching required.\n", version)
		return nil
	}
	_, _ = fmt.Printf("Switching NodeJS version from %v to %v ...\n", currentNodeVersion, version)
	cmd := exec.Command("nvm", "use", version)
	var output string
	if output, err = execute.Command(cmd); err != nil {
		return err
	}
	_, _ = fmt.Printf("NVM output: %q\n", output)
	return nil
}

func getCurrentNodeVersion() (string, error) {
	_, _ = fmt.Print("Getting current NodeJS version...")
	cmd := exec.Command("node", "--version")
	s, err := execute.Command(cmd)
	s = strings.TrimSpace(s)
	if err != nil {
		_, _ = fmt.Println("")
		return "", err
	}
	if strings.HasPrefix(s, "v") {
		s = s[1:]
	}
	_, _ = fmt.Println(" ", s)
	return s, nil
}

func SetRegistry(registry string) (string, error) {
	if registry == "" {
		return "", nil
	}
	_, _ = fmt.Printf("Setting NodeJS repository to: %v...\n", registry)
	return setNodeRegistryUsingNpmConfigSet(registry)
}

func setNodeRegistryUsingNpmConfigSet(registry string) (string, error) {
	if registry == "" {
		return "", nil
	}
	// npm config set registry
	var url string
	if strings.HasPrefix(registry, "https://") && len(registry) > len("https://") ||
		strings.HasPrefix(registry, "http://") && len(registry) > len("http://") {
		url = registry
	} else {
		url = GetRegistryByName(registry).Url
	}

	if url == "" {
		return "", fmt.Errorf("not allowed to set NodeJS registry to empty string, no url defined for: %v", registry)
	}

	return callNpmConfigSet("registry", fmt.Sprintf(`"%v"`, url))
}

func callNpmConfigSet(name, value string) (string, error) {
	fmt.Printf(`Running:
> npm config set %v %v
`, name, value)
	cmd := exec.Command("npm", "config", "set", name, value)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run `npm set config registry`: %w", err)
	}
	//_, _ = fmt.Printf("NPM config set registry output: %q\n", out.String())
	return out.String(), nil
}
