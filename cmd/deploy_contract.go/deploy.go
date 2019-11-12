package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const applicationName = "deploy"

func init() {
	dir, e := os.Getwd()
	if e != nil {
		dir := os.Getenv("HOME")
		if dir == "" {
			usr, err := user.Current()
			if err != nil {
				panic(fmt.Sprintf("cannot get current user: %s", err))
			}
			dir = usr.HomeDir
		}
	}
	ConfigPath = filepath.Join(dir, DefaultConfigPath, DefaultConfigName)
}

func main() {
	app := cli.NewApp()
	app.Version = "v0.0.1"
	app.Name = applicationName
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "load deploy config",
			Value: ConfigPath,
		},
		&cli.StringFlag{
			Name:  "",
			Usage: "",
			Value: "",
		},
	}
	app.Before = before()

	app.Action = action()

}

func before() cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		config, e := LoadConfig(ctx.String("config"))
		if e != nil {
			return e
		}
		_config = config

		return nil
	}
}

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {

	}
}
