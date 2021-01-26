package profiles

type Vars = map[string]string

// Profile holds configs set
type Profile struct {
	EnvVars *EnvVars   `yaml:"environment_variables,omitempty"`
	NodeJS  *NodeJS    `yaml:"nodejs,omitempty"`
	Git     *GitConfig `yaml:"git,omitempty"`
	Hosts   *Hosts     `yaml:"hosts,omitempty"`
}
