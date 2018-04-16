package fs

import "os"

// Wraps os.Create
type Creator interface {
	Create(name string) (*os.File, error)
}

type FileCreator struct {
}

func (FileCreator) Create(name string) (*os.File, error) {
	return os.Create(name)
}
