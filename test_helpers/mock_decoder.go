package test_helpers

type MockDecoder struct {
	Called      bool
	PassedBytes []byte
}

func NewMockDecoder() *MockDecoder {
	return &MockDecoder{
		Called:      false,
		PassedBytes: nil,
	}
}

func (md *MockDecoder) Decode(raw []byte) ([]byte, error) {
	md.Called = true
	md.PassedBytes = raw
	return nil, nil
}
