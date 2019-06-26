package main

import (
	"contract-service/add"
	"contract-service/contract"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/urfave/cli.v2"
	"os"
	"sort"
)

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

	database := &cli.StringFlag{
		Name:    "database",
		Aliases: []string{"d"},
		Usage:   "set the database pathname",
	}

	return []cli.Flag{
		shell, userpass, json, boot, pin, sync, config, database,
	}

}

func runApp() error {
	flags := globalFlags()
	app := &cli.App{
		Version: "v0.0.1",
		Name:    "seed",
		Usage:   "seed is a video manage tool use ipfs,eth,sqlite3 and so on.",
		Action: func(c *cli.Context) error {
			if quick := c.Bool("q"); quick {
				//add.QuickProcess()
			}

			return nil
		},
		Flags: flags,
	}

	app.Commands = append(app.Commands, contract.CmdContract(app), add.CmdAdd(app))

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := runApp(); err != nil {
		panic(err)
	}
	os.Exit(0)
}
