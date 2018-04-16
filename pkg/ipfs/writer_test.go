package ipfs_test

import (
	"github.com/8thlight/block_watcher/pkg/ipfs"
	"github.com/8thlight/block_watcher/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ipfs writer", func() {
	It("Executes command to put block data to ipfs as a dag", func() {
		fakeCommander := test_helpers.NewMockCommander()
		writer := ipfs.NewIpfsEthBlockWriter(fakeCommander)
		filename := "filename"

		_, err := writer.WriteToIpfs(filename)

		Expect(err).NotTo(HaveOccurred())
		Expect(fakeCommander.Called).To(BeTrue())
		Expect(fakeCommander.Name).To(Equal("ipfs"))
		Expect(fakeCommander.Args).To(ConsistOf("dag", "put", "--input-enc", "raw", "--format", "eth-block", filename))
	})
})
