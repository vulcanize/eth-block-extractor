package test_helpers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type MockLevelDatabaseReader struct {
	Called       bool
	PassedHash   common.Hash
	PassedNumber uint64
	ReturnHash   common.Hash
}

func NewMockLevelDatabaseReader() *MockLevelDatabaseReader {
	return &MockLevelDatabaseReader{
		Called:       false,
		PassedHash:   common.Hash{},
		PassedNumber: 0,
	}
}

func (mldr *MockLevelDatabaseReader) SetReturnHash(hash common.Hash) {
	mldr.ReturnHash = hash
}

func (mldr *MockLevelDatabaseReader) GetCanonicalHash(number uint64) common.Hash {
	return mldr.ReturnHash
}

func (mldr *MockLevelDatabaseReader) GetHeaderRLP(hash common.Hash, number uint64) rlp.RawValue {
	mldr.Called = true
	mldr.PassedHash = hash
	mldr.PassedNumber = number
	return nil
}
