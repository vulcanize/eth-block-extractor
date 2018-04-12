package pkg_test

import (
	"errors"
	"github.com/8thlight/block_watcher/pkg"
	"github.com/8thlight/block_watcher/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
)

var _ = Describe("Transformer", func() {
	var mockDatabase *test_helpers.MockDatabase
	var mockPublisher *test_helpers.MockPublisher

	Describe("Fetching one block", func() {
		var fakeBytes []byte
		var blockNumber int64

		BeforeEach(func() {
			mockDatabase = test_helpers.NewMockDatabase()
			mockPublisher = test_helpers.NewMockPublisher()
			blockNumber = 54321
			fakeBytes = []byte{6, 7, 8, 9, 0}
			mockDatabase.SetReturnBytes([][]byte{fakeBytes})
			mockPublisher.SetReturnBytes([][]byte{{0, 9, 8, 7, 6}})
			log.SetOutput(ioutil.Discard)
		})

		It("Fetches block RLP data from ethereum db", func() {
			transformer := pkg.NewTransformer(mockDatabase, mockPublisher)

			err := transformer.Execute(blockNumber, blockNumber)

			Expect(err).NotTo(HaveOccurred())
			Expect(mockDatabase.CalledCount).To(Equal(1))
			Expect(mockDatabase.PassedBlocks).To(ConsistOf(blockNumber))
		})

		It("Returns error if fetching block RLP data from ethereum DB fails", func() {
			fakeError := errors.New("Failed")
			mockDatabase.SetError(fakeError)
			transformer := pkg.NewTransformer(mockDatabase, mockPublisher)

			err := transformer.Execute(blockNumber, blockNumber)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(pkg.NewExecuteError(pkg.GetBlockRlpErr, fakeError)))
			Expect(mockDatabase.CalledCount).To(Equal(1))
			Expect(mockPublisher.CalledCount).To(BeZero())
		})

		It("Persists block RLP data to IPFS", func() {
			transformer := pkg.NewTransformer(mockDatabase, mockPublisher)

			err := transformer.Execute(blockNumber, blockNumber)

			Expect(err).NotTo(HaveOccurred())
			Expect(mockPublisher.CalledCount).To(Equal(1))
			Expect(mockPublisher.PassedBlockDatas[0]).To(Equal(fakeBytes))
		})

		It("Returns err if persisting block RLP data to IPFS fails", func() {
			fakeError := errors.New("Failed")
			mockPublisher.SetError(fakeError)
			transformer := pkg.NewTransformer(mockDatabase, mockPublisher)

			err := transformer.Execute(blockNumber, blockNumber)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(pkg.NewExecuteError(pkg.PutIpldErr, fakeError)))
			Expect(mockDatabase.CalledCount).To(Equal(1))
			Expect(mockPublisher.CalledCount).To(Equal(1))
		})
	})

	Describe("Fetching multiple blocks", func() {
		var fakeRlpBytes []byte
		var allFakeRlpBytes [][]byte
		var fakeOutputBytes []byte
		var startingBlockNumber int64
		var endingBlockNumber int64
		var blockNumbers []int64
		var numBlocks int

		BeforeEach(func() {
			mockDatabase = test_helpers.NewMockDatabase()
			mockPublisher = test_helpers.NewMockPublisher()
			startingBlockNumber = 54321
			endingBlockNumber = 54326
			numBlocks = int(endingBlockNumber) - int(startingBlockNumber) + 1
			for i := startingBlockNumber; i <= endingBlockNumber; i++ {
				blockNumbers = append(blockNumbers, i)
			}
			fakeRlpBytes = []byte{6, 7, 8, 9, 0}
			allFakeRlpBytes = [][]byte{fakeRlpBytes, fakeRlpBytes, fakeRlpBytes, fakeRlpBytes, fakeRlpBytes, fakeRlpBytes}
			mockDatabase.SetReturnBytes(allFakeRlpBytes)
			fakeOutputBytes = []byte{0, 9, 8, 7, 6}
			mockPublisher.SetReturnBytes([][]byte{fakeOutputBytes, fakeOutputBytes, fakeOutputBytes, fakeOutputBytes, fakeOutputBytes, fakeOutputBytes})
			log.SetOutput(ioutil.Discard)
		})

		It("Fetches block RLP data from ethereum db for every block in range", func() {
			transformer := pkg.NewTransformer(mockDatabase, mockPublisher)

			err := transformer.Execute(startingBlockNumber, endingBlockNumber)

			Expect(err).NotTo(HaveOccurred())
			Expect(mockDatabase.CalledCount).To(Equal(numBlocks))
			Expect(mockDatabase.PassedBlocks).To(ConsistOf(blockNumbers))
		})

		It("Persists block RLP data to IPFS for every block in range", func() {
			transformer := pkg.NewTransformer(mockDatabase, mockPublisher)

			err := transformer.Execute(startingBlockNumber, endingBlockNumber)

			Expect(err).NotTo(HaveOccurred())
			Expect(mockPublisher.CalledCount).To(Equal(numBlocks))
			Expect(mockPublisher.PassedBlockDatas).To(ConsistOf(allFakeRlpBytes))
		})
	})

})
