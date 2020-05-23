package cmd

import (
	"github.com/luthermonson/urfave-cli-skeleton/cmd/example"
	"github.com/urfave/cli/v2"
)

func Commands() cli.Commands {
	return cli.Commands{
		example.Command,
	}
}
