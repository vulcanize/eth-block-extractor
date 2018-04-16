package level

import (
	"github.com/ethereum/go-ethereum/common"
	vulcCore "github.com/vulcanize/vulcanizedb/pkg/core"
)

// Implements Database interface for LevelDB
type LevelDatabase struct {
	Reader
}

func NewLevelDatabase(ldbReader Reader) *LevelDatabase {
	return &LevelDatabase{
		Reader: ldbReader,
	}
}

func (l LevelDatabase) Get(block vulcCore.Block) ([]byte, error) {
	blockHash := common.HexToHash(block.Hash)
	blockNumber := uint64(block.Number)
	return l.Reader.GetHeaderRLP(blockHash, blockNumber), nil
}
