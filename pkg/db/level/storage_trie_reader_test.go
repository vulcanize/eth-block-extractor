package level_test

import (
	"github.com/ethereum/go-ethereum/core/state"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/eth-block-extractor/pkg/db/level"
	"github.com/vulcanize/eth-block-extractor/test_helpers"
	state_wrapper "github.com/vulcanize/eth-block-extractor/test_helpers/mocks/wrappers/core/state"
	mock_rlp "github.com/vulcanize/eth-block-extractor/test_helpers/mocks/wrappers/rlp"
	"github.com/vulcanize/eth-block-extractor/test_helpers/mocks/wrappers/trie"
)

var _ = Describe("Storage trie reader", func() {
	It("decodes passed state trie leaf node into account", func() {
		db := state_wrapper.NewMockStateDatabase()
		trieDb := db.CreateFakeUnderlyingDatabase()
		db.ReturnDB = trieDb
		mockIteratror := trie.NewMockIterator(1)
		mockIteratror.SetIncludeLeaf()
		mockTrie := state_wrapper.NewMockTrie()
		mockTrie.SetReturnIterator(mockIteratror)
		db.ReturnTrie = mockTrie
		decoder := mock_rlp.NewMockDecoder()
		acct := &test_helpers.FakeStateAccount
		decoder.SetReturnOut(acct)
		reader := level.NewStorageTrieReader(db, decoder)

		_, err := reader.GetStorageTrie(test_helpers.FakeStateLeaf)

		Expect(err).NotTo(HaveOccurred())
		decoder.AssertDecodeCalledWith(test_helpers.FakeStateLeaf, &state.Account{})
	})

	It("fetches node associated with storage root", func() {
		db := state_wrapper.NewMockStateDatabase()
		trieDb := db.CreateFakeUnderlyingDatabase()
		db.ReturnDB = trieDb
		mockIteratror := trie.NewMockIterator(0)
		mockTrie := state_wrapper.NewMockTrie()
		mockTrie.SetReturnIterator(mockIteratror)
		db.ReturnTrie = mockTrie
		decoder := mock_rlp.NewMockDecoder()
		acct := &test_helpers.FakeStateAccount
		decoder.SetReturnOut(acct)
		reader := level.NewStorageTrieReader(db, decoder)

		storageTrieNodes, err := reader.GetStorageTrie(test_helpers.FakeStateLeaf)

		Expect(err).NotTo(HaveOccurred())
		Expect(len(storageTrieNodes)).To(Equal(1))
	})

	It("returns nodes found traversing storage trie", func() {
		db := state_wrapper.NewMockStateDatabase()
		trieDb := db.CreateFakeUnderlyingDatabase()
		db.ReturnDB = trieDb
		mockIteratror := trie.NewMockIterator(1)
		mockIteratror.SetIncludeLeaf()
		mockTrie := state_wrapper.NewMockTrie()
		mockTrie.SetReturnIterator(mockIteratror)
		db.ReturnTrie = mockTrie
		decoder := mock_rlp.NewMockDecoder()
		acct := &test_helpers.FakeStateAccount
		decoder.SetReturnOut(acct)
		reader := level.NewStorageTrieReader(db, decoder)

		storageTrieNodes, err := reader.GetStorageTrie(test_helpers.FakeStateLeaf)

		Expect(err).NotTo(HaveOccurred())
		Expect(len(storageTrieNodes)).To(Equal(2))
	})
})
