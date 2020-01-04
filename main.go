package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/luthermonson/urfave-cli-skeleton/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "tail-gunner",
		Usage:    "When tail needs a little more firepower to cover your six",
		Action:   cmd.Stub,
		Commands: cmd.Commands(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Turn on verbose debug logging",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Turn on off all logging",
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
			}
			if ctx.Bool("quiet") {
				logrus.SetOutput(ioutil.Discard)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
