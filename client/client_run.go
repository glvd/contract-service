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

// CmdClient ...
func CmdClient(app *cli.App) *cli.Command {
	flags := append(app.Flags)
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
			//json := ctx.String("json")
			cfg := service.DefaultConfig()
			e := cfg.LoadJSON()
			if e != nil {
				log.Panicw("can't load json file", "error", e)
			}

			rest := restapi.NewRestAPI(cfg.API, restapi.Manager(_manager))

			//go func() {
			//	e := rest.Start()
			//	if e != nil {
			//		log.Errorw("start rest client failed", "error", e)
			//	}
			//}()
			_manager.RegisterClient(api.RestAPI, rest)
			//rpcClient := rpcclient.NewClient(cfg.API)
			//go func() {
			//	e := rpcClient.Start()
			//	if e != nil {
			//		log.Errorw("start rpc client failed", "error", e)
			//	}
			//}()
			//_manager.RegisterClient(api.RPCClient, rpcClient)

			return nil
		},
		After: nil,
		Action: func(ctx *cli.Context) error {
			works, e := _manager.Client(api.RestAPI).GetWorks(_manager)
			if e != nil {
				return e
			}
			log.Infof("%+v", works)
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
