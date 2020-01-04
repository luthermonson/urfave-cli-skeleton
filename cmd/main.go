package cmd

import (
	"github.com/urfave/cli/v2"
)

func Commands() cli.Commands {
	return cli.Commands{
		StubCommand(),
	}
}
