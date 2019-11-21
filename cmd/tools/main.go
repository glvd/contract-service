package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = "v0.0.1"
	app.Name = applicationName

	app.Commands = []*cli.Command{
		{
			Name:   "deploy",
			Before: deployBefore(),
			Action: deployAction(),
			Flags:  deployFlags(),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
