package test_helpers

type MockPublisher struct {
	CalledCount        int
	PassedBlockDatas   [][]byte
	PassedBlockNumbers []int64
	ReturnBytes        [][]byte
	Err                error
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{
		CalledCount:        0,
		PassedBlockDatas:   [][]byte{},
		PassedBlockNumbers: []int64{},
		ReturnBytes:        nil,
		Err:                nil,
	}
}

func (mp *MockPublisher) SetReturnBytes(returnBytes [][]byte) {
	mp.ReturnBytes = returnBytes
}

func (mp *MockPublisher) SetError(err error) {
	mp.Err = err
}

func (mp *MockPublisher) Write(blockData []byte, blockNumber int64) ([]byte, error) {
	mp.CalledCount++
	mp.PassedBlockDatas = append(mp.PassedBlockDatas, blockData)
	mp.PassedBlockNumbers = append(mp.PassedBlockNumbers, blockNumber)
	if mp.Err != nil {
		return nil, mp.Err
	}
	bytesToReturn := mp.ReturnBytes[0]
	mp.ReturnBytes = mp.ReturnBytes[1:]
	return bytesToReturn, nil
}
