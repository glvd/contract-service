package contract

import (
	"github.com/yinhevr/seed"
	"gopkg.in/urfave/cli.v2"
)

// CmdContract ...
func CmdContract(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			Value:   "",
			Usage:   "ct address",
		},
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "video",
			Usage:   "ct process type",
		},
		&cli.StringFlag{
			Name:  "ban",
			Usage: "ban no to check",
		},
		&cli.StringFlag{
			Name:    "release",
			Value:   "v0.0.1",
			Aliases: []string{"r"},
			Usage:   "set the application version",
		},
		&cli.StringFlag{
			Name:  "hash",
			Usage: "set the app ipfs hash",
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "set the app path to add  hash",
		},
		&cli.StringFlag{
			Name:  "limit",
			Usage: "set the ban max numbers",
		},
		&cli.StringFlag{
			Name:    "key",
			Usage:   "set the ct process key",
			EnvVars: []string{"seedKey"},
		},
	)
	return &cli.Command{
		Name:    "ct",
		Aliases: []string{"C"},
		Usage:   "ct share the video info to ct.",
		Action: func(context *cli.Context) error {
			log.Info("ct call")
			key := context.String("key")
			if key == "" {
				panic("key must set use -key or Env(seedKey)")
			}
			address := context.String("a")
			if address == "" {
				panic("address must set use -address,-a")
			}
			version := context.String("release")
			path := context.String("path")
			hash := context.String("hash")
			switch context.String("t") {
			case "video":
				log.Info("video:", context.String("ban"))
				//return seedProc(key, address)
			case "check":
				ban := context.String("ban")
				if ban == "" {
					return nil
				}

				limit := context.Int("limit")
				var checkList []int
				for i := 0; i <= limit; i++ {
					checkList = append(checkList, i)
				}
				log.With("ban", context.String("ban"), "limit", limit).Info("check")
				e := seed.CheckNameExists(context.String("ban"), checkList...)
				if e != nil {
					log.Error(e)
					return e
				}
			case "hot":

			case "app":
				if path != "" {
					if err := UpdateAppWithPath(version, path); err != nil {
						log.Error(err)
						return err
					}
					return nil
				} else if hash != "" {
					e := seed.UpdateApp(version, hash)
					if e != nil {
						log.Error(e)
						return e
					}
				} else {
					ver, lastHash, e := seed.GetLastVersionHash()
					if e != nil {
						log.Error(e)
						return e
					}
					log.With("version", ver, "hash", lastHash).Info("last")
				}

			}

			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
