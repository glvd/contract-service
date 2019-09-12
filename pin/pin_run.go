package pin

import (
	"strings"

	"github.com/glvd/seed"
	"github.com/glvd/seed/model"
	"github.com/glvd/seed/task"
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
			Name:  "table",
			Usage: "which table to pin",
			Value: "video",
		},
		&cli.StringFlag{
			Name:  "skip",
			Usage: "set the skip args",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "get the from address",
		},
		//&cli.IntFlag{
		//	Name:  "limit",
		//	Usage: "set the ban max numbers",
		//},
		&cli.BoolFlag{
			Name:  "showsql",
			Usage: "set the show sql option",
		},
		//&cli.StringFlag{
		//	Name:  "stype",
		//	Usage: "skip type",
		//},
	)
	return &cli.Command{
		Name:   "pin",
		Usage:  "pin file to ipfs",
		Action: nil,
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
					seeder := seed.NewSeed()

					db := context.String("database")
					if db == "" {
						db = "cs.db"
					}

					engine, e := model.InitSQLite3(db)
					if e != nil {
						return e
					}
					database := seed.NewDatabase(engine)
					database.RegisterSync(model.Video{}, model.Pin{}, model.Unfinished{})

					api := seed.NewAPI(context.String("shell"))
					seeder.Register(database, api)

					pin := task.NewPin()
					pin.Type = task.PinTypeAdd

					skip := strings.Split(context.String("skip"), ",")
					for _, s := range skip {
						pin.SkipType = append(pin.SkipType, s)
					}
					//ps := seed.PinStatusAll
					//switch context.String("status") {
					//case "relate":
					//	ps = seed.PinStatusAssignRelate
					//case "hash":
					//	ps = seed.PinStatusAssignHash
					//case "unfinished":
					//	ps = seed.PinStatusUnfinished
					//case "video":
					//	ps = seed.PinStatusVideo
					//case "poster":
					//	ps = seed.PinStatusPoster
					//case "sync":
					//	ps = seed.PinStatusSync
					//}
					//stype := context.String("stype")
					//var csa []string
					//if stype != "" {
					//	csa = strings.Split(stype, ",")
					//}
					//
					//s := seed.NewSeed(seed.DatabaseOption("sqlite3", db),
					//	seed.Pin(seed.PinStatusArg(ps),
					//		seed.PinSkipArg(csa),
					//		seed.PinListArg(context.Args().Slice()...)))
					//s.From = context.String("from")
					//j := context.String("json")
					//if j != "" {
					//	s.Register(seed.Information(context.String("json"), seed.InfoFlagBSON))
					//}
					//shell := context.String("shell")
					//if shell != "" {
					//	s.Register(seed.ShellOption(shell))
					//}
					//
					//if context.Bool("skip") {
					//	s.Register(seed.SkipConvertOption())
					//	s.Register(seed.SkipSourceOption())
					//}
					//
					//s.AfterInit(seed.SyncDatabase())
					//if context.Bool("showsql") {
					//	s.AfterInit(seed.ShowSQLOption())
					//}
					//
					//s.Workspace = context.String("workspace")
					//s.Start()
					//
					//s.Wait()
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
					seeder := seed.NewSeed()

					db := context.String("database")
					if db == "" {
						db = "cs.db"
					}

					engine, e := model.InitSQLite3(db)
					if e != nil {
						return e
					}
					database := seed.NewDatabase(engine)
					//database.RegisterSync(model.Video{}, model.Pin{}, model.Unfinished{})

					api := seed.NewAPI(context.String("shell"))
					seeder.Register(database, api)
					pin := task.NewPin()
					pin.Type = task.PinTypeCheck

					skip := strings.Split(context.String("skip"), ",")
					for _, s := range skip {
						pin.SkipType = append(pin.SkipType, s)
					}

					seeder.Start()
					seeder.AddTasker(pin)
					seeder.Wait()

					//db := context.String("database")
					//if db == "" {
					//	db = "cs.db"
					//}
					//eng, e := model.InitDB("sqlite3", db)
					//if e != nil {
					//	return e
					//}
					//model.InitMainDB(eng)
					//tp := context.String("type")
					//ctp := context.String("ctype")
					//stype := context.String("stype")
					//var csa []string
					//if stype != "" {
					//	csa = strings.Split(stype, ",")
					//}
					//log.With("types", csa).Info("skip check")
					//s := seed.NewSeed(seed.DatabaseOption("sqlite3", db),
					//	seed.Check(seed.CheckPinTypeArg(tp),
					//		seed.CheckTypeArg(seed.CheckType(ctp)),
					//		seed.CheckSkipArg(csa)),
					//)
					//
					//s.AfterInit(seed.ShowSQLOption())
					//
					//s.From = context.String("from")
					//api := context.String("api")
					//if api != "" {
					//	s.Register(seed.APIOption(api))
					//}
					//
					//shell := context.String("shell")
					//if shell != "" {
					//	s.Register(seed.ShellOption(shell))
					//}
					//
					//s.AfterInit(seed.SyncDatabase())
					//s.Workspace = context.String("workspace")
					//s.Start()
					//
					//s.Wait()
					//TODO
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
