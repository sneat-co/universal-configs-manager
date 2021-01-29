package execute

import (
	"bytes"
	"os/exec"
)

func Command(cmd *exec.Cmd) (string, error) {
	if cmd == nil {
		panic("cmd parameter is nil")
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := runCmd(cmd); err != nil {
		return "", err
	}
	return out.String(), nil
}

var runCmd = func(cmd *exec.Cmd) error {
	return cmd.Run()
}
