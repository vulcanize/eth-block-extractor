package ipfs

import (
	"fmt"
	"github.com/8thlight/block_watcher/pkg/fs"
)

type IpfsError struct {
	msg string
	err error
}

func (ie IpfsError) Error() string {
	return fmt.Sprintf("%s: %s", ie.msg, ie.err.Error())
}

type Publisher interface {
	Write(blockData []byte, blockNumber int64) ([]byte, error)
}

type IpfsPublisher struct {
	fs.BlockWriter
	IpfsWriter
}

func NewIpfsPublisher(fileWriter fs.BlockWriter, ipfsWriter IpfsWriter) *IpfsPublisher {
	return &IpfsPublisher{BlockWriter: fileWriter, IpfsWriter: ipfsWriter}
}

func (ip *IpfsPublisher) Write(blockData []byte, blockNumber int64) ([]byte, error) {
	filename, err := ip.BlockWriter.WriteBlockFile(blockData, blockNumber)
	if err != nil {
		return nil, IpfsError{msg: "Error writing block data", err: err}
	}
	output, err := ip.IpfsWriter.WriteToIpfs(filename)
	if err != nil {
		return nil, IpfsError{msg: "Error persisting block data", err: err}
	}
	return output, nil
}
