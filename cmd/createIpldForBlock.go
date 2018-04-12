// Copyright © 2018 Rob Mulholand <rmulholand@8thlight.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/geth"

	"github.com/8thlight/block_watcher/pkg"
	"github.com/8thlight/block_watcher/pkg/db"
	"github.com/8thlight/block_watcher/pkg/ipfs"
	"github.com/8thlight/block_watcher/pkg/ipfs/eth_block_header"
)

// createIpldForBlockCmd represents the createIpldForBlock command
var createIpldForBlockCmd = &cobra.Command{
	Use:   "createIpldForBlock",
	Short: "Create an IPLD object for a block.",
	Long: `Create an IPLD object for a block.

e.g. ./block_watcher createIpldForBlock -b 1234567

Under the hood, the command fetches the block header RLP data from LevelDB and
puts it in IPFS, converting the data as an 'eth-block'`,
	Run: func(cmd *cobra.Command, args []string) {
		createIpldForBlock()
	},
}

var blockNumber int64

func init() {
	rootCmd.AddCommand(createIpldForBlockCmd)
	createIpldForBlockCmd.Flags().Int64VarP(&blockNumber, "block-number", "b", 0, "Create IPLD for this block")
	createIpldForBlockCmd.Flags().BoolVarP(&useParity, "parity", "p", false, "Use Parity's RocksDB instead of Geth's LevelDB")
}

func createIpldForBlock() {
	// init blockchain
	blockchain := geth.NewBlockchain(ipc)
	if blockchain.LastBlock().Int64() == 0 {
		log.Fatal("geth initial: state sync not finished")
	}

	// init pg db with blockchain
	pgDB, err := postgres.NewDB(databaseConfig, blockchain.Node())
	if err != nil {
		log.Fatal("Error connecting to postgres db: ", err)
	}

	// init block repository with pg db
	blockRepository := repositories.BlockRepository{DB: pgDB}

	// init eth db
	var ethDBConfig db.DatabaseConfig
	if useParity {
		ethDBConfig = db.CreateDatabaseConfig(db.Rocks, rocksDbPath)
	} else {
		ethDBConfig = db.CreateDatabaseConfig(db.Level, levelDbPath)
	}
	ethDB, err := db.CreateDatabase(ethDBConfig)
	if err != nil {
		log.Fatal("Error connecting to ethereum db: ", err)
	}

	// init ipfs publisher
	ipfsNode, err := ipfs.InitIPFSNode(ipfsPath)
	if err != nil {
		log.Fatal("Error connecting to IPFS: ", err)
	}
	decoder := ipfs.RlpDecoder{}
	dagPutter := eth_block_header.NewBlockHeaderDagPutter(*ipfsNode, decoder)
	publisher := ipfs.NewIpfsPublisher(dagPutter)

	// execute transformer
	transformer := pkg.NewTransformer(blockRepository, ethDB, publisher)
	err = transformer.Execute(blockNumber, blockNumber)
	if err != nil {
		log.Fatal("Error executing transformer: ", err.Error())
	}
}
