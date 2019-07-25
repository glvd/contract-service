package main

import (
	"contract-service/daemon"
	"contract-service/transfer"
	"contract-service/update"
	"github.com/godcong/go-trait"
	"os"
	"sort"

	"contract-service/add"
	"contract-service/bot"
	"contract-service/contract"
	"contract-service/pin"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yinhevr/seed/model"
	"gopkg.in/urfave/cli.v2"
)

var log = trait.NewZapSugar()

func globalFlags() []cli.Flag {
	shell := &cli.StringFlag{
		Name:        "shell",
		Aliases:     []string{"s"},
		Value:       "localhost:5001",
		Usage:       "set the ipfs api address:port",
		DefaultText: "localhost:5001",
	}

	config := &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "sometime need set config",
	}

	userpass := &cli.StringFlag{
		Name:    "userpass",
		Aliases: []string{"u"},
		Usage:   "set the database user:password",
	}

	json := &cli.StringFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "set the json file path",
	}

	boot := &cli.StringFlag{
		Name:    "bootstrap",
		Aliases: []string{"b"},
		Usage:   "set the ipfs bootstrap swarm address to quick connect",
	}

	pin := &cli.BoolFlag{
		Name:    "pin",
		Aliases: []string{"p"},
		Usage:   "set to pin on ipfs",
	}

	sync := &cli.BoolFlag{
		Name:    "sync",
		Aliases: []string{"n"},
		Usage:   "check if the video is synced",
	}

	skip := &cli.BoolFlag{
		Name:  "skip",
		Value: false,
		Usage: "skip convert if not h264/aac",
	}

	database := &cli.StringFlag{
		Name:    "database",
		Value:   "cs.db",
		Aliases: []string{"d"},
		Usage:   "set the database pathname",
	}

	return []cli.Flag{
		shell, userpass, json, boot, pin, sync, config, database, skip,
	}

}

func runApp() error {
	flags := globalFlags()
	app := &cli.App{
		Version: "v0.0.2",
		Name:    "service",
		Usage:   "service is a video manage tool use ipfs,eth,sqlite3 and so on.",
		Action: func(c *cli.Context) error {
			log.Info("main call")
			eng, e := model.InitDB("sqlite3", c.String("database"))
			if e != nil {
				return e
			}
			model.InitMainDB(eng)
			return nil
		},
		Flags: flags,
	}

	app.Commands = []*cli.Command{
		contract.CmdContract(app),
		add.CmdAdd(app),
		pin.CmdPin(app),
		bot.CmdBot(app),
		daemon.CmdDaemon(app),
		transfer.CmdTransfer(app),
		update.CmdUpdate(app),
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
