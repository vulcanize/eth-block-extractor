package ipfs

import (
	"github.com/vulcanize/eth-block-extractor/pkg/wrappers/go-cid"
	"gx/ipfs/QmWi2BYBL5gJ3CiAiQchg6rn1A8iBsrWy51EYxvHVjFvLb/go-ipld-format"
)

type Reader interface {
	Read(cid string) (format.Node, error)
}

type IpldReader struct {
	decoder  go_cid.ICidDecoder
	getter   Getter
	resolver Resolver
}

func NewIpldReader(decoder go_cid.ICidDecoder, getter Getter, resolver Resolver) *IpldReader {
	return &IpldReader{
		decoder:  decoder,
		getter:   getter,
		resolver: resolver,
	}
}

func (reader *IpldReader) Read(cidString string) (format.Node, error) {
	cid, err := reader.decoder.Decode(cidString)
	if err != nil {
		return nil, err
	}
	block, err := reader.getter.Get(cid)
	if err != nil {
		return nil, err
	}
	return reader.resolver.Resolve(block)
}
