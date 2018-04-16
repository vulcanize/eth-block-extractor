package test_helpers

import "os/exec"

type MockCommander struct {
	Args   []string
	Called bool
	Name   string
}

func NewMockCommander() *MockCommander {
	return &MockCommander{
		Args:   nil,
		Called: false,
		Name:   "",
	}
}

func (mc *MockCommander) Command(name string, arg ...string) *exec.Cmd {
	mc.Called = true
	mc.Name = name
	mc.Args = append(mc.Args, arg...)
	return exec.Command("echo", "hello")
}
