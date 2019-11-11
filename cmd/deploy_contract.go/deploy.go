package main

import (
	"github.com/urfave/cli/v2"
)

const applicationName = "deploy"

func main() {
	app := cli.NewApp()
	app.Version = "v0.0.1"
	app.Name = applicationName

	//Usage:   "service is a video manage tool use ipfs,eth,sqlite3 and so on.",
	//Action: func(c *cli.Context) error {
	//return nil
	//},
	//	Flags: flags,
	//}

}
