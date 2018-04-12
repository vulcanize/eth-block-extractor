package rocks

import (
	"encoding/binary"
	"errors"
)

var (
	ErrBlockHashNotFound   = errors.New("Block hash not found.")
	ErrBlockHeaderNotFound = errors.New("Block header not found.")
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

func (r *RocksDatabase) Get(blockNumber int64) ([]byte, error) {
	blockHashKey := GetKeyForBlockHash(blockNumber)
	rawBlockHash, err := r.reader.GetBlockHash(blockHashKey)
	if err != nil {
		return nil, err
	}
	if len(rawBlockHash) == 0 {
		return nil, ErrBlockHashNotFound
	}
	preparedBlockHash := GetKeyForBlock(rawBlockHash)
	rawBlock, err := r.reader.GetBlockHeader(preparedBlockHash)
	if err != nil {
		return nil, err
	}
	if len(rawBlock) == 0 {
		return nil, ErrBlockHeaderNotFound
	}
	decompressedBlock, err := r.decoder.Decompress(rawBlock)
	if err != nil {
		return nil, err
	}
	return decompressedBlock, nil
}

// The key for a block hash in Parity's Rocks DB is a byte array consisting of the byte 1 concatenated with
// a size 4 byte array consisting of a byte representation of the block number integer.
// e.g. for block 5477822: [1, 0, 83, 149, 190]
func GetKeyForBlockHash(blockNumber int64) []byte {
	n := uint32(blockNumber)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, n)
	blockHash := append([]byte{1}, bs...)
	return blockHash
}

// The block hash returned by `GetKeyForBlockHash` comes back prepended with the byte 160.
// This byte needs to be removed to generate the hash used to fetch the block.
func GetKeyForBlock(rawBlockKey []byte) []byte {
	return rawBlockKey[1:]
}
