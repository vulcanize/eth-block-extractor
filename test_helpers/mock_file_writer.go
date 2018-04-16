package test_helpers

type MockFileWriter struct {
	Called         bool
	PassedFilename string
	PassedData     []byte
}

func NewMockFileWriter() *MockFileWriter {
	return &MockFileWriter{
		Called:         false,
		PassedFilename: "",
		PassedData:     nil,
	}
}

func (mfw *MockFileWriter) WriteFile(filename string, data []byte) error {
	mfw.Called = true
	mfw.PassedFilename = filename
	mfw.PassedData = data
	return nil
}
