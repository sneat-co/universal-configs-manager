package execute

import (
	"os/exec"
	"testing"
)

func TestCommand(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		t.Run("if_nil_command_passed", func(t *testing.T) {
			defer func() {
				err := recover()
				if err == nil {
					t.Fatal("expected to panic if nil passed for cmd parameter")
				}
			}()
			_, _ = Command(nil)
		})
	})
	t.Run("succeed", func(t *testing.T) {
		t.Run("sets_command_output", func(t *testing.T) {
			command := &exec.Cmd{}
			runCmd = func(cmd *exec.Cmd) error {
				if cmd != command {
					t.Errorf("executed not the command passed as parameter")
				}
				if command.Stdout == nil {
					t.Error("command.Stdout == nil")
				}
				return nil
			}
			_, err := Command(command)
			if err != nil {
				t.Errorf("expected to return no error")
			}
		})
	})
}
