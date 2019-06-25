package contract

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/godcong/go-ipfs-restapi"
	jsoniter "github.com/json-iterator/go"
	"github.com/yinhevr/seed/model"
	"golang.org/x/xerrors"
	"math/big"
	"strconv"
	"strings"
)

// GatwayAddress ...
const defaultGatewayAddress = "https://ropsten.infura.io/QVsqBu3yopMu2svcHqRj"

// Contract ...
type Contract struct {
	eth       *Contract
	shell     *shell.Shell
	conn      *ethclient.Client
	key       string
	contract  string
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
		tk, err := NewBangumiData(common.HexToAddress(eth.contract), eth.conn)
		if err != nil {
			return nil
		}
		return tk
	}
}

func dhash() ProcessorFunc {
	return func(eth *Contract) interface{} {
		tk, err := NewDhash(common.HexToAddress(eth.contract), eth.conn)
		if err != nil {
			return nil
		}
		return tk
	}
}

// ProcessorFunc ...
type ProcessorFunc func(eth *Contract) interface{}

// NewContract ...
func NewContract(key, contract string) *Contract {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial(defaultGatewayAddress)
	if err != nil {
		return nil
	}
	return &Contract{
		conn:      conn,
		key:       key,
		gateway:   defaultGatewayAddress,
		contract:  contract,
		processor: []ProcessorFunc{bangumi(), dhash()},
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
func (c *Contract) InfoInput(video *model.Video) (e error) {
	e = infoInput(c, video)
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
		e = c.CheckExist(nb)
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
func (c *Contract) CheckExist(ban string) (e error) {
	return c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := v.(*BangumiData)
		if !b {
			return false, nil
		}
		hash, e := data.QueryHash(&bind.CallOpts{Pending: true}, "ban")
		log.With("size", len(hash), "hash", hash, "name", ban).Info("checked")
		if hash == "" || hash == "," || len(hash) != 46 {
			return true, xerrors.New(ban + " hash is not found!")
		}
		return true, e
	})

}

// Close ...
func (c *Contract) Close() {
	if c.conn == nil {
		return
	}
	c.conn.Close()
}

// ContractProcessor ...
//type ContractProcessor interface {
//	ContractProc(eth *Contract) error
//}

func singleInput(c *Contract, video *model.Video) (e error) {
	return c.ProcContract(func(v interface{}) (b bool, e error) {
		data, b := (v).(*BangumiData)
		if !b {
			return false, nil
		}
		opt := bind.NewKeyedTransactor(c.PrivateKey())
		name := video.Bangumi
		//list := video.VideoGroupList[0]
		//objMax := len(list.Object)
		//objMaxStr := strconv.FormatInt(int64(objMax), 10)
		//for i := 0; i < objMax; i++ {
		//	idxStr := strconv.FormatInt(int64(i+1), 10)
		upperName := strings.ToUpper(name + "@" + video.Episode)
		e = c.CheckExist(upperName)
		if e == nil {
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

		upperName := strings.ToUpper(name + "@" + video.Episode)
		e = c.CheckExist(upperName)
		if e == nil {
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

// InfoInput ...
func infoInput(eth *Contract, video *model.Video) (e error) {
	//if video == nil || video.VideoGroupList == nil {
	//	return
	//}
	//
	//vgMax := len(video.VideoGroupList)
	fn := singleInput
	//if vgMax > 1 {
	//	fn = multipleInput
	//}
	return fn(eth, video)
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
