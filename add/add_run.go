package add

import (
	"os"
	"path/filepath"

	"github.com/glvd/seed/model"

	"github.com/glvd/seed"
	"github.com/godcong/go-trait"
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

		&cli.IntFlag{
			Name:  "scale",
			Usage: "set scale value",
		},
		&cli.StringFlag{
			Name:  "infomove",
			Usage: "move the success info file to ...",
		},
		&cli.BoolFlag{
			Name:  "fix",
			Usage: "fix video",
		},
		&cli.BoolFlag{
			Name:  "exist",
			Usage: "set bool to skip exist file",
		},
		&cli.BoolFlag{
			Name:  "source",
			Usage: "set bool to skip add source file",
		},
		&cli.BoolFlag{
			Name:  "nocheck",
			Usage: "set bool to check added",
		},
		&cli.IntFlag{
			Name:  "limit",
			Usage: "set the max process limit",
		},
		//&cli.StringFlag{
		//	Name:    "key",
		//	Usage:   "set the ct process key",
		//	EnvVars: []string{"seedKey"},
		//},
	)
	return &cli.Command{
		Name:          "add",
		Aliases:       []string{"A"},
		Usage:         "add file to db",
		UsageText:     "",
		Description:   "",
		ArgsUsage:     "",
		Category:      "",
		ShellComplete: nil,
		Before:        nil,
		After:         nil,
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
				s.Register(seed.Information(j, seed.InfoFlagBSON, getList(path)...))
				s.MaxLimit = context.Int("limit")
			}

			imove := context.String("infomove")
			if imove != "" {
				log.Info("infomove", imove)
				s.Register(seed.MoveInfo(imove))
			}

			if path != "" {
				log.Info("path: ", path)
				s.Register(seed.Process(path), seed.Update(seed.UpdateMethodVideo, seed.UpdateContentHash))
				s.Scale = context.Int64("scale")
			}

			s.Register(seed.ShellOption(context.String("shell")))

			if context.Bool("exist") {
				s.Register(seed.SkipExistOption())
			}
			if context.Bool("source") {
				s.Register(seed.SkipSourceOption())
			}

			if context.Bool("skip") {
				s.Register(seed.SkipConvertOption())
			}

			if context.Bool("nocheck") {
				s.NoCheck = true
			}

			s.AfterInit(seed.SyncDatabase())
			s.Workspace = context.String("workspace")
			s.Start()

			s.Wait()
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
	}
}

func getList(path string) (list []string) {
	if path == "" {
		return
	}
	files := getFilters(path)
	if files == nil {
		return
	}
	for key := range getFilters(path) {
		list = append(list, key)
	}
	return
}

func getFilters(path string) (files map[string]string) {
	files = make(map[string]string)
	info, e := os.Stat(path)
	if e != nil {
		return nil
	}
	if info.IsDir() {
		file, e := os.Open(path)
		if e != nil {
			return nil
		}
		defer file.Close()
		names, e := file.Readdirnames(-1)
		if e != nil {
			return nil
		}
		var fullPath string
		for _, name := range names {
			fullPath = filepath.Join(path, name)
			tmp := getFilters(fullPath)
			if tmp != nil {
				for key := range tmp {
					files[key] = ""
				}
			}
		}
		return files
	}
	_, file := filepath.Split(path)
	return map[string]string{onlyName(file): ""}
}

func onlyName(name string) string {
	_, name = filepath.Split(name)
	for i := len(name) - 1; i >= 0 && !os.IsPathSeparator(name[i]); i-- {
		if name[i] == '.' {
			return name[:i]
		}
	}
	return ""
}
