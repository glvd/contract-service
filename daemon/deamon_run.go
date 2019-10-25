package daemon

import (
	"os"

	"service"

	"github.com/godcong/go-trait"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdDaemon ...
func CmdDaemon(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:  "config",
			Value: ".service",
			Usage: "set the config path",
		},
	)
	return &cli.Command{
		Before: func(context *cli.Context) error {
			service.DefaultPath = context.String("config")
			if firstRunCheck() {
				_ = os.MkdirAll(service.DefaultPath, 0755)
			}

			return nil
		},
		Name:  "daemon",
		Usage: "daemon handle the target path to process add",
		Action: func(context *cli.Context) error {
			s := service.NewService()
			e := s.Start()
			if e != nil {
				return e
			}
			s.Wait()
			//path := ""
			//if context.NArg() > 0 {
			//	path = context.Args().Get(0)
			//}
			//
			//db := context.String("database")
			//if db == "" {
			//	db = "cs.db"
			//}
			//eng, e := model.InitDB("sqlite3", db)
			//if e != nil {
			//	return e
			//}
			//model.InitMainDB(eng)
			//
			//s := seed.NewSeed(seed.DatabaseOption("sqlite3", context.String("database")))
			//j := context.String("json")
			//if j != "" {
			//	log.Info("json: ", j)
			//	s.Register(seed.Information(j, seed.InfoFlagBSON))
			//}
			//s.Workspace = context.String("workspace")
			//
			//if path != "" {
			//	log.Info("path: ", path)
			//	s.Register(seed.Process(path), seed.Update(seed.UpdateMethodVideo, seed.UpdateContentHash))
			//	s.Scale = context.Int64("scale")
			//}
			//
			//s.Register(seed.ShellOption(context.String("shell")))
			//
			//if context.Bool("skip") {
			//	s.Register(seed.SkipConvertOption())
			//	s.Register(seed.SkipSourceOption())
			//}
			//
			//move := context.String("move")
			//if move == "" || strings.Index(move, path) != -1 {
			//	log.With("move", move, "path", path).Error("output")
			//	panic("cannot handle without move or same with path")
			//}
			//s.Register(seed.MoveProc(move))
			//s.AfterInit(seed.SyncDatabase())
			//s.Workspace = context.String("workspace")
			//
			//for {
			//	s.Start()
			//	s.Wait()
			//	log.With("Handle", path).Info("waiting for next")
			//	time.Sleep(15 * time.Second)
			//}
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}

func firstRunCheck() bool {
	if _, err := os.Stat(service.DefaultPath); err != nil {
		if os.IsNotExist(err) {
			return true
		}
	}
	return false
}
