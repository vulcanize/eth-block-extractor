package test_helpers

type MockRocksDatabaseReader struct {
	Err                  error
	OpenDatabaseCalled   bool
	GetBlockHeaderCalled bool
	PassedKey            []byte
	ReturnBytes          []byte
}

func NewMockRocksDatabaseReader() *MockRocksDatabaseReader {
	return &MockRocksDatabaseReader{
		Err:                  nil,
		OpenDatabaseCalled:   false,
		GetBlockHeaderCalled: false,
		PassedKey:            nil,
		ReturnBytes:          nil,
	}
}

func (mrdbr *MockRocksDatabaseReader) SetReturnBytes(returnBytes []byte) {
	mrdbr.ReturnBytes = returnBytes
}

func (mrdbr *MockRocksDatabaseReader) SetError(err error) {
	mrdbr.Err = err
}

func (mrdbr *MockRocksDatabaseReader) GetBlockHeader(key []byte) ([]byte, error) {
	mrdbr.GetBlockHeaderCalled = true
	mrdbr.PassedKey = key
	if mrdbr.Err != nil {
		return nil, mrdbr.Err
	}
	return mrdbr.ReturnBytes, nil
}

func (mrdbr *MockRocksDatabaseReader) OpenDatabaseForReadOnlyColumnFamilies(name string) error {
	mrdbr.OpenDatabaseCalled = true
	return nil
}
