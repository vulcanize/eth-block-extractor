package fs

import (
	"fmt"
)

type WriteError struct {
	msg string
	err error
}

func (we WriteError) Error() string {
	return fmt.Sprintf("%s: %s", we.msg, we.err.Error())
}

type BlockWriter interface {
	WriteBlockFile(blockData []byte, blockNumber int64) (string, error)
}

type BlockFileWriter struct {
	Creator
	Writer
}

func NewBlockFileWriter(creator Creator, writer Writer) *BlockFileWriter {
	return &BlockFileWriter{Creator: creator, Writer: writer}
}

func (bfw *BlockFileWriter) WriteBlockFile(blockData []byte, blockNumber int64) (string, error) {
	filename := fmt.Sprintf("blocks/block_%d.bytes", blockNumber)
	file, err := bfw.Creator.Create(filename)
	if err != nil {
		return "", WriteError{msg: "Error creating file for data", err: err}
	}
	defer file.Close()
	err = bfw.Writer.WriteFile(filename, blockData)
	if err != nil {
		return "", WriteError{msg: "Error writing data to file", err: err}
	}
	return filename, nil
}
