package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "endermite",
		Usage: "complex minecraft launcher to use from your favorite terminal",
		Commands: []*cli.Command{
			{
				Name:   "list",
				Action: list,
				Usage:  "lists all minecraft versions to be downloaded. If no version type flag is provided, will list releases by default",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "description",
						Aliases: []string{"d", "desc"},
						Usage:   "shows more info about versions in form of table",
					},
					&cli.BoolFlag{
						Name:    "release",
						Aliases: []string{"r", "rel"},
						Usage:   "include releases while listing",
					},
					&cli.BoolFlag{
						Name:    "snapshot",
						Aliases: []string{"s", "snap"},
						Usage:   "include snapshots while listing",
					},
					&cli.BoolFlag{
						Name:  "beta",
						Usage: "include beta versions while listing",
					},
					&cli.BoolFlag{
						Name:  "alpha",
						Usage: "include alpha versions while listing",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
