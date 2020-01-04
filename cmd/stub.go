package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func StubCommand() *cli.Command {
	return &cli.Command{
		Name:   "stub",
		Usage:  "",
		Action: Stub,
		Flags:  []cli.Flag{},
	}
}

func Stub(c *cli.Context) error {
	logrus.Debug("stub function")
	return nil
}
