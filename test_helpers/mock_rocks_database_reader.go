package test_helpers

type MockRocksDatabaseReader struct {
	GetBlockHashErr      error
	GetBlockHeaderErr    error
	OpenDatabaseCalled   bool
	GetBlockHashCalled   bool
	GetBlockHeaderCalled bool
	PassedBlockHashKey   []byte
	PassedBlockHeaderKey []byte
	ReturnHashBytes      []byte
	ReturnHeaderBytes    []byte
}

func NewMockRocksDatabaseReader() *MockRocksDatabaseReader {
	return &MockRocksDatabaseReader{
		GetBlockHashErr:      nil,
		GetBlockHeaderErr:    nil,
		OpenDatabaseCalled:   false,
		GetBlockHashCalled:   false,
		GetBlockHeaderCalled: false,
		PassedBlockHashKey:   nil,
		PassedBlockHeaderKey: nil,
		ReturnHashBytes:      nil,
		ReturnHeaderBytes:    nil,
	}
}

func (mrdbr *MockRocksDatabaseReader) SetReturnHashBytes(returnBytes []byte) {
	mrdbr.ReturnHashBytes = returnBytes
}

func (mrdbr *MockRocksDatabaseReader) SetReturnHeaderBytes(returnBytes []byte) {
	mrdbr.ReturnHeaderBytes = returnBytes
}

func (mrdbr *MockRocksDatabaseReader) SetGetHashError(err error) {
	mrdbr.GetBlockHashErr = err
}

func (mrdbr *MockRocksDatabaseReader) SetGetHeaderError(err error) {
	mrdbr.GetBlockHeaderErr = err
}

func (mrdbr *MockRocksDatabaseReader) GetBlockHash(key []byte) ([]byte, error) {
	mrdbr.GetBlockHashCalled = true
	mrdbr.PassedBlockHashKey = key
	if mrdbr.GetBlockHashErr != nil {
		return nil, mrdbr.GetBlockHashErr
	}
	return mrdbr.ReturnHashBytes, nil
}

func (mrdbr *MockRocksDatabaseReader) GetBlockHeader(key []byte) ([]byte, error) {
	mrdbr.GetBlockHeaderCalled = true
	mrdbr.PassedBlockHeaderKey = key
	if mrdbr.GetBlockHeaderErr != nil {
		return nil, mrdbr.GetBlockHeaderErr
	}
	return mrdbr.ReturnHeaderBytes, nil
}

func (mrdbr *MockRocksDatabaseReader) OpenDatabaseForReadOnlyColumnFamilies(name string) error {
	mrdbr.OpenDatabaseCalled = true
	return nil
}
