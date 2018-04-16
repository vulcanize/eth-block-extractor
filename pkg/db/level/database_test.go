package level_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/8thlight/block_watcher/pkg/db/level"
	"github.com/8thlight/block_watcher/test_helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("LevelDatabase", func() {
	It("invokes the level database reader to query for block data", func() {
		mockLevelDBReader := test_helpers.NewMockLevelDatabaseReader()
		db := level.NewLevelDatabase(mockLevelDBReader)
		block := core.Block{
			Hash:   "abcde",
			Number: 123456,
		}

		_, err := db.Get(block)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockLevelDBReader.Called).To(BeTrue())
	})

	It("converts block data to required format", func() {
		mockLevelDBReader := test_helpers.NewMockLevelDatabaseReader()
		db := level.NewLevelDatabase(mockLevelDBReader)
		hash := "abcde"
		num := int64(123456)
		block := core.Block{
			Hash:   hash,
			Number: num,
		}

		_, err := db.Get(block)

		Expect(err).NotTo(HaveOccurred())
		expectedHash := common.HexToHash(hash)
		Expect(mockLevelDBReader.PassedHash).To(Equal(expectedHash))
		expectedBlockNumber := uint64(num)
		Expect(mockLevelDBReader.PassedNumber).To(Equal(expectedBlockNumber))
	})
})
