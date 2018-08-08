package eth_block_receipts

import (
	"gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"
	"gx/ipfs/QmZtNq8dArGfnpCZfx2pUNY7UcjGhVp5qqwQ4hH6mpTMRQ/go-ipld-format"
)

type EthReceiptNode struct {
	raw []byte
	cid *cid.Cid
}

func (node *EthReceiptNode) RawData() []byte {
	return node.raw
}

func (node *EthReceiptNode) Cid() *cid.Cid {
	return node.cid
}

func (*EthReceiptNode) String() string {
	panic("implement me")
}

func (*EthReceiptNode) Loggable() map[string]interface{} {
	panic("implement me")
}

func (*EthReceiptNode) Resolve(path []string) (interface{}, []string, error) {
	panic("implement me")
}

func (*EthReceiptNode) Tree(path string, depth int) []string {
	panic("implement me")
}

func (*EthReceiptNode) ResolveLink(path []string) (*format.Link, []string, error) {
	panic("implement me")
}

func (*EthReceiptNode) Copy() format.Node {
	panic("implement me")
}

func (*EthReceiptNode) Links() []*format.Link {
	panic("implement me")
}

func (*EthReceiptNode) Stat() (*format.NodeStat, error) {
	panic("implement me")
}

func (*EthReceiptNode) Size() (uint64, error) {
	panic("implement me")
}
