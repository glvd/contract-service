package contract

import (
	"sync"

	"service/contract/dmessage"
	"service/contract/dnode"
	"service/contract/dtag"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Type ...
type Type int

// ContractNone ...
const (
	ContractNone Type = iota
	ContractWriteAble
	ContractMessage
	ContractNode
	ContractTag
)

var contract *Contract

// Contract ...
type Contract struct {
	contracts *sync.Map
	conn      *ethclient.Client
}

// Options ...
type Options func(c *Contract)

func init() {
	contract = &Contract{}
}

// Register ...
func (c *Contract) Register(p Type, v interface{}) {
	c.contracts.Store(p, v)
}

// NewContract ...
func NewContract(opts ...Options) *Contract {
	c := &Contract{
		contracts: &sync.Map{},
	}
	for _, op := range opts {
		op(c)
	}
	return c
}

// Message ...
func Message(addr string) Options {
	return func(c *Contract) {
		newDmessage, e := dmessage.NewDmessage(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(ContractMessage, newDmessage)
	}
}

// Tag ...
func Tag(addr string) Options {
	return func(c *Contract) {
		newDtag, e := dtag.NewDtag(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(ContractTag, newDtag)
	}
}

// Node ...
func Node(addr string) Options {
	return func(c *Contract) {
		newDnode, e := dnode.NewDnode(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(ContractTag, newDnode)
	}
}
