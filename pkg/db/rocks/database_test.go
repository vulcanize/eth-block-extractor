package rocks_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/8thlight/block_watcher/pkg/db/rocks"
	"github.com/8thlight/block_watcher/test_helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("Rocks database", func() {
	It("fetches block data from the reader with converted block hash", func() {
		decoder := test_helpers.NewMockDecoder()
		reader := test_helpers.NewMockRocksDatabaseReader()
		rocksDB := rocks.NewRocksDatabase(decoder, reader)
		hash := "0xHash"
		block := core.Block{
			Hash: hash,
		}

		_, err := rocksDB.Get(block)

		Expect(err).NotTo(HaveOccurred())
		Expect(reader.GetBlockHeaderCalled).To(BeTrue())
		expectedHash := common.FromHex(hash)
		Expect(reader.PassedKey).To(Equal(expectedHash))
	})

	It("returns error if reader returns error", func() {
		decoder := test_helpers.NewMockDecoder()
		reader := test_helpers.NewMockRocksDatabaseReader()
		fakeErr := errors.New("Failed")
		reader.SetError(fakeErr)
		rocksDB := rocks.NewRocksDatabase(decoder, reader)
		block := core.Block{}

		_, err := rocksDB.Get(block)

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(fakeErr))
	})

	It("decodes block data returned by the reader", func() {
		decoder := test_helpers.NewMockDecoder()
		reader := test_helpers.NewMockRocksDatabaseReader()
		bytesToDecode := []byte{1, 2, 3, 4, 5}
		reader.SetReturnBytes(bytesToDecode)
		rocksDB := rocks.NewRocksDatabase(decoder, reader)
		block := core.Block{}

		_, err := rocksDB.Get(block)

		Expect(err).NotTo(HaveOccurred())
		Expect(decoder.Called).To(BeTrue())
		Expect(decoder.PassedBytes).To(Equal(bytesToDecode))
	})
})
