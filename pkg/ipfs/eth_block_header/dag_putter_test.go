package eth_block_header_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"github.com/8thlight/block_watcher/pkg/ipfs/eth_block_header"
	"github.com/8thlight/block_watcher/test_helpers"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = Describe("Creating an IPLD for a block header", func() {
	It("decodes passed bytes into ethereum block header", func() {
		mockDecoder := test_helpers.NewMockDecoder()
		dagPutter := eth_block_header.NewBlockHeaderDagPutter(test_helpers.NewMockAdder(), mockDecoder)
		fakeBytes := []byte{1, 2, 3, 4, 5}

		_, err := dagPutter.DagPut(fakeBytes)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockDecoder.Called).To(BeTrue())
		Expect(mockDecoder.PassedBytes).To(Equal(fakeBytes))
		Expect(mockDecoder.PassedOut).To(BeAssignableToTypeOf(&types.Header{}))
	})

	It("returns error if decoding fails", func() {
		mockDecoder := test_helpers.NewMockDecoder()
		fakeError := errors.New("Failed")
		mockDecoder.SetError(fakeError)
		dagPutter := eth_block_header.NewBlockHeaderDagPutter(test_helpers.NewMockAdder(), mockDecoder)

		_, err := dagPutter.DagPut([]byte{1, 2, 3, 4, 5})

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(fakeError))
	})

	It("adds ethereum block header to ipfs", func() {
		mockAdder := test_helpers.NewMockAdder()
		dagPutter := eth_block_header.NewBlockHeaderDagPutter(mockAdder, test_helpers.NewMockDecoder())
		fakeBytes := []byte{1, 2, 3, 4, 5}

		_, err := dagPutter.DagPut(fakeBytes)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockAdder.Called).To(BeTrue())
		Expect(mockAdder.PassedNode).To(BeAssignableToTypeOf(&eth_block_header.EthBlockHeaderNode{}))
	})

	It("returns error if adding to ipfs fails", func() {
		mockAdder := test_helpers.NewMockAdder()
		fakeError := errors.New("Failed")
		mockAdder.SetError(fakeError)
		dagPutter := eth_block_header.NewBlockHeaderDagPutter(mockAdder, test_helpers.NewMockDecoder())

		_, err := dagPutter.DagPut([]byte{1, 2, 3, 4, 5})

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(fakeError))
	})
})
