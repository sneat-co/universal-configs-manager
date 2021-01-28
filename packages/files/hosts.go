package files

import (
	"fmt"
	"strings"
	"ucm/packages/profiles"
)

func SwitchHosts(_ *profiles.Hosts) error {
	return nil
}

func ShowHosts(hosts *profiles.Hosts) error {
	if hosts == nil || len(hosts.Entries) == 0 {
		return nil
	}
	// TODO: sort by "reverse" name. E.g. "subdomain.example.com" is sorted by "com.example.subdomain"
	for _, host := range hosts.Entries {
		if strings.TrimSpace(host.Comment) == "" {
			_, _ = fmt.Printf("\t%v: %v\n", host.Name, host.IP)
		} else {
			_, _ = fmt.Printf("\t%v: %v # %v\n", host.Name, host.IP, host.Comment)
		}
	}
	return nil
}
