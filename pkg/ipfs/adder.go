package ipfs

import ipld "gx/ipfs/Qme5bWv7wtjUNGsK2BNGVUFPKiuxWrsqrtvYwCLRw8YFES/go-ipld-format"

type Adder interface {
	Add(node ipld.Node) error
}
