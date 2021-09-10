package shell

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	p := Run("ls", "/data")
	if p.ExitStatus != 0 {
		t.Fatal(p.Error())
	}
	t.Log(p.Stdout.String())

	ls := Run
	p = ls("cd", "/home", "&&", "ls", ".")
	if p.ExitStatus != 0 {
		t.Fatal(p.Error())
	}
	t.Log(p.Stdout.String())

	CP := Cmd("cp", "-Rf").ErrFn()
	err := CP("/data/outdir", "/home/a")
	if err != nil {
		t.Fatal(err)
	}
}
