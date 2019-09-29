package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/glvd/seed/model"
	shell "github.com/godcong/go-ipfs-restapi"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

// GatwayAddress ...
//const defaultGatewayAddress = "http://139.196.215.224:8545"
const defaultGatewayAddress = "https://ropsten.infura.io/QVsqBu3yopMu2svcHqRj"

//const defaultGatewayAddress = "http://192.168.1.13:8545"
//var bangumiContract = "0x333e1e53683FC252E4BF87f1A300A6140d58105d"
var bangumiContract = "0xc15266BB7C1FC17d5aD614dfdbA7316038F35725"
var accContract = "0x80b57a43e5a2c7ad2ce28fbb40c6a58bcf2aa252"
var appContract = "0x4ed0462cb2ecf365dea3503a445cb8e25a23dc03"

func init() {

}

// Contract ...
type Contract struct {
	eth       *Contract
	shell     *shell.Shell
	conn      *ethclient.Client
	key       string
	gateway   string
	processor []ProcessorFunc
}

func (c *Contract) UpdateAppWithPath(code int64, version, path string) error {
	object, e := c.shell.AddFile(path)
	if e != nil {
		return e
	}
	return c.UpdateApp(code, version, object.Hash)
}

func bangumi() ProcessorFunc {
	return func(eth *Contract) interface{} {
		tk, err := NewBangumiData(common.HexToAddress(bangumiContract), eth.conn)
		if err != nil {
			return nil
		}
		return tk
	}
}

func dhash() ProcessorFunc {
	return func(eth *Contract) interface{} {
		tk, err := NewDhash(common.HexToAddress(appContract), eth.conn)
		if err != nil {
			return nil
		}
		return tk
	}
}

func node() ProcessorFunc {
	return func(eth *Contract) interface{} {
		tk, err := NewAccelerateNode(common.HexToAddress(accContract), eth.conn)
		if err != nil {
			return nil
		}
		return tk
	}
}

// ProcessorFunc ...
type ProcessorFunc func(eth *Contract) interface{}

// NewContract ...
func NewContract(key string) *Contract {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial(defaultGatewayAddress)
	if err != nil {
		return nil
	}
	return &Contract{
		conn:      conn,
		key:       key,
		gateway:   defaultGatewayAddress,
		processor: []ProcessorFunc{bangumi(), dhash(), node()},
	}
}

// PrivateKey ...
func (c *Contract) PrivateKey() (key *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(c.key)
	if err != nil {
		return nil
	}
	return privateKey
}

// RegisterContract ...
func (c *Contract) RegisterContract(fn func(eth *Contract) interface{}) {
	c.processor = append(c.processor, fn)
}

// ProcContract ...
func (c *Contract) ProcContract(fn func(v interface{}) (bool, error)) (e error) {
	for _, proc := range c.processor {
		if b, e := fn(proc(c)); b {
			return e
		}
	}
	return nil
}

// InfoInput ...
func (c *Contract) InfoInput(video *model.Video, b bool) (e error) {
	e = singleInput(c, video, b)
	if e != nil {
		return e
	}
	return nil
}

// CheckNameExists ...
func (c *Contract) CheckNameExists(ban string, idx ...int) (e error) {
	idxStr := ""
	nb := ""
	if idx == nil {
		idx = append(idx, 1)
	}

	for _, i := range idx {
		idxStr = strconv.FormatInt(int64(i), 10)
		nb = strings.ToUpper(ban + "@" + idxStr)
		_, e = c.CheckExist(nb)
		if e != nil {
			return e
		}
	}
	return nil
}

// GetLastVersionCode ...
func (c *Contract) GetLastVersionCode() (code *big.Int, e error) {
	err := c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}
		code, _, e = data.GetLatest(&bind.CallOpts{Pending: true})
		return true, e
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *Contract) GetLatest() (code *big.Int, hash string, e error) {
	err := c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}
		code, hash, e = data.GetLatest(&bind.CallOpts{Pending: true})
		if e != nil {
			return false, e
		}
		return true, nil
	})
	if err != nil {
		return &big.Int{}, "", err
	}
	return
}

func (c *Contract) GetCodeVersion(code *big.Int) (ver string, e error) {
	err := c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}
		ver, e = data.VersionTable(&bind.CallOpts{Pending: true}, code)
		if e != nil {
			return false, e
		}
		return true, nil
	})
	if err != nil {
		return "", err
	}
	return
}

