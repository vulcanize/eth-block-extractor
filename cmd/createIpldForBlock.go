package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/spf13/cobra"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// createIpldForBlockCmd represents the createIpldForBlock command
var createIpldForBlockCmd = &cobra.Command{
	Use:   "createIpldForBlock",
	Short: "Create an IPLD object for a block",
	Long: `Create an IPLD object for a block.

e.g. ./block_watcher createIpldForBlock -b 1234567

Under the hood, the command fetches the block header RLP data from LevelDB and
puts it in IPFS, converting the data as an 'eth-block'`,
	Run: func(cmd *cobra.Command, args []string) {
		createIpldForBlock()
	},
}

var blockNumber int

func init() {
	rootCmd.AddCommand(createIpldForBlockCmd)
	createIpldForBlockCmd.Flags().IntVarP(&blockNumber, "block-number", "b", 0, "Block number to create IPLD for")
}

func createIpldForBlock() {
	// initialize VulcanizeDB block repo with database connection
	blockchain := geth.NewBlockchain(ipc)
	if blockchain.LastBlock().Int64() == 0 {
		log.Fatal("geth initial: state sync not finished")
	}
	db, err := postgres.NewDB(databaseConfig, blockchain.Node())
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}
	blockRepository := repositories.BlockRepository{DB: db}

	// fetch block hash for block number from VulcanizeDB
	_blockNumber := int64(blockNumber)
	block, err := blockRepository.GetBlock(_blockNumber)
	if err != nil {
		log.Fatal("Error fetching block: ", err)
	}
	blockHash := common.HexToHash(block.Hash)

	// fetch block RLP data from LevelDB
	levelDB, err := ethdb.NewLDBDatabase(levelDbPath, 128, 1024)
	if err != nil {
		log.Fatal("Error connecting to levelDB: ", err)
	}
	uintBlockNumber := uint64(block.Number)
	blockData := core.GetHeaderRLP(levelDB, blockHash, uintBlockNumber)

	// create file to temporarily store RLP data
	filename := fmt.Sprintf("blocks/block_%d.bytes", blockNumber)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	written := ioutil.WriteFile(filename, blockData, 0644)
	if written != nil {
		log.Fatal("Error writing block data to file: ", written)
	}

	// write file contents to IPFS
	ipfsCommand := exec.Command("ipfs", "dag", "put", "--input-enc", "raw", "--format", "eth-block", filename)
	output, err := ipfsCommand.Output()
	if err != nil {
		log.Fatal("Error writing file to IPFS", err)
	}
	fmt.Println("Created IPFS hash: ", string(output))
}
