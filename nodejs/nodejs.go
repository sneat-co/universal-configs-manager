package nodejs

func GetRegistryByName(name string) RegistryInfo {
	return registries[name]
}
