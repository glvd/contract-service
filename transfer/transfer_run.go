package transfer

import (
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed"
	"github.com/yinhevr/seed/model"
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
	)
	return &cli.Command{
		Name:    "pin",
		Aliases: []string{"P"},
		Usage:   "pin file to ipfs",
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

			from := context.String("from")
			s := seed.NewSeed(seed.DatabaseOption("sqlite3", db), seed.Transfer(from, seed.InfoFlagSQLite, seed.TransferStatusOther))

			s.Start()

			s.Wait()
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
