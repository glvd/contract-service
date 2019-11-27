package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"service/contract"

	"github.com/glvd/conversion"
	"github.com/glvd/seed/model"
	"github.com/godcong/go-trait"
	"github.com/goextension/log"
	"github.com/urfave/cli/v2"
	"github.com/xormsharp/xorm"
)

const applicationName = "tools"

var _db *xorm.Engine
var _contract *contract.Contract

func init() {
	log.Register(trait.NewZapFileSugar("zap.log"))
	//contract.DefaultGatway = "http://139.196.215.224:8545"
	//contract.DefaultMessageAddress = "0xdfbff0f931cf056b9eb5a8b58616024881215f01"
	//contract.DefaultTagAddress = "0x431d6215052cd3f0cec1838289a3d99abc496f5e"
	//contract.DefaultNodeAddress = "0X5A144FECD913688D0A755CEE0275FD8F95A767A4"

	//contract.DefaultGatway = "http://localhost:8545"
	//contract.DefaultNodeAddress = "0xe512a2a61814b8de98a52a0dfd7e13627e40014a"
	//contract.DefaultMessageAddress = "0xcD05aAD053f605EE1A76B0B78cDfdfd32bBcfB7b"
	//contract.DefaultTagAddress = "0x696687f41be961A521dD331A1609Edf7E4E7b2b5"
	dir, e := os.Getwd()
	if e != nil {
		dir := os.Getenv("HOME")
		if dir == "" {
			usr, err := user.Current()
			if err != nil {
				panic(fmt.Sprintf("cannot get current user: %s", err))
			}
			dir = usr.HomeDir
		}
	}
	ConfigPath = filepath.Join(dir, DefaultConfigName)
}
func deployFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "load deploy config",
			Value: ConfigPath,
		},
		&cli.StringFlag{
			Name:        "gateway",
			Usage:       "set the remote gate",
			Destination: &_init.Gateway,
			Value:       contract.DefaultGatway,
		},
		&cli.StringFlag{
			Name:        "keypath",
			Usage:       "set the key file path",
			Destination: &_init.KeyPath,
			Value:       DefaultKeyPath,
		},
		&cli.StringFlag{
			Name:        "keypass",
			Usage:       "set the key file decode pass",
			Destination: &_init.KeyPass,
			Value:       DefaultKeyPass,
		},

		&cli.StringFlag{
			Name:        "tag",
			Usage:       "set the tag contract address",
			Destination: &_init.Contract.DTag,
			Value:       contract.DefaultTagAddress,
		},
		&cli.StringFlag{
			Name:        "node",
			Usage:       "set the node contract address",
			Destination: &_init.Contract.DNode,
			Value:       contract.DefaultNodeAddress,
		},
		&cli.StringFlag{
			Name:        "message",
			Usage:       "set the message contract address",
			Destination: &_init.Contract.DMessage,
			Value:       contract.DefaultMessageAddress,
		},
		&cli.BoolFlag{
			Name:  "init",
			Usage: "init deploy contract address",
			Value: false,
		},
	}
}

func deployBefore() cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		cfgPath := ctx.String("config")
		e := LoadConfig(cfgPath)
		if e != nil {
			log.Errorw("load config error", "error", e)
		}
		log.Infow("config info", "info", _config)
		engine, e := MakeInstance(_config.Database)
		if e != nil {
			return e
		}
		engine.ShowSQL()
		engine.ShowExecTime()
		conversion.RegisterDatabase(engine)
		_db = engine

		if ctx.Bool("init") {
			initC := contract.NewContract(
				contract.ETHClient(_init.Gateway),
				contract.FileKey(_init.KeyPath, _init.KeyPass),
			)
			defer initC.Close()
			msgAddr, e := initC.DeployMessage()
			if e != nil {
				return e
			}
			_config.Contract.DMessage = fmt.Sprintf("0x%x", msgAddr)
			log.Infow("message deployed", "address", _config.Contract.DMessage)
			tagAddr, e := initC.DeployTag(msgAddr)
			if e != nil {
				return e
			}
			_config.Contract.DTag = fmt.Sprintf("0x%x", tagAddr)
			log.Infow("tag deployed", "address", _config.Contract.DTag)
			nodeAddr, e := initC.DeployNode()
			if e != nil {
				return e
			}
			_config.Contract.DNode = fmt.Sprintf("0x%x", nodeAddr)
			log.Infow("node deployed", "address", _config.Contract.DNode)

		}
		initDefault()
		e = SaveConfig(cfgPath)
		if e != nil {
			return e
		}
		log.Infow("contract",
			"message", contract.DefaultMessageAddress,
			"tag", contract.DefaultTagAddress,
			"node", contract.DefaultNodeAddress,
		)

		_contract = contract.NewContract(
			contract.ETHClient(contract.DefaultGatway),
			contract.FileKey(DefaultKeyPath, DefaultKeyPass),
			//contract.HexKey("9efef8ebc3c51e91fb7f9faf7dbd516cb320ade03108c1568c9cee01a39af311"),
			contract.Node(contract.DefaultNodeAddress),
			contract.Tag(contract.DefaultTagAddress),
			contract.Message(contract.DefaultMessageAddress),
		)

		e = _contract.OpenMessageAuthority()
		if e != nil {
			log.Errorw("authority error", "error", e)
			return e
		}
		return nil
	}
}

func deployAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		var seedV model.Video
		log.Info("process video")
		csv := make(chan *model.Video)
		go func(csv chan<- *model.Video) {
			defer close(csv)
			rows, e := _db.Rows(seedV)
			if e != nil {
				log.Errorw("found error", "error", e)
				return
			}

			for rows.Next() {
				var v model.Video
				err := rows.Scan(&v)
				if err != nil {
					log.Errorw("scan error", "error", err)
					return
				}

				//skip
				if v.M3U8Hash == "" {
					continue
				}
				csv <- &v
			}
		}(csv)

		for v := range csv {
			select {
			case <-ctx.Context.Done():
				log.Info("exit")
				return nil
			default:
			}
			cv := SeedVideoToConversionVideo(*v)
			json, err := cv.MarshalJSONVersion()
			if err != nil {
				return err
			}
			msg := contract.VideoMessage{
				ID:      cv.ID(),
				Content: json,
				Version: cv.JSONVersion(),
			}
			log.Infow("contract update", "no", cv.No, "id", cv.ID())
			for retry := 0; retry < 3; retry++ {
				err = _contract.AddOrUpdateVideo(cv.No, msg, false)
				if err != nil {
					log.Errorw("update failed", "error", err)
					time.Sleep(5 * time.Second)
					continue
				}
				break
			}

		}

		log.Info("deploy end")
		return nil
	}
}
