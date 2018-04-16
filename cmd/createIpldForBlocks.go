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

	"github.com/8thlight/block_watcher/pkg"
	"github.com/8thlight/block_watcher/pkg/db"
	"github.com/8thlight/block_watcher/pkg/fs"
	"github.com/8thlight/block_watcher/pkg/ipfs"
	"github.com/spf13/cobra"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres/repositories"
	"github.com/vulcanize/vulcanizedb/pkg/geth"
)

// createIpldForBlocksCmd represents the createIpldForBlocks command
var createIpldForBlocksCmd = &cobra.Command{
	Use:   "createIpldForBlocks",
	Short: "Create IPLD objects for multiple blocks.",
	Long: `Create IPLD objects for multiple blocks.

e.g. ./block_watcher createIpldForBlocks -s 1234567 -e 4567890

Under the hood, the command fetches the block header RLP data from LevelDB and
puts it in IPFS, converting the data as an 'eth-block'.`,
	Run: func(cmd *cobra.Command, args []string) {
		createIpldForBlocks()
	},
}

var startingBlockNumber int64
var endingBlockNumber int64

func init() {
	rootCmd.AddCommand(createIpldForBlocksCmd)
	createIpldForBlocksCmd.Flags().Int64VarP(&startingBlockNumber, "starting-block-number", "s", 0, "First block number to create IPLD for")
	createIpldForBlocksCmd.Flags().Int64VarP(&endingBlockNumber, "ending-block-number", "e", 5430000, "Last block number to create IPLD for")
}

func createIpldForBlocks() {
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
	ethDBConfig := db.CreateDatabaseConfig(db.Level, levelDbPath)
	ethDB, err := db.CreateDatabase(ethDBConfig)
	if err != nil {
		log.Fatal("Error connecting to ethereum db: ", err)
	}

	// init ipfs publisher
	blockWriter := fs.NewBlockFileWriter(fs.FileCreator{}, fs.FileWriter{})
	ipfsWriter := ipfs.NewIpfsEthBlockWriter(fs.ExecCommander{})
	publisher := ipfs.NewIpfsPublisher(blockWriter, ipfsWriter)

	// execute transformer
	transformer := pkg.NewTransformer(blockRepository, ethDB, publisher)
	err = transformer.Execute(startingBlockNumber, endingBlockNumber)
	if err != nil {
		log.Fatal("Error executing transformer: ", err.Error())
	}
}
