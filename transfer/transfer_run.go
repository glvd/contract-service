package transfer

import (
	"github.com/godcong/go-trait"
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
		&cli.BoolFlag{
			Name:  "showsql",
			Usage: "show sql",
		},
	)
	return &cli.Command{
		Name:  "transfer",
		Usage: "transfer from other or to other database",
		Action: func(context *cli.Context) error {
			//db := context.String("database")
			//if db == "" {
			//	db = "cs.db"
			//}
			////eng, e := model.InitDB("sqlite3", db)
			////if e != nil {
			////	return e
			////}
			////model.InitMainDB(eng)
			//
			//path := context.String("from")
			//status := seed.TransferStatusFromOther
			//flag := seed.InfoFlagSQLite
			//switch context.String("status") {
			//case "old":
			//	status = seed.TransferStatusFromOld
			//case "json":
			//	path = context.String("to")
			//	status = seed.TransferStatusToJSON
			//	flag = seed.InfoFlagJSON
			//}
			//
			//s := seed.NewSeed(seed.DatabaseOption("sqlite3", db), seed.Transfer(path, flag, status))
			//
			//if context.Bool("showsql") {
			//	s.AfterInit(seed.ShowSQLOption())
			//}
			//
			//s.AfterInit(seed.SyncDatabase())
			//s.Start()
			//
			//s.Wait()
			//TODO
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
