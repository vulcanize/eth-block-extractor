package level

// Implements Database interface for LevelDB
type LevelDatabase struct {
	Reader
}

func NewLevelDatabase(ldbReader Reader) *LevelDatabase {
	return &LevelDatabase{
		Reader: ldbReader,
	}
}

func (l LevelDatabase) Get(blockNumber int64) ([]byte, error) {
	n := uint64(blockNumber)
	blockHash := l.Reader.GetCanonicalHash(n)
	return l.Reader.GetHeaderRLP(blockHash, n), nil
}
