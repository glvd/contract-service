package daemon

import (
	"strings"
	"time"

	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed"
	"github.com/yinhevr/seed/model"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdDaemon ...
func CmdDaemon(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "run",
			Aliases: []string{"a"},
			Value:   "",
			Usage:   "contract address",
		},
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "video",
			Usage:   "contract process type",
		},
		&cli.StringFlag{
			Name:  "ban",
			Usage: "ban no to check",
		},
		&cli.StringFlag{
			Name:    "release",
			Value:   "v0.0.1",
			Aliases: []string{"r"},
			Usage:   "set the application version",
		},
		&cli.StringFlag{
			Name:  "hash",
			Usage: "set the app ipfs hash",
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "set the app path to add  hash",
		},
		&cli.IntFlag{
			Name:  "limit",
			Usage: "set the ban max numbers",
		},
		&cli.Int64Flag{
			Name:  "code",
			Usage: "set the version code for update",
		},
		&cli.StringFlag{
			Name:    "key",
			Usage:   "set the ct process key",
			EnvVars: []string{"seedKey"},
		},
		&cli.StringFlag{
			Name:  "move",
			Usage: "move to",
		},
		&cli.IntFlag{
			Name:  "scale",
			Usage: "set scale value",
		},
		&cli.StringFlag{
			Name:    "workspace",
			Aliases: []string{"w"},
			Value:   "",
			Usage:   "putted files",
		},
	)
	return &cli.Command{
		Name:    "daemon",
		Aliases: []string{"D"},
		Usage:   "daemon handle the target path to process add",
		Action: func(context *cli.Context) error {
			path := ""
			if context.NArg() > 0 {
				path = context.Args().Get(0)
			}

			db := context.String("database")
			if db == "" {
				db = "cs.db"
			}
			eng, e := model.InitDB("sqlite3", db)
			if e != nil {
				return e
			}
			model.InitMainDB(eng)

			s := seed.NewSeed(seed.DatabaseOption("sqlite3", context.String("database")))
			j := context.String("json")
			if j != "" {
				log.Info("json: ", j)
				s.Register(seed.Information(j, seed.InfoFlagBSON))
			}
			s.Workspace = context.String("workspace")

			if path != "" {
				log.Info("path: ", path)
				s.Register(seed.Process(path), seed.Update(seed.UpdateMethodVideo, seed.UpdateContentHash))
				s.Scale = context.Int64("scale")
			}

			s.Register(seed.ShellOption(context.String("shell")))

			if context.Bool("skip") {
				s.Register(seed.SkipConvertOption())
				s.Register(seed.SkipSourceOption())
			}

			move := context.String("move")
			if move == "" || strings.Index(move, path) != -1 {
				log.With("move", move, "path", path).Error("output")
				panic("cannot handle without move or same with path")
			}
			s.Register(seed.MoveProc(move))
			s.AfterInit(seed.SyncDatabase())
			s.Workspace = context.String("workspace")

			for {
				s.Start()
				s.Wait()
				log.With("Handle", path).Info("waiting for next")
				time.Sleep(15 * time.Second)
			}
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
