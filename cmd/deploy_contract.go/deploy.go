package main

import (
	"path/filepath"

	"github.com/glvd/conversion"
	"github.com/urfave/cli/v2"
)

const applicationName = "deploy"

func main() {
	app := cli.NewApp()
	app.Version = "v0.0.1"
	app.Name = applicationName
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "load deploy config",
			Value: filepath.Join(DefaultConfigPath, DefaultConfigPath),
		},
		&cli.StringFlag{
			Name:  "",
			Usage: "",
			Value: "",
		},
	}
	app.Before = before()
}

func before() cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		LoadConfig()
	}
}
