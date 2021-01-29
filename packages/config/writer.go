package config

import "errors"

var errNotImplemented = errors.New("not implemented yet")

func Write(config Ucm) error {
	if config.Path == "" {
		return errors.New("can't write config to unknown path")
	}
	return errNotImplemented
}
