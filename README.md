# Block Watcher

## Description
Tool for converting Ethereum blocks into IPLDs

## Dependencies
- Requires running [go-ipld-eth](https://github.com/ipfs/go-ipld-eth) for formatting of raw bytes into an eth-block (e.g. you need to be able to pass the `--input-enc raw --format eth-block` with `ipfs dag put`)
    - on OS X, from the go-ipld-eth directory: `./plugin/hacks/osx.sh`
- Requires Geth running LevelDB to fetch block data

## Setup
- modify config file (`environments/public.toml` or desired alternative) to point to local IPC and level DB (`ipcPath` and `levelDbPath`)
- NOTE: currently, the ipc path and level db path cannot be in the same directory, as the ipc requires geth to be running but accessing level DB requires that Geth not be using it :( - a simple workaround is to use `environments/infura.toml` for the time being
- `make build`
- Sync VulcanizeDB with required blocks
- `./block_watcher createIpldForBlock -b <desired_block_number>`