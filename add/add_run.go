package add

import (
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdAdd ...
func CmdAdd(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "workspace",
			Aliases: []string{"w"},
			Value:   "",
			Usage:   "putted files",
		},
		&cli.BoolFlag{
			Name:  "skip",
			Value: false,
			Usage: "skip convert if not h264/aac",
		},
		//&cli.StringFlag{
		//	Name:  "ban",
		//	Usage: "ban no to check",
		//},
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
		Name:    "add",
		Aliases: []string{"A"},
		Usage:   "add file to db",
		Action: func(context *cli.Context) error {
			path := ""
			if context.NArg() > 0 {
				path = context.Args().Get(0)
			}
			db := context.String("database")
			if db == "" {
				db = "cs.db"
			}
			s := seed.NewSeed(seed.DatabaseOption("sqlite3", db))
			j := context.String("json")
			if j != "" {
				s.Register(seed.Information(context.String("json"), seed.InfoFlagBSON))
			}
			if path != "" {
				s.Register(seed.Process(path))
			}

			s.Register(seed.ShellOption(context.String("shell")))

			if context.Bool("skip") {
				s.Register(seed.SkipConvertOption())
			}

			s.AfterInit(seed.SyncDatabase())
			s.Workspace = context.String("workspace")
			s.Start()

			s.Wait()
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
