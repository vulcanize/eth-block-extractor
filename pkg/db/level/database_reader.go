package level

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/rlp"
)

// Wraps go-ethereum/core: GetHeaderRLP
type Reader interface {
	GetHeaderRLP(hash common.Hash, number uint64) rlp.RawValue
}

type LDBReader struct {
	core.DatabaseReader
}

func NewLevelDatabaseReader(reader core.DatabaseReader) *LDBReader {
	return &LDBReader{DatabaseReader: reader}
}

func (ldbr *LDBReader) GetHeaderRLP(hash common.Hash, number uint64) rlp.RawValue {
	return core.GetHeaderRLP(ldbr.DatabaseReader, hash, number)
}
