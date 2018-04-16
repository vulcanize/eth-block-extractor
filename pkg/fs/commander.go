package fs

import "os/exec"

// wraps exec.Command
type Commander interface {
	Command(name string, arg ...string) *exec.Cmd
}

type ExecCommander struct {
}

func (ExecCommander) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}
