package add

import (
	"os"
	"path/filepath"

	"github.com/glvd/seed"
	"github.com/godcong/go-trait"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

// CmdAdd ...
func CmdAdd(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"O"},
			//Value:   "",
			Usage: "slice output path",
		},

		&cli.IntFlag{
			Name:  "scale",
			Value: 720,
			Usage: "set scale value",
		},
		&cli.StringFlag{
			Name:  "resource",
			Usage: "set the resource path",
		},
		&cli.IntFlag{
			Name:  "limit",
			Value: 5000,
			Usage: "set the max process limit",
		},
		&cli.StringFlag{
			Name:  "skip",
			Usage: "skip type",
		},
	)
	return &cli.Command{
		Name:          "add",
		Usage:         "add file to db",
		UsageText:     "",
		Description:   "",
		ArgsUsage:     "",
		Category:      "",
		ShellComplete: nil,
		Before:        nil,
		After:         nil,
		Action: func(context *cli.Context) error {

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

func toScale(i int64) seed.Scale {
	if i == 480 {
		return seed.LowScale
	} else if i == 1080 {
		return seed.MiddleScale
	} else {

	}
	return seed.MiddleScale
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