// GetLastVersionHash ...
func (c *Contract) GetLastVersionHash() (ver, hash string, e error) {
	err := c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}
		_, ver, e = data.GetLatest(&bind.CallOpts{Pending: true})
		if e != nil {
			return false, e
		}
		hash, e = data.GetHash(&bind.CallOpts{Pending: true}, ver)
		if e != nil {
			return true, e
		}
		return true, e
	})
	if err != nil {
		return "", "", err
	}
	return
}

// UpdateHotList ...
func (c *Contract) UpdateHotList(list ...string) (e error) {
	err := c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}
		opt := bind.NewKeyedTransactor(c.PrivateKey())
		bytes, e := jsoniter.Marshal(list)
		if e != nil {
			return false, e
		}
		transaction, e := data.UpdateHotList(opt, string(bytes))
		if e != nil {
			return true, e
		}
		ctx := context.Background()
		receipt, err := bind.WaitMined(ctx, c.conn, transaction)
		if err != nil {
			return true, err
		}
		log.Debugf("receipt is :%x\n", string(receipt.TxHash[:]))
		return true, nil
	})
	if err != nil {
		return err
	}
	return
}

// GetHostList ...
func (c *Contract) GetHostList() (list []string) {
	err := c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}
		ll, e := data.GetHotList(&bind.CallOpts{Pending: true})
		if e != nil {
			return true, e
		}
		e = jsoniter.Unmarshal([]byte(ll), &list)
		if e != nil {
			return false, e
		}
		return true, nil
	})
	if err != nil {
		return nil
	}
	return list
}

// CheckExist ...
func (c *Contract) CheckExist(ban string) (hash string, e error) {
	e = c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*BangumiData)
		if !b {
			return false, nil
		}
		hash, e = data.QueryHash(&bind.CallOpts{Pending: true}, ban)
		log.With("size", len(hash), "hash", hash, "name", ban).Info("checked")
		if hash == "" || hash == "," || len(hash) != 46 {
			return true, xerrors.New(ban + " hash is not found!")
		}
		return true, e
	})
	return
}

func (c *Contract) CheckInfo(ban string, hash1, sp1 string, te1 string, ts1 string) (hash string, e error) {
	log.With("hash", hash1, "sharpness", sp1, "total episode", te1, "total season", ts1).Info("input")
	e = c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*BangumiData)
		if !b {
			return false, nil
		}

		sharpness, e := data.QuerySharpness(&bind.CallOpts{Pending: true}, ban)
		if sharpness != sp1 {
			log.With("size", len(hash), "sharpness", sharpness, "name", ban).Info("checked")
			return true, xerrors.New(ban + " sharpness is not found!")
		}

		te, e := data.QueryTotalEpisode(&bind.CallOpts{Pending: true}, ban)
		if te != te1 {
			log.With("size", len(hash), "te", te, "name", ban).Info("checked")
			return true, xerrors.New(ban + " total episode is not found!")
		}

		ts, e := data.QueryTotalSeason(&bind.CallOpts{Pending: true}, ban)
		if ts != ts1 {
			log.With("size", len(hash), "ts", ts, "name", ban).Info("checked")
			return true, xerrors.New(ban + " total season is not found!")
		}
		hash = hash1
		return true, e
	})
	return
}

func (c *Contract) AddReplaceNode(index int64, node ...string) (e error) {
	return addReplaceNode(c, index, node...)
}

// Close ...
func (c *Contract) Close() {
	if c.conn == nil {
		return
	}
	c.conn.Close()
}

func singleInput(c *Contract, video *model.Video, update bool) (e error) {
	return c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := (v).(*BangumiData)
		if !b {
			return false, nil
		}
		opt := bind.NewKeyedTransactor(c.PrivateKey())
		name := video.Bangumi
		var hash string
		upperName := strings.ToUpper(name + "@" + video.Episode)

		hash, e = c.CheckExist(upperName)

		if e == nil {
			if update {
				hash, e = c.CheckInfo(upperName, hash, video.Sharpness, video.TotalEpisode, video.Season)
			}
			if e == nil && hash == video.M3U8Hash {
				return
			}
		}

		if len(video.Role) > 5 {
			video.Role = video.Role[:5]
		}
		roles := strings.Join(video.Role, " ")

		transaction, err := data.InfoInput(opt,
			strings.ToUpper(upperName),
			video.PosterHash,
			roles,
			video.M3U8Hash,
			video.Intro,
			video.Sharpness,
			video.Episode,
			video.TotalEpisode,
			video.Season,
			video.Format,
			"",
			"")
		if err != nil {
			return true, err
		}
		ctx := context.Background()
		receipt, err := bind.WaitMined(ctx, c.conn, transaction)
		if err != nil {
			return true, err
		}
		log.Info(name + "@" + video.Episode + " success")
		log.Debugf("receipt is :%x\n", string(receipt.TxHash[:]))
		return true, nil
	})

}

