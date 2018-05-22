package transformers_test

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gx/ipfs/QmVmDhyTTUcQXFD1rRQ64fGLMSAoaQvNH3hwuaCFAPq2hy/errors"

	"github.com/vulcanize/block_watcher/pkg/transformers"
	"github.com/vulcanize/block_watcher/test_helpers"
)

var _ = Describe("Ethereum state trie transformer", func() {
	BeforeEach(func() {
		log.SetOutput(ioutil.Discard)
	})

	It("fetches block header for block", func() {
		mockDatabase := test_helpers.NewMockDatabase()
		mockDatabase.SetHeaderByBlockNumberReturnBytes([][]byte{{1, 2, 3, 4, 5}})
		mockDecoder := test_helpers.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{})
		transformer := transformers.NewEthStateTrieTransformer(mockDatabase, mockDecoder, test_helpers.NewMockPublisher())

		err := transformer.Execute(0, 0)

		Expect(err).NotTo(HaveOccurred())
		mockDatabase.AssertGetBlockHeaderByBlockNumberCalledWith([]int64{0})
	})

	It("returns err if fetching block header returns err", func() {
		mockDatabase := test_helpers.NewMockDatabase()
		fakeError := errors.New("header fetching failed")
		mockDatabase.SetHeaderByBlockNumberError(fakeError)
		transformer := transformers.NewEthStateTrieTransformer(mockDatabase, test_helpers.NewMockDecoder(), test_helpers.NewMockPublisher())

		err := transformer.Execute(0, 0)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring(fakeError.Error()))
	})

	It("fetches state trie nodes with state root from decoded block header", func() {
		mockDatabase := test_helpers.NewMockDatabase()
		fakeHeader := []byte{6, 7, 8, 9, 0}
		mockDatabase.SetHeaderByBlockNumberReturnBytes([][]byte{fakeHeader})
		mockDecoder := test_helpers.NewMockDecoder()
		rootHash := common.HexToHash("0x12345")
		mockDecoder.SetReturnOut(&types.Header{Root: rootHash})
		transformer := transformers.NewEthStateTrieTransformer(mockDatabase, mockDecoder, test_helpers.NewMockPublisher())

		err := transformer.Execute(0, 0)

		Expect(err).NotTo(HaveOccurred())
		mockDecoder.AssertDecodeCalledWith(fakeHeader, &types.Header{})
		mockDatabase.AssertGetStateTrieNodesCalledWith(rootHash.Bytes())
	})

	It("returns err if fetching state trie returns err", func() {
		mockDatabase := test_helpers.NewMockDatabase()
		mockDatabase.SetHeaderByBlockNumberReturnBytes([][]byte{{1, 2, 3, 4, 5}})
		fakeError := errors.New("state trie fetching failed")
		mockDatabase.SetStateTrieNodesError(fakeError)
		mockDecoder := test_helpers.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{Root: common.Hash{}})
		transformer := transformers.NewEthStateTrieTransformer(mockDatabase, mockDecoder, test_helpers.NewMockPublisher())

		err := transformer.Execute(0, 0)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring(fakeError.Error()))
	})

	It("writes state trie nodes to ipfs", func() {
		mockDatabase := test_helpers.NewMockDatabase()
		mockDatabase.SetHeaderByBlockNumberReturnBytes([][]byte{{1, 2, 3, 4, 5}})
		fakeStateTrieNodes := [][]byte{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 0}}
		mockDatabase.SetStateTrieNodesReturnBytes(fakeStateTrieNodes)
		mockDecoder := test_helpers.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{})
		mockPublisher := test_helpers.NewMockPublisher()
		mockPublisher.SetReturnStrings([][]string{{"one"}, {"two"}})
		transformer := transformers.NewEthStateTrieTransformer(mockDatabase, mockDecoder, mockPublisher)

		err := transformer.Execute(0, 0)

		Expect(err).NotTo(HaveOccurred())
		mockPublisher.AssertWriteCalledWith(fakeStateTrieNodes)
	})
})