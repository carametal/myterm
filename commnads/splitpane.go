package commands

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var SplitPane = &cli.Command{
	Name:     "split-pane",
	Aliases:  []string{"split", "sp"},
	Usage:    "Split pane",
	HideHelp: true,
	Subcommands: []*cli.Command{
		{
			Name:    "horizontal",
			Aliases: []string{"h"},
			Usage:   "Split pane horizontal",
			Action: func(cCtx *cli.Context) error {
				err := exec.Command("wezterm", "cli", "split-pane", "--horizontal").Run()
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
		{
			Name:    "vertical",
			Aliases: []string{"v"},
			Usage:   "Split pane vertical",
			Action: func(cCtx *cli.Context) error {
				err := exec.Command("wezterm", "cli", "split-pane").Run()
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	},
}
