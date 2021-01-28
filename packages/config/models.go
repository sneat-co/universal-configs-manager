package config

import "fmt"

// Ucm holds configuration for what settings are required and what not.
type Ucm struct {
	Path          string            `yaml:"-"`
	Raw           []byte            `yaml:"-"`
	ActiveProfile string            `yaml:"active_profile"`
	Profiles      []ProfileSettings `yaml:"profiles"`
	Git           *ModuleSettings   `yaml:"git,omitempty"`
	Hosts         *HostsSettings    `yaml:"hosts,omitempty"`
	EnvVars       *ModuleSettings   `yaml:"environment_variables,omitempty"`
	NodeJs        *NodeJsSettings   `yaml:"nodejs,omitempty"`
}

// ProfileSettings points to profile
type ProfileSettings struct {
	Name string `yaml:"name"`
	Path string `yaml:"path,omitempty"`
	URL  string `yaml:"url,omitempty"`
	ETag string `yaml:"etag,omitempty"`
}

func (v ProfileSettings) DisplayPath() string {
	if v.Path != "" {
		return v.Path
	}
	return fmt.Sprintf("~/.ucm.%v.yaml", v.Name)
}

// ModuleSettings
type ModuleSettings struct {
	Disabled bool     `yaml:"disabled,omitempty"`
	Requires []string `yaml:"requires,omitempty"`
}

type HostsSettings struct {
	ModuleSettings
	HeaderComment string `yaml:"header"`
	FooterComment string `yaml:"footer"`
	EditMode      string `yaml:"editMode"` // Options: inline, block_top, block_bottom, block_first, block_last
}

// NodeJsSettings holds configuration about settings related to NodeJS
type NodeJsSettings struct {
	ModuleSettings
	VersionManager string `yaml:"version_manager"`
}
