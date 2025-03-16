package commands

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func f(args ...string) error {
	err := exec.Command("wezterm", args...).Run()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

var splitPaneHorizontal = &cli.Command{
	Name:    "horizontal",
	Aliases: []string{"h"},
	Usage:   "Split pane horizontal",
	Action: func(cCtx *cli.Context) error {
		return f("cli", "split-pane", "--horizontal")
	},
}

var splitPaneVertical = &cli.Command{
	Name:    "vertical",
	Aliases: []string{"v"},
	Usage:   "Split pane vertical",
	Action: func(cCtx *cli.Context) error {
		return f("cli", "split-pane")
	},
}

var SplitPane = &cli.Command{
	Name:     "split-pane",
	Aliases:  []string{"split", "sp"},
	Usage:    "Split pane in the terminal",
	HideHelp: true,
	Subcommands: []*cli.Command{
		splitPaneHorizontal,
		splitPaneVertical,
	},
}
