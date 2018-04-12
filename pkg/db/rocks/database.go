package rocks

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

type RocksDatabase struct {
	decoder Decompressor
	reader  Reader
}

func NewRocksDatabase(decoder Decompressor, reader Reader) *RocksDatabase {
	return &RocksDatabase{
		decoder: decoder,
		reader:  reader,
	}
}

func (r *RocksDatabase) Get(block core.Block) ([]byte, error) {
	key := common.FromHex(block.Hash)
	toReturn, err := r.reader.GetBlockHeader(key)
	if err != nil {
		return nil, err
	}
	result, err := r.decoder.Decompress(toReturn)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not decode data: %s", err.Error()))
	}
	return result, nil
}
