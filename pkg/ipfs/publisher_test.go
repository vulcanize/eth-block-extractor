package ipfs_test

import (
	"github.com/8thlight/block_watcher/pkg/ipfs"
	"github.com/8thlight/block_watcher/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IPFS publisher", func() {
	It("writes block data to file", func() {
		mockBlockFileWriter := test_helpers.NewMockBlockFileWriter()
		mockIpfsWriter := test_helpers.NewMockIpfsWriter()
		publisher := ipfs.NewIpfsPublisher(mockBlockFileWriter, mockIpfsWriter)
		blockData := []byte{1, 2, 3, 4, 5}
		blockNumber := int64(67890)

		_, err := publisher.Write(blockData, blockNumber)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockBlockFileWriter.Called).To(BeTrue())
		Expect(mockBlockFileWriter.PassedBlockData).To(Equal(blockData))
		Expect(mockBlockFileWriter.PassedBlockNumber).To(Equal(blockNumber))
	})

	It("persists file to IPFS", func() {
		mockBlockFileWriter := test_helpers.NewMockBlockFileWriter()
		mockIpfsWriter := test_helpers.NewMockIpfsWriter()
		publisher := ipfs.NewIpfsPublisher(mockBlockFileWriter, mockIpfsWriter)
		filename := "filename"
		mockBlockFileWriter.SetReturnString(filename)

		_, err := publisher.Write([]byte{1, 2, 3, 4, 5}, int64(67890))

		Expect(err).NotTo(HaveOccurred())
		Expect(mockIpfsWriter.Called).To(BeTrue())
		Expect(mockIpfsWriter.PassedFilename).To(Equal(filename))
	})
})
