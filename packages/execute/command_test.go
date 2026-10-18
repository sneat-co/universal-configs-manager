package execute

import (
	"errors"
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
			origRunCmd := runCmd
			defer func() { runCmd = origRunCmd }()

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

		t.Run("executes_real_command", func(t *testing.T) {
			cmd := exec.Command("echo", "test-output")
			out, err := Command(cmd)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if out != "test-output\n" {
				t.Errorf("expected 'test-output\\n', got %q", out)
			}
		})
	})

	t.Run("returns_error", func(t *testing.T) {
		origRunCmd := runCmd
		defer func() { runCmd = origRunCmd }()
		runCmd = func(cmd *exec.Cmd) error {
			return errors.New("exec error")
		}
		_, err := Command(&exec.Cmd{})
		if err == nil || err.Error() != "exec error" {
			t.Errorf("expected exec error, got %v", err)
		}
	})
}
