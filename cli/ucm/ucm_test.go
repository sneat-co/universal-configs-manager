package main

import (
	"errors"
	"os"
	"testing"

	"github.com/jessevdk/go-flags"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestMain_Execution(t *testing.T) {
	origParse := parserParse
	origExit := osExit
	t.Cleanup(func() {
		parserParse = origParse
		osExit = origExit
	})

	t.Run("succeeds_without_exit", func(t *testing.T) {
		called := false
		parserParse = func() ([]string, error) { return nil, nil }
		osExit = func(int) { called = true }
		main()
		if called {
			t.Fatal("did not expect osExit to be called")
		}
	})

	t.Run("exits_zero_on_help", func(t *testing.T) {
		exitCode := -1
		parserParse = func() ([]string, error) {
			return nil, &flags.Error{Type: flags.ErrHelp}
		}
		osExit = func(code int) { exitCode = code }
		main()
		if exitCode != 0 {
			t.Fatalf("expected exit code 0, got %d", exitCode)
		}
	})

	t.Run("exits_one_on_other_error", func(t *testing.T) {
		exitCode := -1
		parserParse = func() ([]string, error) {
			return nil, errors.New("parse error")
		}
		osExit = func(code int) { exitCode = code }
		main()
		if exitCode != 1 {
			t.Fatalf("expected exit code 1, got %d", exitCode)
		}
	})
}
