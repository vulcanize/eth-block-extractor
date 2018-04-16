package test_helpers

type MockIpfsWriter struct {
	Called         bool
	PassedFilename string
}

func NewMockIpfsWriter() *MockIpfsWriter {
	return &MockIpfsWriter{
		Called:         false,
		PassedFilename: "",
	}
}

func (miw *MockIpfsWriter) WriteToIpfs(filename string) ([]byte, error) {
	miw.Called = true
	miw.PassedFilename = filename
	return []byte{}, nil
}
