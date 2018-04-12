package test_helpers

import ipld "gx/ipfs/Qme5bWv7wtjUNGsK2BNGVUFPKiuxWrsqrtvYwCLRw8YFES/go-ipld-format"

type MockAdder struct {
	Called     bool
	PassedNode ipld.Node
	Err        error
}

func NewMockAdder() *MockAdder {
	return &MockAdder{
		Called:     false,
		PassedNode: nil,
		Err:        nil,
	}
}

func (ma *MockAdder) SetError(err error) {
	ma.Err = err
}

func (ma *MockAdder) Add(node ipld.Node) error {
	ma.Called = true
	ma.PassedNode = node
	return ma.Err
}
