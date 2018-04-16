package test_helpers

import "github.com/vulcanize/vulcanizedb/pkg/core"

type MockDatabase struct {
	CalledCount  int
	PassedBlocks []core.Block
	ReturnBytes  [][]byte
	Err          error
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		CalledCount:  0,
		PassedBlocks: nil,
		ReturnBytes:  nil,
		Err:          nil,
	}
}

func (md *MockDatabase) SetReturnBytes(returnBytes [][]byte) {
	md.ReturnBytes = returnBytes
}

func (md *MockDatabase) SetError(err error) {
	md.Err = err
}

func (md *MockDatabase) Get(block core.Block) ([]byte, error) {
	md.CalledCount++
	md.PassedBlocks = append(md.PassedBlocks, block)
	if md.Err != nil {
		return nil, md.Err
	}
	bytesToReturn := md.ReturnBytes[0]
	md.ReturnBytes = md.ReturnBytes[1:]
	return bytesToReturn, nil
}
