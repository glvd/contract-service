package pin

import (
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed"
	"github.com/yinhevr/seed/model"
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
		Subcommands: nil,
		Flags:       flags,
	}
}
