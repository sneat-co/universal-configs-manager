package profiles

// GitConfig defines settings available for Git
type GitConfig struct {
	User      *GitUser `yaml:"user,omitempty"`
	Author    *GitUser `yaml:"author,omitempty"`
	Committer *GitUser `yaml:"committer,omitempty"`
}

// GitConfig defines user related parameters
type GitUser struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
