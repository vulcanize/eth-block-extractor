package test_helpers

type MockDecmpressor struct {
	Called      bool
	PassedBytes []byte
}

func NewMockDecompressor() *MockDecmpressor {
	return &MockDecmpressor{
		Called:      false,
		PassedBytes: nil,
	}
}

func (md *MockDecmpressor) Decompress(raw []byte) ([]byte, error) {
	md.Called = true
	md.PassedBytes = raw
	return nil, nil
}
