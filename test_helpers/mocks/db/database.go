package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/gomega"
)

type MockDatabase struct {
	computeBlockStateTrieErr                          error
	computeBlockStateTriePassedCurrentBlock           *types.Block
	computeBlockStateTriePassedParentBlock            *types.Block
	computeBlockStateTrieReturnHash                   common.Hash
	getBlockBodyByBlockNumberErr                      error
	getBlockBodyByBlockNumberPassedBlockNumbers       []int64
	getBlockBodyByBlockNumberReturnBytes              [][]byte
	getBlockByBlockNumberPassedNumbers                []int64
	getBlockByBlockNumberReturnBlock                  *types.Block
	getBlockHeaderByBlockNumberErr                    error
	getBlockHeaderByBlockNumberPassedBlockNumbers     []int64
	getBlockHeaderByBlockNumberReturnBytes            [][]byte
	getBlockReceiptsPassedBlockNumbers                []int64
	getBlockReceiptsReturnReceipts                    types.Receipts
	getStateAndStorageTrieNodesErr                    error
	getStateAndStorageTrieNodesPassedRoot             common.Hash
	getStateAndStorageTrieNodesReturnStateTrieBytes   [][]byte
	getStateAndStorageTrieNodesReturnStorageTrieBytes [][]byte
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		computeBlockStateTrieErr:                          nil,
		computeBlockStateTriePassedCurrentBlock:           nil,
		computeBlockStateTriePassedParentBlock:            nil,
		computeBlockStateTrieReturnHash:                   common.Hash{},
		getBlockBodyByBlockNumberErr:                      nil,
		getBlockBodyByBlockNumberPassedBlockNumbers:       nil,
		getBlockBodyByBlockNumberReturnBytes:              nil,
		getBlockByBlockNumberPassedNumbers:                nil,
		getBlockByBlockNumberReturnBlock:                  nil,
		getBlockHeaderByBlockNumberErr:                    nil,
		getBlockHeaderByBlockNumberPassedBlockNumbers:     nil,
		getBlockHeaderByBlockNumberReturnBytes:            nil,
		getBlockReceiptsPassedBlockNumbers:                nil,
		getBlockReceiptsReturnReceipts:                    nil,
		getStateAndStorageTrieNodesErr:                    nil,
		getStateAndStorageTrieNodesPassedRoot:             common.Hash{},
		getStateAndStorageTrieNodesReturnStateTrieBytes:   nil,
		getStateAndStorageTrieNodesReturnStorageTrieBytes: nil,
	}
}

func (db *MockDatabase) SetComputeBlockStateTrieError(err error) {
	db.computeBlockStateTrieErr = err
}

func (db *MockDatabase) SetComputeBlockStateTrieReturnHash(hash common.Hash) {
	db.computeBlockStateTrieReturnHash = hash
}

func (db *MockDatabase) SetGetBlockBodyByBlockNumberError(err error) {
	db.getBlockBodyByBlockNumberErr = err
}

func (db *MockDatabase) SetGetBlockBodyByBlockNumberReturnBytes(returnBytes [][]byte) {
	db.getBlockBodyByBlockNumberReturnBytes = returnBytes
}

func (db *MockDatabase) SetGetBlockByBlockNumberReturnBlock(returnBlock *types.Block) {
	db.getBlockByBlockNumberReturnBlock = returnBlock
}

func (db *MockDatabase) SetGetBlockHeaderByBlockNumberError(err error) {
	db.getBlockHeaderByBlockNumberErr = err
}

func (db *MockDatabase) SetGetBlockHeaderByBlockNumberReturnBytes(returnBytes [][]byte) {
	db.getBlockHeaderByBlockNumberReturnBytes = returnBytes
}

func (db *MockDatabase) SetGetBlockReceiptsReturnReceipts(receipts types.Receipts) {
	db.getBlockReceiptsReturnReceipts = receipts
}

func (db *MockDatabase) SetGetStateAndStorageTrieNodesError(err error) {
	db.getStateAndStorageTrieNodesErr = err
}

func (db *MockDatabase) SetGetStateAndStorageTrieNodesReturnStateTrieBytes(returnBytes [][]byte) {
	db.getStateAndStorageTrieNodesReturnStateTrieBytes = returnBytes
}

func (db *MockDatabase) SetGetStateAndStorageTrieNodesReturnStorageTrieBytes(returnBytes [][]byte) {
	db.getStateAndStorageTrieNodesReturnStorageTrieBytes = returnBytes
}

