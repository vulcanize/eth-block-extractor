package test_helpers

type MockPublisher struct {
	CalledCount      int
	PassedBlockDatas [][]byte
	ReturnBytes      [][]byte
	Err              error
}

func (mp *MockPublisher) Write(blockData []byte) (string, error) {
	mp.CalledCount++
	mp.PassedBlockDatas = append(mp.PassedBlockDatas, blockData)
	if mp.Err != nil {
		return "", mp.Err
	}
	bytesToReturn := mp.ReturnBytes[0]
	mp.ReturnBytes = mp.ReturnBytes[1:]
	return string(bytesToReturn), nil
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{
		CalledCount:      0,
		PassedBlockDatas: [][]byte{},
		ReturnBytes:      nil,
		Err:              nil,
	}
}

func (mp *MockPublisher) SetReturnBytes(returnBytes [][]byte) {
	mp.ReturnBytes = returnBytes
}

func (mp *MockPublisher) SetError(err error) {
	mp.Err = err
}
