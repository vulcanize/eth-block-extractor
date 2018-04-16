package fs

import (
	"io/ioutil"
	"os"
)

// wraps ioutil.WriteFile
type Writer interface {
	WriteFile(filename string, data []byte) error
}

type FileWriter struct {
}

func (FileWriter) WriteFile(filename string, data []byte) error {
	permissions := os.FileMode(0644)
	return ioutil.WriteFile(filename, data, permissions)
}
