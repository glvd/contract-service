package bot

import (
	"github.com/godcong/go-trait"
	"github.com/yinhevr/yinhe_bot/message"
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
		//&cli.StringFlag{
		//	Name:    "release",
		//	Value:   "v0.0.1",
		//	Aliases: []string{"r"},
		//	Usage:   "set the application version",
		//},
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
		Name:    "bot",
		Aliases: []string{"B"},
		Usage:   "pin file to ipfs",
		Action: func(context *cli.Context) error {
			message.BootWithGAE(context.String("config"), "8443")
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
