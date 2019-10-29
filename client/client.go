package client

import (
	"context"
	"os"
	"path/filepath"

	"service"
	"service/api"
	"service/api/restapi"

	"github.com/goextension/log"
	"gopkg.in/urfave/cli.v2"
)

var _manager *api.Manager

// CmdClientAdd ...
func CmdClientAdd(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:  "thumb",
			Usage: "set the video thumb",
		},
		&cli.StringFlag{
			Name:  "poster",
			Usage: "set the video poster",
		},
		&cli.StringSliceFlag{
			Name:  "sample",
			Usage: "set the sample",
		},
	)
	return &cli.Command{
		Name:          "add",
		Usage:         "add file ipfs",
		UsageText:     "",
		Description:   "",
		ArgsUsage:     "",
		Category:      "",
		ShellComplete: nil,
		Before: func(ctx *cli.Context) error {
			service.DefaultPath = ctx.String("config")
			if firstRunCheck() {
				log.Panic("config file was not found")
			}
			_manager = api.NewManager(context.Background())
			cfg := service.DefaultConfig()
			e := cfg.LoadJSON()
			if e != nil {
				log.Panicw("can't load json file", "error", e)
			}

			rest := restapi.NewRestAPI(cfg.API, restapi.Manager(_manager))
			_manager.RegisterClient(api.RestAPI, rest)
			return nil
		},
		After: nil,
		Action: func(ctx *cli.Context) error {
			work := api.Work{
				WorkStatus: 0,
				VideoPath:  ctx.Args().Slice(),
				PosterPath: ctx.String("poster"),
				ThumbPath:  ctx.String("thumb"),
				SamplePath: ctx.StringSlice("sample"),
				VideoInfo:  ctx.String("json"),
			}

			e := _manager.Client(api.RestAPI).AddWork(_manager, work)
			if e != nil {
				return e
			}
			log.Infof("work added")
			return nil
		},
		Flags: flags,
	}
}

// CmdClientStatus ...
func CmdClientStatus(app *cli.App) *cli.Command {
	flags := append(app.Flags)
	return &cli.Command{
		Name:          "status",
		Usage:         "work status",
		UsageText:     "",
		Description:   "",
		ArgsUsage:     "",
		Category:      "",
		ShellComplete: nil,
		Before: func(ctx *cli.Context) error {
			service.DefaultPath = ctx.String("config")
			if firstRunCheck() {
				log.Panic("config file was not found")
			}
			_manager = api.NewManager(context.Background())
			cfg := service.DefaultConfig()
			e := cfg.LoadJSON()
			if e != nil {
				log.Panicw("can't load json file", "error", e)
			}

			rest := restapi.NewRestAPI(cfg.API, restapi.Manager(_manager))
			_manager.RegisterClient(api.RestAPI, rest)
			return nil
		},
		After: nil,
		Action: func(ctx *cli.Context) error {
			//work := api.Work{
			//	WorkStatus: 0,
			//	VideoPath:  ctx.Args().Slice(),
			//	PosterPath: ctx.String("poster"),
			//	ThumbPath:  ctx.String("thumb"),
			//	SamplePath: ctx.StringSlice("sample"),
			//	VideoInfo:  ctx.String("json"),
			//}

			works, e := _manager.Client(api.RestAPI).GetWorks(_manager)
			if e != nil {
				return e
			}
			log.Infof("works result:%+v", works)
			return nil
		},
		Flags: flags,
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
