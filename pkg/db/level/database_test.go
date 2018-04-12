package level_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/8thlight/block_watcher/pkg/db/level"
	"github.com/8thlight/block_watcher/test_helpers"
	"github.com/ethereum/go-ethereum/common"
)

var _ = Describe("LevelDatabase", func() {
	It("invokes the level database reader to query for block data", func() {
		mockLevelDBReader := test_helpers.NewMockLevelDatabaseReader()
		db := level.NewLevelDatabase(mockLevelDBReader)

		_, err := db.Get(12345)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockLevelDBReader.Called).To(BeTrue())
	})

	It("converts block data to required format", func() {
		mockLevelDBReader := test_helpers.NewMockLevelDatabaseReader()
		hash := common.HexToHash("abcde")
		mockLevelDBReader.SetReturnHash(hash)
		db := level.NewLevelDatabase(mockLevelDBReader)
		num := int64(123456)

		_, err := db.Get(num)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockLevelDBReader.PassedHash).To(Equal(hash))
		expectedBlockNumber := uint64(num)
		Expect(mockLevelDBReader.PassedNumber).To(Equal(expectedBlockNumber))
	})
})