func (db *MockDatabase) ComputeBlockStateTrie(currentBlock *types.Block, parentBlock *types.Block) (common.Hash, error) {
	db.computeBlockStateTriePassedCurrentBlock = currentBlock
	db.computeBlockStateTriePassedParentBlock = parentBlock
	return db.computeBlockStateTrieReturnHash, db.computeBlockStateTrieErr
}

func (db *MockDatabase) GetBlockBodyByBlockNumber(blockNumber int64) ([]byte, error) {
	db.getBlockBodyByBlockNumberPassedBlockNumbers = append(db.getBlockBodyByBlockNumberPassedBlockNumbers, blockNumber)
	if db.getBlockBodyByBlockNumberErr != nil {
		return nil, db.getBlockBodyByBlockNumberErr
	}
	returnBytes := db.getBlockBodyByBlockNumberReturnBytes[0]
	db.getBlockBodyByBlockNumberReturnBytes = db.getBlockBodyByBlockNumberReturnBytes[1:]
	return returnBytes, nil
}

func (db *MockDatabase) GetBlockByBlockNumber(blockNumber int64) *types.Block {
	db.getBlockByBlockNumberPassedNumbers = append(db.getBlockByBlockNumberPassedNumbers, blockNumber)
	return db.getBlockByBlockNumberReturnBlock
}

func (db *MockDatabase) GetBlockHeaderByBlockNumber(blockNumber int64) ([]byte, error) {
	db.getBlockHeaderByBlockNumberPassedBlockNumbers = append(db.getBlockHeaderByBlockNumberPassedBlockNumbers, blockNumber)
	if db.getBlockHeaderByBlockNumberErr != nil {
		return nil, db.getBlockHeaderByBlockNumberErr
	}
	returnBytes := db.getBlockHeaderByBlockNumberReturnBytes[0]
	db.getBlockHeaderByBlockNumberReturnBytes = db.getBlockHeaderByBlockNumberReturnBytes[1:]
	return returnBytes, nil
}

func (db *MockDatabase) GetBlockReceipts(blockNumber int64) types.Receipts {
	db.getBlockReceiptsPassedBlockNumbers = append(db.getBlockReceiptsPassedBlockNumbers, blockNumber)
	return db.getBlockReceiptsReturnReceipts
}

func (db *MockDatabase) GetStateAndStorageTrieNodes(root common.Hash) ([][]byte, [][]byte, error) {
	db.getStateAndStorageTrieNodesPassedRoot = root
	return db.getStateAndStorageTrieNodesReturnStateTrieBytes, db.getStateAndStorageTrieNodesReturnStorageTrieBytes, db.getStateAndStorageTrieNodesErr
}

func (db *MockDatabase) AssertComputeBlockStateTrieCalledWith(currentBlock *types.Block, parentBlock *types.Block) {
	Expect(db.computeBlockStateTriePassedCurrentBlock).To(Equal(currentBlock))
	Expect(db.computeBlockStateTriePassedParentBlock).To(Equal(parentBlock))
}

func (db *MockDatabase) AssertGetBlockBodyByBlockNumberCalledWith(blockNumbers []int64) {
	Expect(db.getBlockBodyByBlockNumberPassedBlockNumbers).To(Equal(blockNumbers))
}

func (db *MockDatabase) AssertGetBlockByBlockNumberCalledwith(blockNumbers []int64) {
	for i := 0; i < len(blockNumbers); i++ {
		Expect(db.getBlockByBlockNumberPassedNumbers).To(ContainElement(blockNumbers[i]))
	}
	for i := 0; i < len(db.getBlockByBlockNumberPassedNumbers); i++ {
		Expect(blockNumbers).To(ContainElement(db.getBlockByBlockNumberPassedNumbers[i]))
	}
}

func (db *MockDatabase) AssertGetBlockHeaderByBlockNumberCalledWith(blockNumbers []int64) {
	Expect(db.getBlockHeaderByBlockNumberPassedBlockNumbers).To(Equal(blockNumbers))
}

func (db *MockDatabase) AssertGetBlockReceiptsCalledWith(blockNumbers []int64) {
	Expect(db.getBlockReceiptsPassedBlockNumbers).To(Equal(blockNumbers))
}

func (db *MockDatabase) AssertGetStateTrieNodesCalledWith(root common.Hash) {
	Expect(db.getStateAndStorageTrieNodesPassedRoot).To(Equal(root))
}
