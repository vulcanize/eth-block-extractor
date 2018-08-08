// Copyright © 2018 James Christie <jchristie@8thlight.com>
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
	"github.com/spf13/cobra"
	"github.com/vulcanize/eth-block-extractor/cmd/ipfs"
	"github.com/vulcanize/eth-block-extractor/cmd/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/config"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize IPFS and Postgres schema",
	Long: `Initialize IPFS and Postgres schema

e.g. ./eth-block-extractor init

This command establishes (or ensures) the existence of an IPFS config directory
in a non-standard location: '~/.vulcanize/eth-block-extractor-ipfs' and creates
a database for the Postgres driver with associated schema (using the current
user)`,
	Run: func(cmd *cobra.Command, args []string) {
		//var initErr = initializeIpfsWithPostgres()
		//if initErr != nil {
		//	log.Fatal("Error creating IPFS configuration", ipfsErr)
		//}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	createIpldForBlockHeaderCmd.Flags()
}

func initializeIpfsWithPostgres(ipfsManager ipfs.IpfsManager, ipfsPath string, postgresManager postgres.PostgresManager, databaseConfig config.Database) error {
	var ipfsErr = ipfsManager.EnsureConfig(ipfsPath)

	if ipfsErr != nil {
		return ipfsErr
	}

	var postgresErr = postgresManager.EnsureSchema(databaseConfig)

	if postgresErr != nil {
		return postgresErr
	}

	return nil
}
