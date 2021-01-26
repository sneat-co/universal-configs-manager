package profiles

type Host struct {
	IP      string `yaml:"ip"`
	Name    string `yaml:"name"`
	Comment string `yaml:"comment"`
}

type Hosts struct {
	Entries []Host `yaml:"entries"`
}
