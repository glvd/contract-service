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
			Value: "default",
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
				Name:          "sync",
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

					engine, e := model.InitSQLite3(context.String("database"))
					if e != nil {
						return e
					}
					database := seed.NewDatabase(engine)
					database.RegisterSync(model.Video{}, model.Pin{}, model.Unfinished{})
					//e = model.Sync(engine)
					//if e != nil {
					//	return e
					//}
					api := seed.NewAPI(context.String("shell"))
					seeder.Register(database, api)
					pin := task.NewPin()
					pin.Type = task.PinTypeSync
					pin.Table = task.PinTablePin
					if v := context.String("table"); v != "default" {
						pin.Table = task.PinTable(v)
					}

					pin.From = context.String("from")
					skip := strings.Split(context.String("skip"), ",")
					for _, s := range skip {
						pin.SkipType = append(pin.SkipType, s)
					}

					seeder.Start()
					seeder.AddTasker(pin)
					seeder.Wait()

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

					engine, e := model.InitSQLite3(context.String("database"))
					if e != nil {
						return e
					}
					database := seed.NewDatabase(engine)
					database.RegisterSync(model.Video{}, model.Pin{}, model.Unfinished{})
					//e = model.Sync(engine)
					//if e != nil {
					//	return e
					//}
					api := seed.NewAPI(context.String("shell"))
					seeder.Register(database, api)
					pin := task.NewPin()
					pin.Type = task.PinTypeCheck
					pin.Table = task.PinTablePin
					if v := context.String("table"); v != "default" {
						pin.Table = task.PinTable(v)
					}
					skip := strings.Split(context.String("skip"), ",")
					for _, s := range skip {
						pin.SkipType = append(pin.SkipType, s)
					}

					seeder.Start()
					seeder.AddTasker(pin)
					seeder.Wait()
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
