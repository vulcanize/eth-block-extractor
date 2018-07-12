package ipfs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ipfs2 "github.com/vulcanize/eth-block-extractor/pkg/ipfs"
	"github.com/vulcanize/eth-block-extractor/test_helpers"
	"github.com/vulcanize/eth-block-extractor/test_helpers/mocks/ipfs"
	"github.com/vulcanize/eth-block-extractor/test_helpers/mocks/wrappers/go-cid"
)

var _ = Describe("Ipld reader", func() {
	It("decodes passed string into cid", func() {
		decoder := go_cid.NewMockCidDecoder()
		getter := ipfs.NewMockGetter()
		resolver := ipfs.NewMockResolver()
		reader := ipfs2.NewIpldReader(decoder, getter, resolver)

		_, err := reader.Read("cid")

		Expect(err).NotTo(HaveOccurred())
		decoder.AssertDecodeCalledWith("cid")
	})

	It("returns error if decoder returns err", func() {
		decoder := go_cid.NewMockCidDecoder()
		decoder.SetDecodeError(test_helpers.FakeError)
		getter := ipfs.NewMockGetter()
		resolver := ipfs.NewMockResolver()
		reader := ipfs2.NewIpldReader(decoder, getter, resolver)

		_, err := reader.Read("cid")

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})

	It("gets block corresponding to decoded cid", func() {
		decoder := go_cid.NewMockCidDecoder()
		decoder.SetDecodeReturnCid(test_helpers.FakeCid)
		getter := ipfs.NewMockGetter()
		resolver := ipfs.NewMockResolver()
		reader := ipfs2.NewIpldReader(decoder, getter, resolver)

		_, err := reader.Read("cid")

		Expect(err).NotTo(HaveOccurred())
		getter.AssertGetCalledWith(test_helpers.FakeCid)
	})

	It("returns error if getting block returns err", func() {
		decoder := go_cid.NewMockCidDecoder()
		getter := ipfs.NewMockGetter()
		getter.SetReturnErr(test_helpers.FakeError)
		resolver := ipfs.NewMockResolver()
		reader := ipfs2.NewIpldReader(decoder, getter, resolver)

		_, err := reader.Read("cid")

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})

	It("resolves block into IPLD node", func() {
		decoder := go_cid.NewMockCidDecoder()
		getter := ipfs.NewMockGetter()
		getter.SetReturnBlock(test_helpers.FakeIpfsBlock)
		resolver := ipfs.NewMockResolver()
		reader := ipfs2.NewIpldReader(decoder, getter, resolver)

		_, err := reader.Read("cid")

		Expect(err).NotTo(HaveOccurred())
		resolver.AssertResolveCalledWith(test_helpers.FakeIpfsBlock)
	})

	It("returns error if resolving block returns err", func() {
		decoder := go_cid.NewMockCidDecoder()
		getter := ipfs.NewMockGetter()
		resolver := ipfs.NewMockResolver()
		resolver.SetReturnErr(test_helpers.FakeError)
		reader := ipfs2.NewIpldReader(decoder, getter, resolver)

		_, err := reader.Read("cid")

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})
})
