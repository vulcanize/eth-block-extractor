package test_helpers

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

var (
	FakeError        = errors.New("failed")
	FakeHash         = common.HexToHash("0x123")
	FakeString       = "Test"
	FakeTrieNodes    = [][]byte{{1, 1, 1, 1, 1}, {2, 2, 2, 2, 2}, {3, 3, 3, 3, 3}}
	FakeTrieNode     = []byte{248, 68, 1, 128, 160, 6, 180, 135, 209, 92, 2, 139, 109, 245, 108, 62, 187, 155, 112, 134, 150, 94, 186, 58, 36, 8, 87, 166, 71, 250, 236, 226, 255, 19, 38, 159, 43, 160, 206, 51, 34, 13, 92, 127, 13, 9, 215, 92, 239, 247, 108, 5, 134, 60, 94, 125, 110, 128, 28, 112, 223, 231, 213, 212, 93, 76, 68, 232, 6, 84}
	FakeStateLeaf, _ = rlp.EncodeToBytes(state.Account{
		Nonce:    0,
		Balance:  big.NewInt(10000),
		CodeHash: common.HexToHash("0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470").Bytes(),
		Root:     common.HexToHash("0x821e2556a290c86405f8160a2d662042a431ba456b9db265c79bb837c04be5f0"),
	})
	FakeStateAccount = state.Account{
		Nonce:    0,
		Balance:  big.NewInt(10000),
		CodeHash: common.HexToHash("0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470").Bytes(),
		Root:     common.HexToHash("0x821e2556a290c86405f8160a2d662042a431ba456b9db265c79bb837c04be5f0"),
	}
)
