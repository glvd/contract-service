package contract

import (
	"encoding/csv"
	"io"
	"math/big"
	"os"

	shell "github.com/godcong/go-ipfs-restapi"

	"github.com/glvd/seed/model"
	"github.com/go-xorm/xorm"
	"github.com/gocarina/gocsv"
	"github.com/godcong/go-trait"
	"golang.org/x/xerrors"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdContract ...
func CmdContract(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:  "address",
			Value: "",
			Usage: "contract address",
		},
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "video",
			Usage:   "contract process type",
		},
		&cli.StringFlag{
			Name:  "node",
			Usage: "node add the node to accelerate",
		},
		&cli.Int64Flag{
			Name:  "index",
			Value: -1,
			Usage: "index set the node replace index",
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
			Value: "db",
			Usage: "tell me where the list from db/args",
		},
		&cli.BoolFlag{
			Name:  "update",
			Usage: "set the bool to update hash",
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
			//address := context.String("address")
			//if address == "" {
			//	panic("address must set use -address,-a")
			//}
			version := context.String("release")
			path := context.String("path")
			hash := context.String("hash")
			code := context.Int64("code")
			contract := NewContract(key)
			if contract == nil {
				panic("null contract")
			}
			sh := context.String("shell")
			if sh == "" {
				sh = "localhost:5001"
			}
			contract.shell = shell.NewShell(sh)

			db := context.String("database")
			log.Info("database:", db)
			eng, e := model.InitSQLite3(db)
			if e != nil {
				return e
			}

			switch context.String("type") {
			case "video":
				var session *xorm.Session
				session = eng.NoCache()
				if context.NArg() > 0 {
					session = eng.In("bangumi", context.Args().Slice())
				}

				if context.String("from") == "db" {
					videos, e := model.Top(session, 0)

					if e != nil {
						log.Error()
						return e
					}
					log.Info("find:", len(*videos))
					update := context.Bool("update")
					for _, video := range *videos {
						if !update && video.M3U8Hash == "" {
							continue
						}
						e := contract.InfoInput(video, update)
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
				for i := 1; i <= limit; i++ {
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
				if from := context.String("from"); from == "db" {
					//TODO
					videos, e := model.Top(nil, 50)
					if e != nil {
						log.Error(e)
						return e
					}
					for _, video := range *videos {
						list = append(list, video.Bangumi)
					}
				} else if from == "csv" {
					lists, e := getHotList(context.String("path"))
					if e != nil {
						return e
					}
					for _, li := range lists {
						list = append(list, li.EventName)
					}
				}
				e := contract.UpdateHotList(list...)
				if e != nil {
					log.Error(e)
					return e
				}
			case "status":
				var ver string
				var lastHash string
				var e error
				log.With("key", key, "address", address).Info("input info")
				if code > 0 {
					ver, e = contract.GetCodeVersion(big.NewInt(code))
					if e != nil {
						log.Error(e)
						return e
					}
					log.With("version", ver, "hash", lastHash, "code", code).Info("version")
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
			case "acc":
				var nodes []string
				if context.NArg() > 0 {
					nodes = context.Args().Slice()
				}
				log.With("nodes", nodes).Info("add node")
				e := contract.AddReplaceNode(context.Int64("index"), nodes...)
				if e != nil {
					log.Error(e)
					return e
				}
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

type HotList struct {
	EventName             string  `csv:"Event Name"`
	AppName               string  `csv:"App Name"`
	TotalOccurrences      int64   `csv:"Total Occurrences"`
	AvgEventsPerSession   float64 `csv:"Avg Events per Session"`
	OccurrencesDailyAvg   int64   `csv:"Occurrences Daily Avg"`
	UniqueDevicesDailyAvg int64   `csv:"Unique Devices Daily Avg"`
}

func getHotList(path string) ([]*HotList, error) {
	file, e := os.Open(path)
	if e != nil {
		log.Error(e)
		return nil, e
	}
	var hl []*HotList
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ','
		r.Comment = '#'
		return r // Allows use pipe as delimiter
	})

	e = gocsv.UnmarshalFile(file, &hl)
	if e != nil {
		return nil, e
	}
	return hl, nil
}
