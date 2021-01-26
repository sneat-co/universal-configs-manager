package profiles

// NodeJS defines settings for NodeJS
type NodeJS struct {
	Version  string        `yaml:"version,omitempty"`
	Registry *NodeRegistry `yaml:"registry,omitempty"`
}

// NodeRegistry defines possible settings for NodeJS repository
type NodeRegistry struct {
	ID  string `yaml:"id,omitempty"`
	URL string `yaml:"url,omitempty"`
}
