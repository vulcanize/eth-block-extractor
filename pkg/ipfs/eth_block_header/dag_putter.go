package eth_block_header

import (
	"github.com/8thlight/block_watcher/pkg/ipfs"
	"github.com/ethereum/go-ethereum/core/types"

	mh "gx/ipfs/QmZyZDi491cCNTLfAhwcaDii2Kg4pwKRkhqQzURGDvY6ua/go-multihash"
	cid "gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
	ipld "gx/ipfs/Qme5bWv7wtjUNGsK2BNGVUFPKiuxWrsqrtvYwCLRw8YFES/go-ipld-format"
)

const (
	EthBlockHeaderCode = 0x90
)

type BlockHeaderDagPutter struct {
	adder   ipfs.Adder
	decoder ipfs.Decoder
}

func NewBlockHeaderDagPutter(adder ipfs.Adder, decoder ipfs.Decoder) *BlockHeaderDagPutter {
	return &BlockHeaderDagPutter{adder: adder, decoder: decoder}
}

func (bhdp *BlockHeaderDagPutter) DagPut(raw []byte) (string, error) {
	nd, err := bhdp.getNodeForBlockHeader(raw)
	if err != nil {
		return "", err
	}
	err = bhdp.adder.Add(nd)
	if err != nil {
		return "", err
	}
	return nd.Cid().String(), nil
}

func (bhdp *BlockHeaderDagPutter) getNodeForBlockHeader(raw []byte) (ipld.Node, error) {
	var blockHeader types.Header
	err := bhdp.decoder.Decode(raw, &blockHeader)
	if err != nil {
		return nil, err
	}
	blockHeaderCid, err := rawToCid(EthBlockHeaderCode, raw)
	if err != nil {
		return nil, err
	}
	return &EthBlockHeaderNode{
		Header:  &blockHeader,
		cid:     blockHeaderCid,
		rawdata: raw,
	}, nil
}

func rawToCid(codec uint64, raw []byte) (*cid.Cid, error) {
	c, err := cid.Prefix{
		Codec:    codec,
		Version:  1,
		MhType:   mh.KECCAK_256,
		MhLength: -1,
	}.Sum(raw)
	if err != nil {
		return nil, err
	}
	return c, nil
}
