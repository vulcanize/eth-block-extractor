package rocks_test

import (
	"github.com/8thlight/block_watcher/pkg/db/rocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rocks RLP decoder", func() {
	It("returns error if input empty", func() {
		emptyInput := []byte{}
		decoder := rocks.EthBlockHeaderDecoder{}

		_, err := decoder.Decode(emptyInput)

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(rocks.NewDecodeError("No input to decode", nil)))
	})

	It("returns error if input does not begin with byte 249", func() {
		wrongHeaderInput := []byte{200, 1, 2, 3, 4}
		decoder := rocks.EthBlockHeaderDecoder{}

		_, err := decoder.Decode(wrongHeaderInput)

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(rocks.NewDecodeError("Received byte array not beginning with 249: 200", nil)))
	})

	It("returns err if input cannot be split as RLP data", func() {
		malformedInput := []byte{249, 1, 2, 3, 4}
		decoder := rocks.EthBlockHeaderDecoder{}

		_, err := decoder.Decode(malformedInput)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Error decoding header:"))
	})

	It("replaces uncles hash in ethereum block", func() {
		// block 5000002
		rawRocksBlock := []byte{249, 1, 250, 160, 5, 107, 244, 73, 195, 48, 48, 199, 75, 1, 205, 254, 44, 118, 224, 92, 37, 137, 83, 64, 182, 133, 117, 229, 18, 37, 229, 170, 220, 122, 246, 225, 129, 2, 148, 90, 11, 84, 213, 220, 23, 224, 170, 220, 56, 61, 45, 180, 59, 10, 13, 62, 2, 156, 76, 160, 193, 65, 87, 125, 76, 169, 59, 49, 39, 182, 2, 223, 140, 52, 182, 150, 42, 86, 24, 62, 183, 15, 101, 89, 179, 2, 167, 112, 17, 219, 124, 82, 160, 195, 247, 201, 21, 206, 212, 214, 47, 214, 141, 206, 201, 90, 209, 12, 234, 138, 241, 237, 172, 55, 239, 220, 57, 123, 78, 98, 232, 161, 225, 221, 132, 160, 11, 41, 115, 71, 71, 166, 169, 231, 139, 47, 255, 81, 171, 92, 248, 41, 38, 28, 8, 142, 43, 210, 36, 116, 250, 71, 23, 87, 175, 30, 167, 97, 185, 1, 0, 0, 18, 131, 8, 34, 72, 129, 81, 40, 137, 4, 0, 12, 5, 128, 0, 0, 8, 1, 20, 148, 0, 48, 0, 10, 192, 5, 0, 26, 36, 145, 106, 0, 8, 0, 177, 8, 16, 17, 33, 145, 176, 66, 132, 114, 0, 52, 16, 0, 1, 80, 4, 0, 133, 2, 24, 1, 32, 65, 0, 98, 160, 16, 6, 0, 10, 17, 192, 16, 24, 0, 155, 80, 162, 130, 10, 32, 40, 193, 226, 0, 48, 0, 0, 13, 33, 128, 80, 108, 1, 133, 72, 129, 0, 0, 8, 4, 32, 0, 40, 32, 2, 2, 66, 132, 65, 0, 1, 32, 42, 4, 1, 48, 32, 0, 33, 40, 140, 0, 68, 194, 0, 36, 17, 68, 128, 0, 228, 0, 32, 0, 6, 128, 92, 4, 0, 8, 2, 78, 6, 132, 64, 17, 1, 192, 128, 144, 17, 8, 9, 0, 164, 138, 64, 128, 129, 32, 0, 100, 0, 2, 68, 1, 10, 32, 196, 164, 0, 198, 0, 8, 24, 17, 64, 128, 148, 18, 6, 5, 1, 160, 66, 160, 1, 205, 0, 128, 16, 2, 0, 48, 0, 16, 3, 49, 2, 32, 13, 9, 0, 32, 45, 129, 4, 16, 156, 2, 1, 129, 19, 128, 167, 130, 11, 18, 138, 0, 16, 20, 20, 8, 2, 161, 193, 68, 146, 1, 148, 1, 131, 0, 1, 4, 64, 137, 68, 128, 64, 128, 0, 72, 16, 40, 16, 0, 144, 0, 91, 65, 100, 96, 0, 192, 0, 6, 192, 135, 9, 13, 67, 73, 169, 214, 215, 131, 76, 75, 66, 131, 121, 213, 68, 131, 121, 143, 207, 132, 90, 112, 118, 31, 152, 69, 84, 72, 46, 69, 84, 72, 70, 65, 78, 83, 46, 79, 82, 71, 45, 67, 68, 70, 57, 49, 53, 52, 54, 160, 2, 138, 25, 214, 220, 240, 151, 89, 114, 40, 95, 113, 162, 63, 227, 164, 104, 136, 106, 210, 76, 198, 63, 150, 145, 88, 212, 39, 85, 109, 107, 181, 136, 201, 160, 68, 32, 29, 217, 152, 242}
		matchingLevelBlock := []byte{249, 2, 25, 160, 5, 107, 244, 73, 195, 48, 48, 199, 75, 1, 205, 254, 44, 118, 224, 92, 37, 137, 83, 64, 182, 133, 117, 229, 18, 37, 229, 170, 220, 122, 246, 225, 160, 29, 204, 77, 232, 222, 199, 93, 122, 171, 133, 181, 103, 182, 204, 212, 26, 211, 18, 69, 27, 148, 138, 116, 19, 240, 161, 66, 253, 64, 212, 147, 71, 148, 90, 11, 84, 213, 220, 23, 224, 170, 220, 56, 61, 45, 180, 59, 10, 13, 62, 2, 156, 76, 160, 193, 65, 87, 125, 76, 169, 59, 49, 39, 182, 2, 223, 140, 52, 182, 150, 42, 86, 24, 62, 183, 15, 101, 89, 179, 2, 167, 112, 17, 219, 124, 82, 160, 195, 247, 201, 21, 206, 212, 214, 47, 214, 141, 206, 201, 90, 209, 12, 234, 138, 241, 237, 172, 55, 239, 220, 57, 123, 78, 98, 232, 161, 225, 221, 132, 160, 11, 41, 115, 71, 71, 166, 169, 231, 139, 47, 255, 81, 171, 92, 248, 41, 38, 28, 8, 142, 43, 210, 36, 116, 250, 71, 23, 87, 175, 30, 167, 97, 185, 1, 0, 0, 18, 131, 8, 34, 72, 129, 81, 40, 137, 4, 0, 12, 5, 128, 0, 0, 8, 1, 20, 148, 0, 48, 0, 10, 192, 5, 0, 26, 36, 145, 106, 0, 8, 0, 177, 8, 16, 17, 33, 145, 176, 66, 132, 114, 0, 52, 16, 0, 1, 80, 4, 0, 133, 2, 24, 1, 32, 65, 0, 98, 160, 16, 6, 0, 10, 17, 192, 16, 24, 0, 155, 80, 162, 130, 10, 32, 40, 193, 226, 0, 48, 0, 0, 13, 33, 128, 80, 108, 1, 133, 72, 129, 0, 0, 8, 4, 32, 0, 40, 32, 2, 2, 66, 132, 65, 0, 1, 32, 42, 4, 1, 48, 32, 0, 33, 40, 140, 0, 68, 194, 0, 36, 17, 68, 128, 0, 228, 0, 32, 0, 6, 128, 92, 4, 0, 8, 2, 78, 6, 132, 64, 17, 1, 192, 128, 144, 17, 8, 9, 0, 164, 138, 64, 128, 129, 32, 0, 100, 0, 2, 68, 1, 10, 32, 196, 164, 0, 198, 0, 8, 24, 17, 64, 128, 148, 18, 6, 5, 1, 160, 66, 160, 1, 205, 0, 128, 16, 2, 0, 48, 0, 16, 3, 49, 2, 32, 13, 9, 0, 32, 45, 129, 4, 16, 156, 2, 1, 129, 19, 128, 167, 130, 11, 18, 138, 0, 16, 20, 20, 8, 2, 161, 193, 68, 146, 1, 148, 1, 131, 0, 1, 4, 64, 137, 68, 128, 64, 128, 0, 72, 16, 40, 16, 0, 144, 0, 91, 65, 100, 96, 0, 192, 0, 6, 192, 135, 9, 13, 67, 73, 169, 214, 215, 131, 76, 75, 66, 131, 121, 213, 68, 131, 121, 143, 207, 132, 90, 112, 118, 31, 152, 69, 84, 72, 46, 69, 84, 72, 70, 65, 78, 83, 46, 79, 82, 71, 45, 67, 68, 70, 57, 49, 53, 52, 54, 160, 2, 138, 25, 214, 220, 240, 151, 89, 114, 40, 95, 113, 162, 63, 227, 164, 104, 136, 106, 210, 76, 198, 63, 150, 145, 88, 212, 39, 85, 109, 107, 181, 136, 201, 160, 68, 32, 29, 217, 152, 242}

		decoder := rocks.EthBlockHeaderDecoder{}

		result, err := decoder.Decode(rawRocksBlock)

		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(matchingLevelBlock))

	})

	It("works for another block", func() {
		// block 5350025
		rawRocksBlock := []byte{249, 1, 228, 160, 177, 179, 134, 146, 143, 123, 230, 168, 49, 88, 47, 153, 7, 142, 9, 227, 168, 70, 10, 3, 112, 63, 73, 59, 190, 132, 145, 97, 204, 188, 239, 22, 129, 2, 148, 178, 147, 11, 53, 132, 74, 35, 15, 0, 229, 20, 49, 172, 174, 150, 254, 84, 58, 3, 71, 160, 116, 117, 200, 28, 26, 78, 100, 82, 134, 188, 201, 56, 219, 125, 146, 11, 238, 130, 226, 247, 74, 191, 1, 204, 74, 244, 83, 4, 76, 212, 181, 250, 160, 189, 128, 4, 234, 182, 89, 153, 180, 200, 36, 70, 249, 121, 18, 95, 64, 182, 64, 37, 144, 34, 29, 201, 152, 248, 226, 195, 119, 148, 31, 70, 27, 160, 164, 45, 134, 79, 66, 230, 65, 243, 130, 2, 57, 253, 138, 74, 224, 252, 107, 208, 110, 196, 37, 221, 142, 61, 87, 105, 255, 134, 106, 125, 204, 124, 185, 1, 0, 0, 0, 17, 161, 0, 72, 33, 0, 128, 35, 162, 0, 0, 4, 161, 2, 0, 3, 0, 47, 128, 16, 40, 8, 72, 129, 134, 160, 0, 44, 3, 98, 16, 36, 34, 48, 16, 0, 1, 193, 128, 64, 16, 2, 136, 16, 0, 0, 18, 0, 0, 0, 0, 1, 8, 144, 16, 0, 8, 4, 8, 32, 32, 8, 0, 64, 2, 192, 24, 0, 9, 17, 129, 64, 0, 8, 136, 129, 2, 4, 0, 0, 0, 32, 0, 8, 128, 0, 0, 16, 0, 5, 8, 128, 128, 0, 130, 32, 0, 40, 3, 32, 8, 64, 0, 80, 48, 144, 0, 0, 8, 16, 36, 32, 64, 0, 8, 140, 4, 64, 16, 32, 2, 17, 16, 4, 32, 1, 0, 160, 0, 34, 24, 0, 0, 68, 4, 66, 65, 137, 1, 32, 128, 8, 10, 38, 132, 0, 32, 16, 64, 5, 8, 0, 0, 0, 32, 22, 52, 64, 74, 12, 9, 9, 128, 192, 80, 0, 4, 0, 0, 128, 0, 0, 146, 6, 34, 1, 135, 64, 32, 2, 0, 67, 144, 2, 128, 64, 176, 0, 4, 2, 17, 64, 32, 2, 128, 0, 2, 2, 0, 16, 1, 0, 2, 129, 64, 32, 20, 10, 0, 1, 4, 10, 64, 9, 64, 0, 0, 0, 0, 0, 32, 16, 68, 16, 33, 1, 4, 16, 0, 0, 8, 20, 0, 128, 0, 1, 172, 36, 1, 104, 0, 0, 128, 64, 32, 4, 0, 4, 128, 132, 1, 0, 0, 16, 135, 11, 167, 186, 160, 182, 78, 10, 131, 81, 162, 137, 131, 122, 18, 29, 131, 122, 2, 183, 132, 90, 190, 109, 40, 130, 108, 49, 160, 42, 56, 46, 147, 159, 181, 193, 238, 25, 34, 247, 150, 128, 12, 202, 140, 47, 124, 196, 94, 170, 44, 20, 53, 232, 48, 49, 116, 92, 135, 19, 30, 136, 227, 251, 169, 172, 6, 106, 193, 153}
		matchingLevelBlock := []byte{249, 2, 3, 160, 177, 179, 134, 146, 143, 123, 230, 168, 49, 88, 47, 153, 7, 142, 9, 227, 168, 70, 10, 3, 112, 63, 73, 59, 190, 132, 145, 97, 204, 188, 239, 22, 160, 29, 204, 77, 232, 222, 199, 93, 122, 171, 133, 181, 103, 182, 204, 212, 26, 211, 18, 69, 27, 148, 138, 116, 19, 240, 161, 66, 253, 64, 212, 147, 71, 148, 178, 147, 11, 53, 132, 74, 35, 15, 0, 229, 20, 49, 172, 174, 150, 254, 84, 58, 3, 71, 160, 116, 117, 200, 28, 26, 78, 100, 82, 134, 188, 201, 56, 219, 125, 146, 11, 238, 130, 226, 247, 74, 191, 1, 204, 74, 244, 83, 4, 76, 212, 181, 250, 160, 189, 128, 4, 234, 182, 89, 153, 180, 200, 36, 70, 249, 121, 18, 95, 64, 182, 64, 37, 144, 34, 29, 201, 152, 248, 226, 195, 119, 148, 31, 70, 27, 160, 164, 45, 134, 79, 66, 230, 65, 243, 130, 2, 57, 253, 138, 74, 224, 252, 107, 208, 110, 196, 37, 221, 142, 61, 87, 105, 255, 134, 106, 125, 204, 124, 185, 1, 0, 0, 0, 17, 161, 0, 72, 33, 0, 128, 35, 162, 0, 0, 4, 161, 2, 0, 3, 0, 47, 128, 16, 40, 8, 72, 129, 134, 160, 0, 44, 3, 98, 16, 36, 34, 48, 16, 0, 1, 193, 128, 64, 16, 2, 136, 16, 0, 0, 18, 0, 0, 0, 0, 1, 8, 144, 16, 0, 8, 4, 8, 32, 32, 8, 0, 64, 2, 192, 24, 0, 9, 17, 129, 64, 0, 8, 136, 129, 2, 4, 0, 0, 0, 32, 0, 8, 128, 0, 0, 16, 0, 5, 8, 128, 128, 0, 130, 32, 0, 40, 3, 32, 8, 64, 0, 80, 48, 144, 0, 0, 8, 16, 36, 32, 64, 0, 8, 140, 4, 64, 16, 32, 2, 17, 16, 4, 32, 1, 0, 160, 0, 34, 24, 0, 0, 68, 4, 66, 65, 137, 1, 32, 128, 8, 10, 38, 132, 0, 32, 16, 64, 5, 8, 0, 0, 0, 32, 22, 52, 64, 74, 12, 9, 9, 128, 192, 80, 0, 4, 0, 0, 128, 0, 0, 146, 6, 34, 1, 135, 64, 32, 2, 0, 67, 144, 2, 128, 64, 176, 0, 4, 2, 17, 64, 32, 2, 128, 0, 2, 2, 0, 16, 1, 0, 2, 129, 64, 32, 20, 10, 0, 1, 4, 10, 64, 9, 64, 0, 0, 0, 0, 0, 32, 16, 68, 16, 33, 1, 4, 16, 0, 0, 8, 20, 0, 128, 0, 1, 172, 36, 1, 104, 0, 0, 128, 64, 32, 4, 0, 4, 128, 132, 1, 0, 0, 16, 135, 11, 167, 186, 160, 182, 78, 10, 131, 81, 162, 137, 131, 122, 18, 29, 131, 122, 2, 183, 132, 90, 190, 109, 40, 130, 108, 49, 160, 42, 56, 46, 147, 159, 181, 193, 238, 25, 34, 247, 150, 128, 12, 202, 140, 47, 124, 196, 94, 170, 44, 20, 53, 232, 48, 49, 116, 92, 135, 19, 30, 136, 227, 251, 169, 172, 6, 106, 193, 153}
		decoder := rocks.EthBlockHeaderDecoder{}

		result, err := decoder.Decode(rawRocksBlock)

		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(matchingLevelBlock))
	})

	It("replaces transaction root and receipts root in ethereum block", func() {
		// block 1234571
		rawRocksBlock := []byte{249, 1, 215, 160, 133, 36, 75, 179, 255, 39, 30, 80, 157, 6, 231, 63, 74, 157, 51, 139, 19, 197, 245, 200, 147, 216, 112, 202, 210, 160, 108, 7, 37, 93, 150, 44, 160, 225, 251, 49, 240, 63, 105, 125, 51, 39, 42, 246, 139, 85, 31, 29, 204, 61, 36, 120, 183, 176, 224, 112, 131, 224, 195, 8, 65, 77, 172, 136, 64, 148, 21, 18, 85, 221, 158, 56, 228, 77, 179, 142, 160, 110, 198, 109, 13, 17, 61, 108, 190, 55, 160, 65, 38, 223, 208, 132, 66, 107, 82, 154, 101, 243, 120, 196, 126, 236, 141, 75, 95, 243, 89, 228, 205, 196, 131, 222, 73, 102, 98, 82, 246, 22, 228, 129, 0, 129, 0, 185, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 134, 20, 78, 150, 37, 74, 14, 131, 18, 214, 139, 131, 71, 231, 196, 128, 132, 86, 249, 137, 239, 152, 215, 131, 1, 3, 5, 132, 71, 101, 116, 104, 135, 103, 111, 49, 46, 53, 46, 49, 133, 108, 105, 110, 117, 120, 160, 3, 243, 173, 181, 223, 76, 79, 216, 126, 35, 230, 139, 1, 34, 228, 168, 223, 251, 77, 252, 114, 83, 220, 163, 77, 233, 214, 115, 75, 95, 150, 191, 136, 143, 243, 191, 142, 116, 216, 20, 72}
		matchingLevelBlock := []byte{249, 2, 21, 160, 133, 36, 75, 179, 255, 39, 30, 80, 157, 6, 231, 63, 74, 157, 51, 139, 19, 197, 245, 200, 147, 216, 112, 202, 210, 160, 108, 7, 37, 93, 150, 44, 160, 225, 251, 49, 240, 63, 105, 125, 51, 39, 42, 246, 139, 85, 31, 29, 204, 61, 36, 120, 183, 176, 224, 112, 131, 224, 195, 8, 65, 77, 172, 136, 64, 148, 21, 18, 85, 221, 158, 56, 228, 77, 179, 142, 160, 110, 198, 109, 13, 17, 61, 108, 190, 55, 160, 65, 38, 223, 208, 132, 66, 107, 82, 154, 101, 243, 120, 196, 126, 236, 141, 75, 95, 243, 89, 228, 205, 196, 131, 222, 73, 102, 98, 82, 246, 22, 228, 160, 86, 232, 31, 23, 27, 204, 85, 166, 255, 131, 69, 230, 146, 192, 248, 110, 91, 72, 224, 27, 153, 108, 173, 192, 1, 98, 47, 181, 227, 99, 180, 33, 160, 86, 232, 31, 23, 27, 204, 85, 166, 255, 131, 69, 230, 146, 192, 248, 110, 91, 72, 224, 27, 153, 108, 173, 192, 1, 98, 47, 181, 227, 99, 180, 33, 185, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 134, 20, 78, 150, 37, 74, 14, 131, 18, 214, 139, 131, 71, 231, 196, 128, 132, 86, 249, 137, 239, 152, 215, 131, 1, 3, 5, 132, 71, 101, 116, 104, 135, 103, 111, 49, 46, 53, 46, 49, 133, 108, 105, 110, 117, 120, 160, 3, 243, 173, 181, 223, 76, 79, 216, 126, 35, 230, 139, 1, 34, 228, 168, 223, 251, 77, 252, 114, 83, 220, 163, 77, 233, 214, 115, 75, 95, 150, 191, 136, 143, 243, 191, 142, 116, 216, 20, 72}
		decoder := rocks.EthBlockHeaderDecoder{}

		result, err := decoder.Decode(rawRocksBlock)

		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(matchingLevelBlock))
	})

	It("works for another block", func() {
		// block 1234575
		rawRocksBlock := []byte{249, 1, 248, 160, 199, 172, 103, 167, 55, 246, 146, 140, 46, 65, 122, 63, 85, 43, 38, 4, 210, 179, 139, 74, 15, 33, 239, 214, 109, 187, 130, 131, 78, 56, 212, 29, 129, 2, 148, 115, 141, 183, 20, 192, 139, 138, 50, 162, 158, 14, 104, 175, 0, 33, 80, 121, 170, 156, 92, 160, 127, 128, 94, 207, 180, 25, 138, 89, 113, 141, 33, 193, 213, 207, 190, 138, 169, 230, 172, 200, 112, 229, 26, 191, 181, 147, 158, 41, 33, 2, 157, 106, 160, 223, 44, 171, 177, 118, 108, 159, 167, 76, 2, 143, 181, 163, 134, 126, 122, 177, 3, 117, 39, 32, 123, 94, 199, 62, 187, 122, 140, 231, 79, 162, 68, 160, 154, 171, 151, 108, 57, 187, 218, 7, 175, 100, 30, 145, 156, 108, 66, 156, 232, 11, 203, 219, 59, 75, 192, 181, 161, 18, 243, 100, 127, 192, 59, 92, 185, 1, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 128, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 134, 20, 83, 170, 28, 29, 185, 131, 18, 214, 143, 131, 71, 231, 196, 130, 253, 132, 132, 86, 249, 138, 18, 152, 215, 131, 1, 3, 5, 132, 71, 101, 116, 104, 135, 103, 111, 49, 46, 53, 46, 49, 133, 108, 105, 110, 117, 120, 160, 22, 196, 33, 18, 113, 99, 69, 100, 60, 108, 101, 160, 82, 136, 7, 202, 230, 194, 79, 43, 230, 202, 212, 119, 46, 219, 10, 29, 220, 220, 36, 252, 136, 239, 68, 31, 52, 98, 54, 65, 174}
		matchingLevelBlock := []byte{249, 2, 23, 160, 199, 172, 103, 167, 55, 246, 146, 140, 46, 65, 122, 63, 85, 43, 38, 4, 210, 179, 139, 74, 15, 33, 239, 214, 109, 187, 130, 131, 78, 56, 212, 29, 160, 29, 204, 77, 232, 222, 199, 93, 122, 171, 133, 181, 103, 182, 204, 212, 26, 211, 18, 69, 27, 148, 138, 116, 19, 240, 161, 66, 253, 64, 212, 147, 71, 148, 115, 141, 183, 20, 192, 139, 138, 50, 162, 158, 14, 104, 175, 0, 33, 80, 121, 170, 156, 92, 160, 127, 128, 94, 207, 180, 25, 138, 89, 113, 141, 33, 193, 213, 207, 190, 138, 169, 230, 172, 200, 112, 229, 26, 191, 181, 147, 158, 41, 33, 2, 157, 106, 160, 223, 44, 171, 177, 118, 108, 159, 167, 76, 2, 143, 181, 163, 134, 126, 122, 177, 3, 117, 39, 32, 123, 94, 199, 62, 187, 122, 140, 231, 79, 162, 68, 160, 154, 171, 151, 108, 57, 187, 218, 7, 175, 100, 30, 145, 156, 108, 66, 156, 232, 11, 203, 219, 59, 75, 192, 181, 161, 18, 243, 100, 127, 192, 59, 92, 185, 1, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 128, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 134, 20, 83, 170, 28, 29, 185, 131, 18, 214, 143, 131, 71, 231, 196, 130, 253, 132, 132, 86, 249, 138, 18, 152, 215, 131, 1, 3, 5, 132, 71, 101, 116, 104, 135, 103, 111, 49, 46, 53, 46, 49, 133, 108, 105, 110, 117, 120, 160, 22, 196, 33, 18, 113, 99, 69, 100, 60, 108, 101, 160, 82, 136, 7, 202, 230, 194, 79, 43, 230, 202, 212, 119, 46, 219, 10, 29, 220, 220, 36, 252, 136, 239, 68, 31, 52, 98, 54, 65, 174}
		decoder := rocks.EthBlockHeaderDecoder{}

		result, err := decoder.Decode(rawRocksBlock)

		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(matchingLevelBlock))
	})

	It("is a no-op on a properly constructed level block", func() {
		levelBlock := []byte{249, 2, 23, 160, 199, 172, 103, 167, 55, 246, 146, 140, 46, 65, 122, 63, 85, 43, 38, 4, 210, 179, 139, 74, 15, 33, 239, 214, 109, 187, 130, 131, 78, 56, 212, 29, 160, 29, 204, 77, 232, 222, 199, 93, 122, 171, 133, 181, 103, 182, 204, 212, 26, 211, 18, 69, 27, 148, 138, 116, 19, 240, 161, 66, 253, 64, 212, 147, 71, 148, 115, 141, 183, 20, 192, 139, 138, 50, 162, 158, 14, 104, 175, 0, 33, 80, 121, 170, 156, 92, 160, 127, 128, 94, 207, 180, 25, 138, 89, 113, 141, 33, 193, 213, 207, 190, 138, 169, 230, 172, 200, 112, 229, 26, 191, 181, 147, 158, 41, 33, 2, 157, 106, 160, 223, 44, 171, 177, 118, 108, 159, 167, 76, 2, 143, 181, 163, 134, 126, 122, 177, 3, 117, 39, 32, 123, 94, 199, 62, 187, 122, 140, 231, 79, 162, 68, 160, 154, 171, 151, 108, 57, 187, 218, 7, 175, 100, 30, 145, 156, 108, 66, 156, 232, 11, 203, 219, 59, 75, 192, 181, 161, 18, 243, 100, 127, 192, 59, 92, 185, 1, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 128, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 134, 20, 83, 170, 28, 29, 185, 131, 18, 214, 143, 131, 71, 231, 196, 130, 253, 132, 132, 86, 249, 138, 18, 152, 215, 131, 1, 3, 5, 132, 71, 101, 116, 104, 135, 103, 111, 49, 46, 53, 46, 49, 133, 108, 105, 110, 117, 120, 160, 22, 196, 33, 18, 113, 99, 69, 100, 60, 108, 101, 160, 82, 136, 7, 202, 230, 194, 79, 43, 230, 202, 212, 119, 46, 219, 10, 29, 220, 220, 36, 252, 136, 239, 68, 31, 52, 98, 54, 65, 174}
		decoder := rocks.EthBlockHeaderDecoder{}

		result, err := decoder.Decode(levelBlock)

		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(levelBlock))
	})
})
