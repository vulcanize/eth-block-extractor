package fs_test

import (
	"github.com/8thlight/block_watcher/pkg/fs"
	"github.com/8thlight/block_watcher/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block file writer", func() {
	It("Writes block data to a file", func() {
		mockFileCreator := test_helpers.NewMockFileCreator()
		mockFileWriter := test_helpers.NewMockFileWriter()
		blockFileWriter := fs.NewBlockFileWriter(mockFileCreator, mockFileWriter)
		blockData := []byte{1, 2, 3, 4, 5}
		blockNumber := int64(67890)

		_, err := blockFileWriter.WriteBlockFile(blockData, blockNumber)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockFileCreator.Called).To(BeTrue())
		Expect(mockFileCreator.PassedName).To(Equal("blocks/block_67890.bytes"))
		Expect(mockFileWriter.Called).To(BeTrue())
		Expect(mockFileWriter.PassedFilename).To(Equal("blocks/block_67890.bytes"))
		Expect(mockFileWriter.PassedData).To(Equal(blockData))
	})
})
