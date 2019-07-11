package transfer

import (
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdPin ...
func CmdTransfer(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:  "from",
			Usage: "set the other sqlite3 path",
		},
		&cli.StringFlag{
			Name:  "to",
			Usage: "set the json output path",
		},
		&cli.StringFlag{
			Name:  "status",
			Usage: "set transfer status",
		},
	)
	return &cli.Command{
		Name:    "transfer",
		Aliases: []string{"T"},
		Usage:   "transfer from other or to other database",
		Action: func(context *cli.Context) error {
			db := context.String("database")
			if db == "" {
				db = "cs.db"
			}
			//eng, e := model.InitDB("sqlite3", db)
			//if e != nil {
			//	return e
			//}
			//model.InitMainDB(eng)

			from := context.String("from")
			status := seed.TransferStatusOther
			switch context.String("status") {
			case "old":
				status = seed.TransferStatusOld
			}

			s := seed.NewSeed(seed.DatabaseOption("sqlite3", db), seed.Transfer(from, seed.InfoFlagSQLite, status))
			s.AfterInit(seed.ShowSQLOption(), seed.SyncDatabase())
			s.Start()

			s.Wait()
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
