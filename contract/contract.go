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

// None ...
const (
	None Type = iota
	WriteAble
	DMessage
	DNode
	DTag
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

// LoadMessage ...
func LoadMessage(msg *dmessage.Dmessage) {

}

// Message ...
func Message(addr string) Options {
	return func(c *Contract) {
		newDmessage, e := dmessage.NewDmessage(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(DMessage, newDmessage)
	}
}

// Tag ...
func Tag(addr string) Options {
	return func(c *Contract) {
		newDtag, e := dtag.NewDtag(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(DTag, newDtag)
	}
}

// Node ...
func Node(addr string) Options {
	return func(c *Contract) {
		newDnode, e := dnode.NewDnode(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(DTag, newDnode)
	}
}

func (c *Contract) message() (msg *dmessage.Dmessage) {
	if v, b := c.contracts.Load(DMessage); b {
		if msg, b = v.(*dmessage.Dmessage); b {
			return
		}
	}
	return nil
}

func (c *Contract) tag() (tag *dtag.Dtag) {
	if v, b := c.contracts.Load(DTag); b {
		if tag, b = v.(*dtag.Dtag); b {
			return
		}
	}
	return nil
}

func (c *Contract) node() (node *dnode.Dnode) {
	if v, b := c.contracts.Load(DNode); b {
		if node, b = v.(*dnode.Dnode); b {
			return
		}
	}
	return nil
}
