package pin

import (
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
			Name:        "type",
			DefaultText: "all",
			Usage:       "set the pin check type",
		},
		//&cli.StringFlag{
		//	Name:  "path",
		//	Usage: "set the app path to add  hash",
		//},
		//&cli.IntFlag{
		//	Name:  "limit",
		//	Usage: "set the ban max numbers",
		//},
		//&cli.Int64Flag{
		//	Name:  "code",
		//	Usage: "set the version code for update",
		//},
		//&cli.StringFlag{
		//	Name:    "key",
		//	Usage:   "set the ct process key",
		//	EnvVars: []string{"seedKey"},
		//},
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
					case "slice":
						ps = seed.PinStatusSliceOnly
					case "video":
						ps = seed.PinStatusVideo
					case "poster":
						ps = seed.PinStatusPoster
					}

					s := seed.NewSeed(seed.DatabaseOption("sqlite3", db), seed.Pin(ps, context.Args().Slice()...))
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

					s := seed.NewSeed(seed.DatabaseOption("sqlite3", db), seed.Check(tp))

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
