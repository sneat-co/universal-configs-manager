package profiles

type EnvVars struct {
	Remove []string          `yaml:"remove"`
	Set    map[string]string `yaml:"set"`
}

func (v *EnvVars) IsEmpty() bool {
	return v == nil || len(v.Remove) == 0 && len(v.Set) == 0
}
