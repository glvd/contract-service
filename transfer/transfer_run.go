package transfer

import (
	"errors"

	"github.com/glvd/seed"
	"github.com/glvd/seed/model"
	"github.com/glvd/seed/task"
	"github.com/godcong/go-trait"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdPin ...
func CmdTransfer(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name: "path",
			//Value: ".json",
			Usage: "transfer path input/output",
		},
		&cli.StringFlag{
			Name:  "status",
			Usage: "set transfer status",
		},
	)
	return &cli.Command{
		Name:  "transfer",
		Usage: "transfer from other or to other database",
		Subcommands: []*cli.Command{
			{
				Name:          "db",
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

					path := context.Path("path")
					if path == "" {
						path = "source.db"
					}
					dbt := task.NewDBTransfer(model.MustDatabase(model.InitSQLite3(path)))
					//jst := task.NewJSONTransfer(path)
					dbt.Status = task.TransferStatusFromOther

					sdb := seed.NewDatabase(model.MustDatabase(model.InitSQLite3(context.String("database"))))
					//sdb.RegisterSync(model.Video{}, model.Pin{}, model.Unfinished{})

					proc := seed.NewProcess()
					s := seed.NewSeed(sdb, proc)
					s.Start()
					s.AddTasker(dbt)
					s.Wait()
					log.Info("transfer end")
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
				Name:          "json",
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
					//dbt := task.NewDBTransfer(model.MustDatabase(model.InitSQLite3("0916.db")))
					path := context.Path("path")
					if path == "" {
						path = "output.json"
					}
					jst := task.NewJSONTransfer(path)
					jst.Status = task.TransferStatusToJSON

					sdb := seed.NewDatabase(model.MustDatabase(model.InitSQLite3(context.String("database"))))
					//sdb.RegisterSync(model.Video{}, model.Pin{}, model.Unfinished{})

					proc := seed.NewProcess()
					s := seed.NewSeed(sdb, proc)
					s.Start()
					s.AddTasker(jst)
					s.Wait()
					log.Info("transfer end")
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
		Action: func(context *cli.Context) error {
			return errors.New("not enough args to call")
		},
		Flags: flags,
	}
}
