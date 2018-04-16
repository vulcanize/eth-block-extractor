package test_helpers

type MockBlockFileWriter struct {
	Called            bool
	PassedBlockData   []byte
	PassedBlockNumber int64
	ReturnString      string
}

func NewMockBlockFileWriter() *MockBlockFileWriter {
	return &MockBlockFileWriter{
		Called:            false,
		PassedBlockData:   nil,
		PassedBlockNumber: 0,
		ReturnString:      "",
	}
}

func (mbfw *MockBlockFileWriter) WriteBlockFile(blockData []byte, blockNumber int64) (string, error) {
	mbfw.Called = true
	mbfw.PassedBlockData = blockData
	mbfw.PassedBlockNumber = blockNumber
	return mbfw.ReturnString, nil
}

func (mbfw *MockBlockFileWriter) SetReturnString(arg string) {
	mbfw.ReturnString = arg
}
