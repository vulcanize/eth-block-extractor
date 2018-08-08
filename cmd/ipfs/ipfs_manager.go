package ipfs

type IpfsManager interface {
	EnsureConfig(ipfsPath string) error
}
