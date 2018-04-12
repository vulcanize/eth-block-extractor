package test_helpers

type MockDecoder struct {
	Called      bool
	PassedBytes []byte
	PassedOut   interface{}
	Err         error
}

func NewMockDecoder() *MockDecoder {
	return &MockDecoder{
		Called:      false,
		PassedBytes: nil,
		PassedOut:   nil,
		Err:         nil,
	}
}

func (md *MockDecoder) SetError(err error) {
	md.Err = err
}

func (md *MockDecoder) Decode(raw []byte, out interface{}) error {
	md.Called = true
	md.PassedBytes = raw
	md.PassedOut = out
	return md.Err
}
