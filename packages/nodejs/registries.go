package nodejs

// RegistryInfo holds information about Node package manager registry
type RegistryInfo struct {
	Url      string `yaml:"url" json:"url"`
	Home     string `yaml:"home" json:"home"`
	IsCustom bool   `yaml:"isCustom,omitempty" json:"isCustom,omitempty"`
}

var registries = map[string]RegistryInfo{
	"npm": {
		Url:  "https://registry.npmjs.org/",
		Home: "https://www.npmjs.org",
	},
	"yarn": {
		Url:  "https://registry.yarnpkg.com/",
		Home: "https://yarnpkg.com",
	},
}
