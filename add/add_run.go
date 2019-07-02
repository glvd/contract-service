package add

import (
	"github.com/yinhevr/seed/model"
	"os"
	"path/filepath"

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
			}

			if path != "" {
				log.Info("path: ", path)
				s.Register(seed.Process(path), seed.Update(seed.UpdateMethodVideo, seed.UpdateContentHash))
			}

			s.Register(seed.ShellOption(context.String("shell")))

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
