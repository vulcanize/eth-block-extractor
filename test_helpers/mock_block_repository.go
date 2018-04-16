package test_helpers

import "github.com/vulcanize/vulcanizedb/pkg/core"

type MockBlockRepository struct {
	CalledCount        int
	PassedBlockNumbers []int64
	ReturnBlocks       []core.Block
	Err                error
}

func NewMockBlockRepository() *MockBlockRepository {
	return &MockBlockRepository{
		CalledCount:        0,
		PassedBlockNumbers: nil,
		ReturnBlocks:       nil,
		Err:                nil,
	}
}

func (mbr *MockBlockRepository) SetReturnBlocks(blocks []core.Block) {
	mbr.ReturnBlocks = blocks
}

func (mbr *MockBlockRepository) SetError(err error) {
	mbr.Err = err
}

func (MockBlockRepository) CreateOrUpdateBlock(block core.Block) error {
	panic("implement me")
}

func (mbr *MockBlockRepository) GetBlock(blockNumber int64) (core.Block, error) {
	mbr.CalledCount++
	mbr.PassedBlockNumbers = append(mbr.PassedBlockNumbers, blockNumber)
	if mbr.Err != nil {
		return core.Block{}, mbr.Err
	}
	blockToReturn := mbr.ReturnBlocks[0]
	mbr.ReturnBlocks = mbr.ReturnBlocks[1:]
	return blockToReturn, nil
}

func (MockBlockRepository) MissingBlockNumbers(startingBlockNumber int64, endingBlockNumber int64) []int64 {
	panic("implement me")
}

func (MockBlockRepository) SetBlocksStatus(chainHead int64) {
	panic("implement me")
}
