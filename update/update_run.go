package update

import (
	"os"
	"path/filepath"

	"github.com/godcong/go-trait"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdAdd ...
func CmdUpdate(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "workspace",
			Aliases: []string{"w"},
			Value:   "",
			Usage:   "putted files",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "set the other sqlite3 path",
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
		&cli.StringFlag{
			Name:  "method",
			Usage: "set the update method",
		},
	)
	return &cli.Command{
		Name:          "update",
		Aliases:       []string{"U"},
		Usage:         "update db",
		UsageText:     "",
		Description:   "",
		ArgsUsage:     "",
		Category:      "",
		ShellComplete: nil,
		Before:        nil,
		After:         nil,
		Action: func(context *cli.Context) error {
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
			//
			////path := context.String("from")
			//if path != "" {
			//	log.Info("path: ", path)
			//}
			//
			//method := seed.UpdateMethodAll
			//switch context.String("method") {
			//case "video":
			//	method = seed.UpdateMethodVideo
			//case "unfinished":
			//	method = seed.UpdateMethodUnfinished
			//}
			//
			//s.Register(seed.Update(method, seed.UpdateContentHash))
			//s.AfterInit(seed.SyncDatabase())
			//s.Start()
			//
			//s.Wait()
			//TODO
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
