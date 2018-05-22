package transformers

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vulcanize/block_watcher/pkg/db"
	"github.com/vulcanize/block_watcher/pkg/ipfs"
)

type EthStateTrieTransformer struct {
	database  db.Database
	decoder   db.Decoder
	publisher ipfs.Publisher
}

func NewEthStateTrieTransformer(database db.Database, decoder db.Decoder, publisher ipfs.Publisher) *EthStateTrieTransformer {
	return &EthStateTrieTransformer{database: database, decoder: decoder, publisher: publisher}
}

func (t EthStateTrieTransformer) Execute(startingBlockNumber int64, endingBlockNumber int64) error {
	for i := startingBlockNumber; i <= endingBlockNumber; i++ {
		root, err := t.getStateRootForBlock(i)
		if err != nil {
			return err
		}

		stateTrieNodes, err := t.database.GetStateTrieNodes(root.Bytes())
		if err != nil {
			return fmt.Errorf("Error fetching state trie for block %d: %s\n", i, err)
		}

		err = t.writeStateTrieNodesToIpfs(stateTrieNodes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t EthStateTrieTransformer) getStateRootForBlock(blockNumber int64) (common.Hash, error) {
	var header types.Header
	rawHeader, err := t.database.GetBlockHeaderByBlockNumber(blockNumber)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Error fetching header for block %d: %s\n", blockNumber, err)
	}
	out, err := t.decoder.Decode(rawHeader, &header)
	parsedHeader := out.(*types.Header)
	return parsedHeader.Root, nil
}

func (t EthStateTrieTransformer) writeStateTrieNodesToIpfs(stateTrieNodes [][]byte) error {
	for _, node := range stateTrieNodes {
		output, err := t.publisher.Write(node)
		if err != nil {
			return fmt.Errorf("Error writing state trie node to ipfs: %s\n", err)
		}
		log.Println("Created ipld: ", output)
	}
	return nil
}