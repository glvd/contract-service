package contract

import (
	"gopkg.in/urfave/cli.v2"
	"math/big"
)

// CmdContract ...
func CmdContract(app *cli.App) *cli.Command {
	flags := append(app.Flags,
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			Value:   "",
			Usage:   "contract address",
		},
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "video",
			Usage:   "contract process type",
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
		&cli.IntFlag{
			Name:  "limit",
			Usage: "set the ban max numbers",
		},
		&cli.Int64Flag{
			Name:  "code",
			Usage: "set the version code for update",
		},
		&cli.StringFlag{
			Name:    "key",
			Usage:   "set the ct process key",
			EnvVars: []string{"seedKey"},
		},
	)
	return &cli.Command{
		Name:    "contract",
		Aliases: []string{"C"},
		Usage:   "contract share the video info to ct.",
		Action: func(context *cli.Context) error {
			log.Info("contract call")
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
			code := context.Int64("code")
			contract := NewContract(key, address)
			if contract == nil {
				panic("null contract")
			}
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
				e := contract.CheckNameExists(context.String("ban"), checkList...)
				if e != nil {
					log.Error(e)
					return e
				}
			case "hot":
			case "status":
				var ver string
				var lastHash string
				var e error
				if code > 0 {
					ver, e = contract.GetCodeVersion(big.NewInt(code))
					if e != nil {
						log.Error(e)
						return e
					}
					log.With("version", ver, "hash", lastHash, "code", code).Info("version:", ver)
					return nil
				}
				ver, lastHash, e = contract.GetLastVersionHash()
				if e != nil {
					log.Error(e)
					return e
				}
				log.With("version", ver, "hash", lastHash, "code", code).Info("last")
				return nil
			case "app":
				if path != "" {
					if err := contract.UpdateAppWithPath(code, version, path); err != nil {
						log.Error(err)
						return err
					}
					return nil
				}
				e := contract.UpdateApp(code, version, hash)
				if e != nil {
					log.Error(e)
					return e
				}
				return nil
			}
			return nil
		},
		Subcommands: nil,
		Flags:       flags,
	}
}
