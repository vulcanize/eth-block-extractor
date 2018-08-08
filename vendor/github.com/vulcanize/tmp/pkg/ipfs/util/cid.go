package util

import (
	"gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"
	mh "gx/ipfs/QmZyZDi491cCNTLfAhwcaDii2Kg4pwKRkhqQzURGDvY6ua/go-multihash"
)

func RawToCid(codec uint64, raw []byte) (*cid.Cid, error) {
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
