package ipfs

import (
	"context"

	"github.com/vulcanize/go-ipfs/core"
	"github.com/vulcanize/go-ipfs/repo/fsrepo"

	ipld "gx/ipfs/QmZtNq8dArGfnpCZfx2pUNY7UcjGhVp5qqwQ4hH6mpTMRQ/go-ipld-format"
)

type IPFS struct {
	n   *core.IpfsNode
	ctx context.Context
}

func (ipfs IPFS) Add(node ipld.Node) error {
	return ipfs.n.DAG.Add(ipfs.n.Context(), node)
}

func InitIPFSNode(repoPath string) (*IPFS, error) {
	r, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	cfg := &core.BuildCfg{
		Online: false,
		Repo:   r,
	}
	ipfsNode, err := core.NewNode(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &IPFS{n: ipfsNode, ctx: ctx}, nil
}
