package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

var (
	// cat /etc/shells
	Shell = []string{"/bin/bash", "-c"}
)

func Cmd(args ...string) *Command {
	c := new(Command)
	c.args = make([]string, 0)
	c.args = args
	return c
}

func Run(args ...string) *Process {
	return Cmd(args...).Run()
}

type Command struct {
	args []string
}

type Process struct {
	Cmd        *exec.Cmd
	Stdout     *bytes.Buffer
	Stderr     *bytes.Buffer
	ExitStatus int // 0 is ok, others are error

	err error
}

func (p *Process) Error() error {
	return fmt.Errorf("%s\n%v: %s", p.Stdout, p.err, p.Stderr)
}

func (c *Command) addArgs(args ...string) {
	c.args = append(c.args, args...)
}

func (c *Command) Run() *Process {
	sys := runtime.GOOS

	if sys == "linux" {
		return c.linuxRun()
	}

	if sys == "windows" {

	}

	return nil
}

func (c *Command) linuxRun() *Process {
	cmd := exec.Command(Shell[0], append(Shell[1:], strings.Join(c.args, " "))...)

	p := new(Process)
	p.Stdout = bytes.NewBuffer(nil)
	p.Stderr = bytes.NewBuffer(nil)
	p.ExitStatus = 0

	defer func() {
		if err := recover(); err != nil {
			p.err =fmt.Errorf("%v", err)
			p.ExitStatus = -1
		}
	}()

	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stderr

	if err := cmd.Run(); err != nil {
		p.err = err
		if exitErr, ok := err.(*exec.ExitError); ok {
			if stat, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				p.ExitStatus = stat.ExitStatus()
			}
		} else {
			// unknown error
			p.ExitStatus = -1
		}
	}

	p.Cmd = cmd
	return p
}

func (c *Command) ErrFn() func(...string) error {
	return func(args ...string) error {
		cmd := &Command{c.args}
		cmd.addArgs(args...)

		p := cmd.Run()
		if p.ExitStatus != 0 {
			return p.Error()
		}

		return nil
	}
}