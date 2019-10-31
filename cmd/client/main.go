//go:generate statik -f -src=./api/restapi
package main

import (
	"os"
	"sort"

	"service/client"

	"github.com/goextension/log"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	logger, e := zap.NewProduction(zap.AddCallerSkip(1))
	if e != nil {
		panic(e)
	}
	log.Register(logger.Sugar())
}

func globalFlags() []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: ".service",
			Usage: "set the config path",
		},
		&cli.StringFlag{
			Name:  "json",
			Value: "",
			Usage: "load video json",
		},
	}

}

func runApp() error {
	flags := globalFlags()
	app := &cli.App{
		Version: "v0.0.2",
		Name:    "service",
		Usage:   "service is a video manage tool use ipfs,eth,sqlite3 and so on.",
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags: flags,
	}

	app.Commands = []*cli.Command{
		client.CmdClientAdd(app),
		client.CmdClientRemove(app),
		client.CmdClientStatus(app),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	if err := runApp(); err != nil {
		panic(err)
	}
	os.Exit(0)
}
