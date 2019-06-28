package contract

import (
	"github.com/go-xorm/xorm"
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed/model"
	"golang.org/x/xerrors"
	"gopkg.in/urfave/cli.v2"
	"math/big"
)

var log = trait.NewZapSugar()

// CmdContract ...
func CmdContract(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			Value:   "",
			Usage:   "contract address",
		},
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "video",
			Usage:   "contract process type",
		},
		&cli.StringFlag{
			Name:  "ban",
			Usage: "ban no to check",
		},
		&cli.StringFlag{
			Name:    "release",
			Value:   "v0.0.1",
			Aliases: []string{"r"},
			Usage:   "set the application version",
		},
		&cli.StringFlag{
			Name:  "hash",
			Usage: "set the app ipfs hash",
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "set the app path to add  hash",
		},
		&cli.IntFlag{
			Name:  "limit",
			Usage: "set the ban max numbers",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "tell me where the list from db/args",
		},
		&cli.Int64Flag{
			Name:  "code",
			Usage: "set the version code for update",
		},
		&cli.StringFlag{
			Name:    "key",
			Usage:   "set the ct process key",
			EnvVars: []string{"seedKey"},
		},
	)
	return &cli.Command{
		Name:          "contract",
		Aliases:       []string{"C"},
		Usage:         "contract share the video info to ct.",
		UsageText:     "",
		Description:   "",
		ArgsUsage:     "",
		Category:      "",
		ShellComplete: nil,
		Before:        nil,
		After:         nil,
		Action: func(context *cli.Context) error {
			log.Info("contract call")
			key := context.String("key")
			if key == "" {
				panic("key must set use -key or Env(seedKey)")
			}
			address := context.String("a")
			if address == "" {
				panic("address must set use -address,-a")
			}
			version := context.String("release")
			path := context.String("path")
			hash := context.String("hash")
			code := context.Int64("code")
			contract := NewContract(key, address)
			if contract == nil {
				panic("null contract")
			}
			switch context.String("t") {
			case "video":
				var session *xorm.Session
				if context.NArg() > 0 {
					var ins []interface{}
					for _, v := range context.Args().Slice() {
						ins = append(ins, v)
					}
					session = model.DB().In("bangumi", ins...)
				}
				if context.String("from") == "db" {
					videos, e := model.TopList(session, 0)
					if e != nil {
						log.Error()
						return e
					}
					for _, video := range *videos {
						e := contract.InfoInput(video)
						if e != nil {
							log.Error(e)
							continue
						}
					}
				}

			case "check":
				if context.NArg() <= 0 {
					return xerrors.New("nothing to check")
				}
				bans := context.Args().Slice()

				limit := context.Int("limit")
				var checkList []int
				for i := 0; i <= limit; i++ {
					checkList = append(checkList, i)
				}
				for _, ban := range bans {
					log.With("ban", ban, "limit", limit).Info("check")
					e := contract.CheckNameExists(ban, checkList...)
					if e != nil {
						log.Error(e)
						continue
					}
				}
			case "hot":
				list := context.Args().Slice()
				if context.String("from") == "db" {
					videos, e := model.TopList(nil, 50)
					if e != nil {
						log.Error()
						return e
					}
					for _, video := range *videos {
						list = append(list, video.Bangumi)
					}
				}
				e := contract.UpdateHotList(list...)
				if e != nil {
					log.Error()
					return e
				}
			case "status":
				var ver string
				var lastHash string
				var e error
				if code > 0 {
					ver, e = contract.GetCodeVersion(big.NewInt(code))
					if e != nil {
						log.Error(e)
						return e
					}
					log.With("version", ver, "hash", lastHash, "code", code).Info("version:", ver)
					return nil
				}
				ver, lastHash, e = contract.GetLastVersionHash()
				if e != nil {
					log.Error(e)
					return e
				}
				log.With("version", ver, "hash", lastHash, "code", code).Info("last")
				return nil
			case "app":
				if path != "" {
					if err := contract.UpdateAppWithPath(code, version, path); err != nil {
						log.Error(err)
						return err
					}
					return nil
				}
				e := contract.UpdateApp(code, version, hash)
				if e != nil {
					log.Error(e)
					return e
				}
				return nil
			}
			return nil
		},
		OnUsageError:       nil,
		Subcommands:        nil,
		Flags:              flags,
		SkipFlagParsing:    false,
		HideHelp:           false,
		Hidden:             false,
		HelpName:           "",
		CustomHelpTemplate: "",
	}
}
