package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	state_wrapper "github.com/vulcanize/eth-block-extractor/pkg/wrappers/core/state"
	trie_wrapper "github.com/vulcanize/eth-block-extractor/pkg/wrappers/trie"
	"github.com/vulcanize/eth-block-extractor/test_helpers"
)

type MockStateDatabase struct {
	ReturnDB   state.Database
	ReturnTrie state_wrapper.GethTrie
}

func NewMockStateDatabase() *MockStateDatabase {
	return &MockStateDatabase{}
}

func (msdb *MockStateDatabase) CreateFakeUnderlyingDatabase() state.Database {
	return &mockStateDatabase{}
}

func (msdb *MockStateDatabase) Database() state.Database {
	return msdb.ReturnDB
}

func (msdb *MockStateDatabase) OpenTrie(root common.Hash) (state_wrapper.GethTrie, error) {
	return msdb.ReturnTrie, nil
}

func (msdb *MockStateDatabase) TrieDB() trie_wrapper.GethTrieDatabase {
	return msdb.ReturnDB.TrieDB()
}

// implements state.GethStateDatabase interface for testing
type mockStateDatabase struct {
}

func (*mockStateDatabase) ContractCode(addrHash, codeHash common.Hash) ([]byte, error) {
	panic("implement me")
}

func (*mockStateDatabase) ContractCodeSize(addrHash, codeHash common.Hash) (int, error) {
	panic("implement me")
}

func (*mockStateDatabase) CopyTrie(state.Trie) state.Trie {
	panic("implement me")
}

func (*mockStateDatabase) OpenStorageTrie(addrHash, root common.Hash) (state.Trie, error) {
	panic("implement me")
}

func (*mockStateDatabase) OpenTrie(root common.Hash) (state.Trie, error) {
	return &trie.SecureTrie{}, nil
}

func (*mockStateDatabase) TrieDB() *trie.Database {
	trieDB := trie.NewDatabase(&mockEthDB{})
	trieDB.InsertBlob(test_helpers.FakeHash, test_helpers.FakeTrieNode)
	return trieDB
}

// implements eth.GethStateDatabase interface for testing
type mockEthDB struct {
}

func (db *mockEthDB) NewIteratorWithStart(start []byte) ethdb.Iterator {
	panic("implement me")
}

func (*mockEthDB) Put(key []byte, value []byte) error {
	panic("implement me")
}

func (*mockEthDB) Get(key []byte) ([]byte, error) {
	return []byte{1, 2, 3, 4, 5}, nil
}

func (*mockEthDB) Has(key []byte) (bool, error) {
	panic("implement me")
}

func (*mockEthDB) Delete(key []byte) error {
	panic("implement me")
}

func (*mockEthDB) Close() error {
	panic("implement me")
}

func (*mockEthDB) NewBatch() ethdb.Batch {
	panic("implement me")
}

func (*mockEthDB) Compact([]byte, []byte) error {
	panic("implement me")
}

func (*mockEthDB) NewIterator() ethdb.Iterator {
	panic("implement me")
}

func (*mockEthDB) NewIteratorWithPrefix([]byte) ethdb.Iterator {
	panic("implement me")
}

func (*mockEthDB) Stat(string) (string, error) {
	panic("implement me")
}
