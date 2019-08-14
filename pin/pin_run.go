package pin

import (
	"strings"

	"github.com/glvd/seed"
	"github.com/glvd/seed/model"
	"github.com/godcong/go-trait"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdPin ...
func CmdPin(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:  "status",
			Usage: "set the pin status (all/relate/hash/unfinished) default is all",
		},
		&cli.StringFlag{
			Name:  "type",
			Usage: "set the pin check type",
			Value: "recursive",
		},
		&cli.StringFlag{
			Name:  "ctype",
			Usage: "set the pin check type pin/unpin",
			Value: "pin",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "get the from address",
		},
		//&cli.IntFlag{
		//	Name:  "limit",
		//	Usage: "set the ban max numbers",
		//},
		//&cli.Int64Flag{
		//	Name:  "code",
		//	Usage: "set the version code for update",
		//},
		&cli.StringFlag{
			Name:  "stype",
			Usage: "skip type",
		},
	)
	return &cli.Command{
		Name:    "pin",
		Aliases: []string{"P"},
		Usage:   "pin file to ipfs",
		Action:  nil,
		Subcommands: []*cli.Command{
			{
				Name:          "add",
				Aliases:       nil,
				Usage:         "",
				UsageText:     "",
				Description:   "",
				ArgsUsage:     "",
				Category:      "",
				ShellComplete: nil,
				Before:        nil,
				After:         nil,
				Action: func(context *cli.Context) error {
					db := context.String("database")
					if db == "" {
						db = "cs.db"
					}
					eng, e := model.InitDB("sqlite3", db)
					if e != nil {
						return e
					}
					model.InitMainDB(eng)
					ps := seed.PinStatusAll
					switch context.String("status") {
					case "relate":
						ps = seed.PinStatusAssignRelate
					case "hash":
						ps = seed.PinStatusAssignHash
					case "unfinished":
						ps = seed.PinStatusUnfinished
					case "video":
						ps = seed.PinStatusVideo
					case "poster":
						ps = seed.PinStatusPoster
					}
					stype := context.String("stype")
					var csa []string
					if stype != "" {
						csa = strings.Split(stype, ",")
					}

					s := seed.NewSeed(seed.DatabaseOption("sqlite3", db),
						seed.Pin(seed.PinStatusArg(ps),
							seed.PinSkipArg(csa),
							seed.PinListArg(context.Args().Slice()...)))
					j := context.String("json")
					if j != "" {
						s.Register(seed.Information(context.String("json"), seed.InfoFlagBSON))
					}
					shell := context.String("shell")
					if shell != "" {
						s.Register(seed.ShellOption(shell))
					}

					if context.Bool("skip") {
						s.Register(seed.SkipConvertOption())
						s.Register(seed.SkipSourceOption())
					}

					s.AfterInit(seed.SyncDatabase())
					s.Workspace = context.String("workspace")
					s.Start()

					s.Wait()
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
			},
			{
				Name:          "check",
				Aliases:       nil,
				Usage:         "",
				UsageText:     "",
				Description:   "",
				ArgsUsage:     "",
				Category:      "",
				ShellComplete: nil,
				Before:        nil,
				After:         nil,
				Action: func(context *cli.Context) error {
					db := context.String("database")
					if db == "" {
						db = "cs.db"
					}
					eng, e := model.InitDB("sqlite3", db)
					if e != nil {
						return e
					}
					model.InitMainDB(eng)
					tp := context.String("type")
					ctp := context.String("ctype")
					stype := context.String("stype")
					var csa []string
					if stype != "" {
						csa = strings.Split(stype, ",")
					}
					log.With("types", csa).Info("skip check")
					s := seed.NewSeed(seed.DatabaseOption("sqlite3", db),
						seed.Check(seed.CheckPinTypeArg(tp),
							seed.CheckTypeArg(seed.CheckType(ctp)),
							seed.CheckSkipArg(csa)),
					)

					s.AfterInit(seed.ShowSQLOption())

					s.From = context.String("from")
					api := context.String("api")
					if api != "" {
						s.Register(seed.APIOption(api))
					}

					shell := context.String("shell")
					if shell != "" {
						s.Register(seed.ShellOption(shell))
					}

					s.AfterInit(seed.SyncDatabase())
					s.Workspace = context.String("workspace")
					s.Start()

					s.Wait()
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
			},
		},
	}
}
