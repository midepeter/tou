package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type Version string

var VersionCmd = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "prints the version of work queue manager",
	Action: func(cctx *cli.Context) error {
		fmt.Println(Version("1.00.1"))
		return nil
	},
}
