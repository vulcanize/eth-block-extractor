package ipfs

import (
	"github.com/8thlight/block_watcher/pkg/fs"
)

type IpfsWriter interface {
	WriteToIpfs(filename string) ([]byte, error)
}

type IpfsEthBlockWriter struct {
	fs.Commander
}

func NewIpfsEthBlockWriter(commander fs.Commander) *IpfsEthBlockWriter {
	return &IpfsEthBlockWriter{Commander: commander}
}

func (iebw *IpfsEthBlockWriter) WriteToIpfs(filename string) ([]byte, error) {
	ipfsCommand := iebw.Commander.Command("ipfs", "dag", "put", "--input-enc", "raw", "--format", "eth-block", filename)
	return ipfsCommand.Output()
}
