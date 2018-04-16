package test_helpers

import (
	"os"
)

type MockFileCreator struct {
	Called     bool
	PassedName string
}

func NewMockFileCreator() *MockFileCreator {
	return &MockFileCreator{
		Called:     false,
		PassedName: "",
	}
}

func (mfc *MockFileCreator) Create(name string) (*os.File, error) {
	mfc.Called = true
	mfc.PassedName = name
	return nil, nil
}