func multipleInput(c *Contract, video *model.Video) (e error) {
	return c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := (v).(*BangumiData)
		if !b {
			return false, nil
		}
		opt := bind.NewKeyedTransactor(c.PrivateKey())
		name := video.Bangumi

		var hash string
		upperName := strings.ToUpper(name + "@" + video.Episode)
		hash, e = c.CheckExist(upperName)
		if e == nil && hash == video.M3U8Hash {
			return
		}
		transaction, err := data.InfoInput(opt,
			strings.ToUpper(upperName),
			video.PosterHash,
			video.Role[0],
			video.M3U8Hash,
			video.Alias[0],
			video.Sharpness,
			video.Episode,
			video.TotalEpisode,
			video.Season,
			video.Format,
			"",
			"")
		if err != nil {
			return true, err
		}
		ctx := context.Background()
		receipt, err := bind.WaitMined(ctx, c.conn, transaction)
		if err != nil {
			//log.Fatalf("tx mining error:%v\n", err)
			return true, err
		}
		log.Info(name + "@" + video.Episode + " success")
		log.Debugf("receipt is :%x\n", string(receipt.TxHash[:]))

		return true, nil
	})

}

// UpdateApp ...
func (c *Contract) UpdateApp(code int64, version string, hash string) (e error) {
	return updateAppHash(c, code, version, hash)
}

func updateAppHash(c *Contract, code int64, version string, hash string) (e error) {
	return c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*Dhash)
		if !b {
			return false, nil
		}

		code := big.NewInt(code)
		key := c.PrivateKey()
		opt := bind.NewKeyedTransactor(key)
		transaction, err := data.UpdateVersion(opt, version, hash, code)
		if err != nil {
			return true, e
		}

		ctx := context.Background()
		receipt, err := bind.WaitMined(ctx, c.conn, transaction)
		if err != nil {
			return true, e
		}
		log.Info(version + "@" + hash + " success")
		log.Debugf("receipt is :%x\n", string(receipt.TxHash[:]))
		return true, nil
	})
}

func addReplaceNode(c *Contract, index int64, newNode ...string) (e error) {
	return c.ProcContract(func(v interface{}) (b bool, e error) {
		node, b := v.(*AccelerateNode)
		if !b {
			return false, nil
		}

		key := c.PrivateKey()
		opt := bind.NewKeyedTransactor(key)
		ctx := context.Background()

		if index == -1 {
			nodes, e := node.Get(&bind.CallOpts{Pending: true})
			if e != nil {
				return true, e
			}
			log.With("total", len(nodes)).Info(nodes)
			var exist bool
			for _, nodeT := range newNode {
				exist = false
				for i, n := range nodes {
					if nodeT == n {
						log.With("index", i, "node", n).Info("exist")
						exist = true
						break
					}
				}
				if !exist {
					log.With("node", nodeT).Info("add")
					transaction, e := node.Add(opt, nodeT)
					if e != nil {
						return true, e
					}
					receipt, err := bind.WaitMined(ctx, c.conn, transaction)
					if err != nil {
						return true, e
					}
					log.Debugf("receipt is :%x\n", string(receipt.TxHash[:]))
				}
			}
			return true, nil
		}

		if newNode == nil || len(newNode) == 0 {
			return true, xerrors.New("no data")
		}

		transaction, err := node.Replace(opt, uint32(index), newNode[0])
		if err != nil {
			return true, e
		}

		receipt, err := bind.WaitMined(ctx, c.conn, transaction)
		if err != nil {
			return true, e
		}
		log.Debugf("receipt is :%x\n", string(receipt.TxHash[:]))
		return true, nil
	})
}
