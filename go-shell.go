package shell

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

var (
	Shell = []string{"/bin/sh", "-c"}
)

func Cmd(args ...string) *Command {
	c := new(Command)
	c.args = args
	return c
}

func Run(args ...string) (string, error) {
	return Cmd(args...).Run()
}

type Command struct {
	args []string
}

func (c *Command) Run() (string, error) {
	cmd := exec.Command(Shell[0], append(Shell[1:], c.args...)...)

	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), fmt.Errorf("%v: %s", err, stderr.String())
	}

	if stderr.String() != "" {
		return stdout.String(), errors.New(stderr.String())
	}

	return stdout.String(), nil
}
