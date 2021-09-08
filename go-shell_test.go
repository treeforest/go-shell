package shell

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	_, err := Run("ls", "/data")
	if err != nil {
		t.Fatal(err)
	}

	ls := Run
	_, err = ls("cd", "/home", "&&", "ls", ".")
	if err != nil {
		t.Fatal(err)
	}
}
