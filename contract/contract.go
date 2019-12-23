package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"service/contract/dhashcoin"
	"strings"
	"sync"
	"time"

	"service/contract/dmessage"
	"service/contract/dnode"
	"service/contract/dtag"
	"service/dhcrypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	DHC
)

// Type ...
type Type int

// Contract ...
type Contract struct {
	contracts *sync.Map
	conn      *ethclient.Client
	key       *ecdsa.PrivateKey
	gasLimit  *big.Int
}

// Options ...
type Options func(c *Contract)

// TransactOpts ...
type TransactOpts func(c *Contract, opts *bind.TransactOpts) (*types.Transaction, error)

// CallOpts ...
type CallOpts func(c *Contract, opts *bind.CallOpts) error

// DefaultGatway ...
//var DefaultGatway = "http://139.196.215.224:8545"
var DefaultGatway = "http://13.124.213.107:8545"

// DefaultNodeAddress ...
var DefaultNodeAddress = "0x3c87dad9055b8075a8515d5c5a68d72d0d2c099e"

// DefaultMessageAddress ...
var DefaultMessageAddress = "0x4833267bde3aa043803a1fa8c3e071f708367da4"

// DefaultTagAddress ...
var DefaultTagAddress = "0x8d64f6d57c7ee984cce09f89969f706216ac03d9"

// DefaultDHCAddress ...
var DefaultDHCAddress = "0x89b92612b9795f8bfed840df9a33a0a3d6745250"

// DefaultGasLimit ...
var DefaultGasLimit = "0x7A1200"

func init() {

}

// Close ...
func (c *Contract) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// Register ...
func (c *Contract) Register(p Type, v interface{}) {
	c.contracts.Store(p, v)
}

