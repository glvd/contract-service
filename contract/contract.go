package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
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
var DefaultNodeAddress = "0x54737e15b826bc72fda0acb7c6e713cf1a987a5d"

// DefaultMessageAddress ...
var DefaultMessageAddress = "0xce50de6720bb96a9ccbd36c46e963720ac2e9ff8"

// DefaultTagAddress ...
var DefaultTagAddress = "0x408cbbb78c10f199a8371e26970be70138e87b20"

// DefaultGasLimit ...
var DefaultGasLimit = "0x7A1200"

func init() {

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

		privateKey, err := keystore.DecryptKey(bys, pass)
		if err != nil {
			panic(e)
		}
		c.key = privateKey.PrivateKey
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

	e = c.Transact(context.Background(), func(c *Contract, opts *bind.TransactOpts) (transaction *types.Transaction, e error) {
		transaction, e = c.tag().SetTagIds(opts, "hot", time.Now().Format(TimeStampFormat), uno)
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
