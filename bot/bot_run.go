package bot

import (
	"github.com/glvd/bot/message"
	"github.com/godcong/go-trait"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdPin ...
func CmdBot(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:  "status",
			Usage: "set the pin status (all/relate/hash/unfinished) default is all",
		},
		&cli.StringFlag{
			Name:  "port",
			Value: "443",
			Usage: "set the bot handle port",
		},
		//&cli.StringFlag{
		//	Name:  "hash",
		//	Usage: "set the app ipfs hash",
		//},
		//&cli.StringFlag{
		//	Name:  "path",
		//	Value: "config.json",
		//	Usage: "set the config file path to run",
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
		Name:  "bot",
		Usage: "pin file to ipfs",
		Action: func(context *cli.Context) error {
			db := context.String("database")
			if db == "" {
				db = "cs.db"
			}
			e := message.InitDB(db)
			if e != nil {
				return e
			}
			//model.InitMainDB(eng)

			message.BootWithUpdate(context.String("config"))
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