// NewContract ...
func NewContract(opts ...Options) *Contract {
	c := &Contract{
		contracts: &sync.Map{},
		gasLimit:  hexutil.MustDecodeBig(DefaultGasLimit),
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

// FileKey ...
func FileKey(path string, pass string) Options {
	return func(c *Contract) {
		bys, e := ioutil.ReadFile(path)
		if e != nil {
			panic(e)
		}

		keys, err := keystore.DecryptKey(bys, pass)
		if err != nil {
			panic(e)
		}
		c.key = keys.PrivateKey
	}
}

// ETHClient ...
func ETHClient(addr string) Options {
	return func(c *Contract) {
		conn, err := ethclient.Dial(addr)
		if err != nil {
			panic(fmt.Errorf("address:%s,error:%w", addr, err))
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
		newDmessage, e := dmessage.NewDMessage(common.HexToAddress(addr), c.conn)
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
		newDtag, e := dtag.NewDTag(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(DTag, newDtag)
	}
}

// Node ...
func Node(addr string) Options {
	return func(c *Contract) {
		if c.conn == nil {
			panic("null connect")
		}
		newDnode, e := dnode.NewDNode(common.HexToAddress(addr), c.conn)
		if e != nil {
			panic(e)
		}
		c.Register(DNode, newDnode)
	}
}

func (c *Contract) message() (msg *dmessage.DMessage) {
	if v, b := c.contracts.Load(DMessage); b {
		if msg, b = v.(*dmessage.DMessage); b {
			return
		}
	}
	return nil
}

func (c *Contract) tag() (tag *dtag.DTag) {
	if v, b := c.contracts.Load(DTag); b {
		if tag, b = v.(*dtag.DTag); b {
			return
		}
	}
	return nil
}

func (c *Contract) node() (node *dnode.DNode) {
	if v, b := c.contracts.Load(DNode); b {
		if node, b = v.(*dnode.DNode); b {
			return
		}
	}
	return nil
}
func (c *Contract) dhc() (node *dhashcoin.DHC) {
	if v, b := c.contracts.Load(DHC); b {
		if node, b = v.(*dhashcoin.DHC); b {
			return
		}
	}
	return nil
}

// Transact ...
func (c *Contract) Transact(ctx context.Context, opt TransactOpts) error {
	o := bind.NewKeyedTransactor(c.key)
	o.GasLimit = c.gasLimit.Uint64()
	transaction, e := opt(c, o)
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

// Call ...
func (c *Contract) Call(ctx context.Context, opt CallOpts) error {
	o := &bind.CallOpts{Pending: true}
	return opt(c, o)
}

// RemoveNode ...
func (c *Contract) RemoveNode(idx int) {

}

// AddNodes ...
func (c *Contract) AddNodes(copyOld bool, ts time.Time, ss ...string) (e error) {
	ctx := context.Background()
	var last *big.Int
	e = c.Call(ctx, func(c *Contract, opts *bind.CallOpts) (e error) {
		last, e = c.node().GetNodeLast(opts)
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

		e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (*types.Transaction, error) {
			ts := big.NewInt(n.Unix())
			return c.node().SetNodeLast(opts, ts)
		})
		if e != nil {
			return e
		}
	}

	enc := dhcrypto.NewCipherEncoder([]byte(PublicKey), int(n.UnixNano()), n)
	for _, s := range ss {
		e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
			encoded, e := enc.Encode(s)
			if e != nil {
				log.Errorw("encode error", "errors", e, "source", s)
				return nil, e
			}
			return c.node().StoreNode(opts, string(encoded))
		})
		if e != nil {
			log.Errorw("encode error", "errors", e, "source", s)
		}
	}
	return nil
}

// SetVersion ...
func (c *Contract) SetVersion(device int64, version string, hash string) (e error) {
	ctx := context.Background()
	e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, err error) {
		return c.node().SetDeviceVersion(opts, big.NewInt(device), version, hash)
	})
	if e != nil {
		return e
	}
	return nil
}

// GetVersion ...
func (c *Contract) GetVersion(device int64) (version string, hash string, e error) {
	ctx := context.Background()
	e = c.Call(ctx, func(c *Contract, opts *bind.CallOpts) (e error) {
		version, hash, e = c.node().GetDeviceVersion(opts, big.NewInt(device))
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		return "", "", e
	}
	return
}

// GetNodes ...
func (c *Contract) GetNodes(ts time.Time) ([]string, *big.Int, error) {
	ctx := context.Background()
	var bi *big.Int
	var ss []string
	e := c.Call(ctx, func(c *Contract, opts *bind.CallOpts) (e error) {
		if ts.IsZero() {
			bi, e = c.node().GetNodeLast(opts)
			if e != nil {
				return e
			}
			ss, _, e = c.node().GetNodeLastData(opts)
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
		bys, e := dec.Decode(s)
		if e != nil {
			log.Errorw("decode errors", "error", e, "source", s)
			continue
		}
		results = append(results, string(bys))
	}

	return results, bi, nil
}

// DeployNode ...
func (c *Contract) DeployNode() (addr *common.Address, e error) {
	ctx := context.Background()
	e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		address, tx, _, e := dnode.DeployDNode(opts, c.conn)
		if e != nil {
			return nil, e
		}
		addr = &address
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		addressAfterMined, err := bind.WaitDeployed(ctx, c.conn, tx)
		if err != nil {
			log.Errorf("failed to deploy contact when mining :%v", err)
			return nil, err
		}
		if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
			log.Errorf("mined address :%s,before mined address:%s", addressAfterMined, address)
			return nil, errors.New("compare error")
		}
		return tx, nil
	})
	if e != nil {
		return nil, e
	}
	return addr, nil
}

// DeployTag ...
func (c *Contract) DeployTag(msgAddr *common.Address) (addr *common.Address, e error) {
	ctx := context.Background()
	if msgAddr == nil {
		return nil, errors.New("message address cant intput with null")
	}
	e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		address, tx, _, e := dtag.DeployDTag(opts, c.conn, *msgAddr)
		if e != nil {
			return nil, e
		}
		addr = &address
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		addressAfterMined, err := bind.WaitDeployed(ctx, c.conn, tx)
		if err != nil {
			log.Errorf("failed to deploy contact when mining :%v", err)
			return nil, err
		}
		if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
			log.Errorf("mined address :%s,before mined address:%s", addressAfterMined, address)
			return nil, errors.New("compare error")
		}
		return tx, nil
	})
	if e != nil {
		return nil, e
	}
	return addr, nil
}

// DeployMessage ...
func (c *Contract) DeployMessage() (addr *common.Address, e error) {
	ctx := context.Background()
	e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		address, tx, _, e := dmessage.DeployDMessage(opts, c.conn)
		if e != nil {
			return nil, e
		}
		addr = &address
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
		startTime := time.Now()
		fmt.Printf("TX start @:%s", time.Now())
		ctx := context.Background()
		addressAfterMined, err := bind.WaitDeployed(ctx, c.conn, tx)
		if err != nil {
			log.Fatalf("failed to deploy contact when mining :%v", err)
		}
		fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
		if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
			log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
		}
		return tx, nil
	})
	if e != nil {
		return nil, e
	}
	return addr, nil
}

// DeployDHC ...
func (c *Contract) DeployDHC() (addr *common.Address, e error) {
	ctx := context.Background()
	e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		address, tx, _, e := dhashcoin.DeployDHC(opts, c.conn)
		if e != nil {
			return nil, e
		}
		addr = &address
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		addressAfterMined, err := bind.WaitDeployed(ctx, c.conn, tx)
		if err != nil {
			log.Errorf("failed to deploy contact when mining :%v", err)
			return nil, err
		}
		if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
			log.Errorf("mined address :%s,before mined address:%s", addressAfterMined, address)
			return nil, errors.New("compare error")
		}
		return tx, nil
	})
	if e != nil {
		return nil, e
	}
	return addr, nil
}

