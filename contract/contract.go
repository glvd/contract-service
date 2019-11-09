package contract

import (
	"context"
	"crypto/ecdsa"
	"dhcrypto"
	"math/big"
	"strings"
	"sync"
	"time"

	"service/contract/dmessage"
	"service/contract/dnode"
	"service/contract/dtag"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goextension/log"
)

// TimeStampFormat ...
const TimeStampFormat = "20060102"

// None ...
const (
	None Type = iota
	WriteAble
	DMessage
	DNode
	DTag
)

// Type ...
type Type int

// Contract ...
type Contract struct {
	contracts *sync.Map
	conn      *ethclient.Client
	key       *ecdsa.PrivateKey
}

// Options ...
type Options func(c *Contract)

var contract *Contract

// DefaultGatway ...
var DefaultGatway = ""

// DefaultNodeAddress ...
var DefaultNodeAddress = ""

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

// HexKey ...
func HexKey(key string) Options {
	return func(c *Contract) {
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			panic(err)
		}
		c.key = privateKey
	}
}

// ETHClient ...
func ETHClient(addr string) Options {
	return func(c *Contract) {
		conn, err := ethclient.Dial(addr)
		if err != nil {
			panic(err)
		}
		c.conn = conn
	}
}

// Message ...
func Message(addr string) Options {
	return func(c *Contract) {
		if c.conn == nil {
			panic("null connect")
		}
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
		if c.conn == nil {
			panic("null connect")
		}
		//newDtag, e := dtag.NewDtag(common.HexToAddress(addr), c.conn)
		//if e != nil {
		//	panic(e)
		//}
		//c.Register(DTag, newDtag)
	}
}

// Node ...
func Node(addr string) Options {
	return func(c *Contract) {
		if c.conn == nil {
			panic("null connect")
		}
		newDnode, e := dnode.NewDnode(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(DNode, newDnode)
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

// TransactOpts ...
type TransactOpts func(opts *bind.TransactOpts) (*types.Transaction, error)

// Transact ...
func (c *Contract) Transact(ctx context.Context, opt TransactOpts) error {
	o := bind.NewKeyedTransactor(c.key)
	transaction, e := opt(o)
	if e != nil {
		return e
	}
	receipt, err := bind.WaitMined(ctx, c.conn, transaction)
	if err != nil {
		return err
	}
	log.Infof("receipt is :%x", receipt.TxHash[:])
	return nil
}

// CallOpts ...
type CallOpts func(opts *bind.CallOpts) error

// Call ...
func (c *Contract) Call(ctx context.Context, opt CallOpts) error {
	o := &bind.CallOpts{Pending: true}
	return opt(o)
}

// AddNodes ...
func (c *Contract) AddNodes(copyOld bool, ts time.Time, ss ...string) (e error) {
	ctx := context.Background()
	var last *big.Int
	e = c.Call(ctx, func(opts *bind.CallOpts) (e error) {
		last, e = c.node().GetLast(opts)
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		return e
	}

	n := time.Now()
	if !ts.IsZero() {
		n = ts
	}
	newTS := n.Format(TimeStampFormat)
	old := time.Unix(last.Int64(), 0)
	oldTS := old.Format(TimeStampFormat)
	if strings.Compare(oldTS, newTS) != 0 {
		if copyOld {
			oldNodes, _, e := c.GetNodes(old)
			if e != nil {
				return e
			}
			ss = append(ss, oldNodes...)
		}

		e = c.Transact(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
			ts := big.NewInt(n.Unix())
			return c.node().SetLast(opts, ts)
		})
		if e != nil {
			return e
		}
	}

	enc := dhcrypto.NewCipherEncoder([]byte(PublicKey), int(n.UnixNano()), n)
	for _, s := range ss {
		e = c.Transact(ctx, func(opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
			encoded, e := enc.Encode(s)
			if e != nil {
				log.Errorw("encode error", "errors", e, "source", s)
				return nil, e
			}
			return c.node().Store(opts, string(encoded))
		})
		if e != nil {
			log.Errorw("encode error", "errors", e, "source", s)
		}
	}
	return nil
}

// GetNodes ...
func (c *Contract) GetNodes(ts time.Time) ([]string, *big.Int, error) {
	ctx := context.Background()
	var bi *big.Int
	var ss []string
	e := c.Call(ctx, func(opts *bind.CallOpts) (e error) {
		if ts.IsZero() {
			bi, e = c.node().GetLast(opts)
			if e != nil {
				return e
			}
			ss, _, e = c.node().GetLastNode(opts)
		} else {
			bi = big.NewInt(ts.Unix())
			ss, _, e = c.node().GetNode(opts, big.NewInt(ts.Unix()))
		}
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		return nil, nil, e
	}
	var results []string
	dec := dhcrypto.NewCipherDecode([]byte(PrivateKey), time.Unix(bi.Int64(), 0))
	for _, s := range ss {
		bytes, e := dec.Decode(s)
		if e != nil {
			log.Errorw("decode errors", "error", e, "source", s)
			continue
		}
		results = append(results, string(bytes))
	}

	return results, bi, nil
}

// AddVideo ...
func (c *Contract) AddVideo(json string, version string) (e error) {
	e = c.Transact(context.Background(), func(opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		//transaction, e = c.tag().(opts, "video", json, version)i
		//if e != nil {
		return nil, e
		//}
		//return transaction, nil
	})
	if e != nil {
		return e
	}
	return nil
}

// GetVideo ...
func (c *Contract) GetVideo() {

}
