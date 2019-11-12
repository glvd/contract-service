package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"service/contract"

	"github.com/glvd/conversion"
	"github.com/glvd/seed/model"
	"github.com/godcong/go-trait"
	"github.com/goextension/log"
	"github.com/urfave/cli/v2"
	"github.com/xormsharp/xorm"
)

const applicationName = "deploy"

var _db *xorm.Engine
var _contract *contract.Contract

func init() {
	log.Register(trait.NewZapFileSugar("zap.log"))
	contract.DefaultGatway = "http://139.196.215.224:8545"
	contract.DefaultMessageAddress = "0xdfbff0f931cf056b9eb5a8b58616024881215f01"
	contract.DefaultTagAddress = "0x431d6215052cd3f0cec1838289a3d99abc496f5e"
	contract.DefaultNodeAddress = "0X5A144FECD913688D0A755CEE0275FD8F95A767A4"

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
	ConfigPath = filepath.Join(dir, DefaultConfigPath, DefaultConfigName)
}

func main() {
	app := cli.NewApp()
	app.Version = "v0.0.1"
	app.Name = applicationName
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "load deploy config",
			Value: ConfigPath,
		},
		&cli.StringFlag{
			Name:  "gateway",
			Usage: "set the remote gate",
			Value: contract.DefaultGatway,
		},
		&cli.StringFlag{
			Name:  "keypath",
			Usage: "set the key file path",
			Value: "945d35cd4a6549213e8d37feb5d708ec98906902",
		},
		&cli.StringFlag{
			Name:  "keypass",
			Usage: "set the key file decode pass",
			Value: "123",
		},

		&cli.StringFlag{
			Name:  "tag",
			Usage: "set the tag contract address",
			Value: contract.DefaultTagAddress,
		},
		&cli.StringFlag{
			Name:  "node",
			Usage: "set the node contract address",
			Value: contract.DefaultNodeAddress,
		},
		&cli.StringFlag{
			Name:  "message",
			Usage: "set the message contract address",
			Value: contract.DefaultMessageAddress,
		},
	}
	app.Before = before()

	app.Action = action()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func before() cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		e := LoadConfig(ctx.String("config"))
		if e != nil {
			log.Errorw("load config error", "error", e)
		}
		engine, e := MakeInstance(_config.DBConfig)
		if e != nil {
			return e
		}
		engine.ShowSQL()
		engine.ShowExecTime()
		conversion.RegisterDatabase(engine)
		_db = engine
		gateway := ctx.String("gateway")
		keypath := ctx.String("keypath")
		keypass := ctx.String("keypass")
		dnode := ctx.String("node")
		dtag := ctx.String("tag")
		dmessage := ctx.String("dmessage")
		_contract = contract.NewContract(contract.ETHClient(gateway),
			contract.FileKey(keypath, keypass),
			//contract.HexKey("9efef8ebc3c51e91fb7f9faf7dbd516cb320ade03108c1568c9cee01a39af311"),
			contract.Node(dnode),
			contract.Tag(dtag),
			contract.Message(dmessage))

		return nil
	}
}

func action() cli.ActionFunc {
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
			cv := SeedVideoToConversionVideo(*v)
			json, err := cv.MarshalJSONVersion()
			if err != nil {
				return err
			}
			err = _contract.AddVideo(cv.No, cv.ID(), json, cv.JSONVersion())
			if err != nil {
				return err
			}
			log.Infow("contract update", "no", cv.No, "id", cv.ID())
		}

		return nil
	}
}