// OpenMessageAuthority ...
func (c *Contract) OpenMessageAuthority() (e error) {
	ctx := context.Background()
	isWriter := false
	e = c.Call(ctx, func(c *Contract, opts *bind.CallOpts) error {
		writerList, e := c.message().Writers(opts)
		if e != nil {
			return e
		}
		for _, wr := range writerList {
			if CompareNoCap(wr.Hex(), DefaultTagAddress) == 0 {
				isWriter = true
				return nil
			}
		}
		return nil
	})
	if e != nil {
		return e
	}
	log.Infow("iswriter", "address", DefaultTagAddress, "is", isWriter)
	if isWriter {
		log.Info("already is writers")
		return nil
	}
	e = c.Transact(ctx, func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		trans, e := c.message().IncreasedWritership(opts, common.HexToAddress(DefaultTagAddress))
		if e != nil {
			return nil, e
		}
		return trans, nil
	})
	if e != nil {
		return e
	}
	return nil
}

// AddOrUpdateVideo ...
func (c *Contract) AddOrUpdateVideo(no string, message VideoMessage, update bool) (e error) {
	msg, e := c.GetVideo(no)

	if e == nil && CompareNoCap(msg.ID, message.ID) == 0 {
		if update {
			e = c.UpdateVideo(no, message)
			if e != nil {
				return e
			}
		}
		log.Infow("exist", "update", update)
		return nil
	}
	e = c.AddVideo(no, message)
	if e != nil {
		return e
	}
	return nil
}

// UpdateVideo ...
func (c *Contract) UpdateVideo(no string, message VideoMessage) (e error) {
	no = strings.ToUpper(no)
	e = c.Transact(context.Background(), func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		transaction, e = c.message().UpdateMessage(opts, message.ID, MustJSON(message.Encode()))
		if e != nil {
			return nil, e
		}
		return transaction, nil
	})
	if e != nil {
		return e
	}
	return nil
}

// AddVideo ...
func (c *Contract) AddVideo(no string, message VideoMessage) (e error) {
	no = strings.ToUpper(no)
	e = c.Transact(context.Background(), func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		transaction, e = c.tag().AddTagMessage(opts, "video", no, message.ID, MustJSON(message.Encode()))
		if e != nil {
			return nil, e
		}
		return transaction, nil
	})
	if e != nil {
		return e
	}
	return nil
}

// AddHot ...
func (c *Contract) AddHot(no ...string) (e error) {
	return c.AddTagVideos("hot", time.Now().Format(TimeStampFormat), no...)
}

// AddEveryDay ...
func (c *Contract) AddEveryDay(no ...string) (e error) {
	return c.AddTagVideos("everyday", time.Now().Format(TimeStampFormat), no...)
}

// AddTagVideos ...
func (c *Contract) AddTagVideos(tag string, date string, no ...string) (e error) {
	var uno []string
	for _, n := range no {
		n = strings.ToUpper(n)
		msg, e := c.GetVideo(n)
		if e != nil {
			log.Errorw("error", "no", n, "error", e)
			continue
		}
		uno = append(uno, msg.ID)
	}

	if len(uno) == 0 {
		return errors.New("no not added")
	}

	if date == "" {
		date = time.Now().Format(TimeStampFormat)
	}

	e = c.Transact(context.Background(), func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		transaction, e = c.tag().SetTagIds(opts, tag, date, uno)
		if e != nil {
			return nil, e
		}
		return transaction, nil
	})
	if e != nil {
		return e
	}
	return nil
}

// GetTagVideos ...
func (c *Contract) GetTagVideos(tag, date string) (v []VideoMessage, i int64, e error) {
	var msg struct {
		Value []string
		Size  *big.Int
	}
	e = c.Call(context.Background(), func(c *Contract, opts *bind.CallOpts) error {
		msg, e = c.tag().GetTagMessage(opts, tag, date)
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		return nil, 0, e
	}
	m, err := DecodeMessages(msg.Value)
	if err != nil {
		return nil, 0, err
	}
	return m, msg.Size.Int64(), nil
}

// GetVideo ...
func (c *Contract) GetVideo(no string) (messages VideoMessage, err error) {
	no = strings.ToUpper(no)
	struct0s, size, err := c.GetVideos(no)
	if err != nil {
		return VideoMessage{}, err
	}
	if size <= 0 {
		return VideoMessage{}, errors.New("data not found")
	}
	return struct0s[0], nil
}

// GetVideos ...
func (c *Contract) GetVideos(no string) (messages []VideoMessage, size int64, err error) {
	no = strings.ToUpper(no)
	var msg struct {
		Value []string
		Size  *big.Int
	}
	e := c.Call(context.Background(), func(c *Contract, opts *bind.CallOpts) error {
		msg, err = c.tag().GetTagMessage(opts, "video", no)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, 0, e
	}
	m, err := DecodeMessages(msg.Value)
	if err != nil {
		return nil, 0, err
	}
	return m, msg.Size.Int64(), nil
}
