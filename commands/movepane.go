package commands

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func callMovePane(args ...string) error {
	err := exec.Command("wezterm", args...).Run()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

var MovePane = &cli.Command{
	Name:    "move-pane",
	Aliases: []string{"mp"},
	Usage:   "Move a pane to a new position",
	Subcommands: []*cli.Command{
		{
			Name:    "up",
			Aliases: []string{"u"},
			Usage:   "Move the pane up",
			Action: func(c *cli.Context) error {
				return callMovePane("cli", "activate-pane-direction", "Up")
			},
		},
		{
			Name:    "down",
			Aliases: []string{"d"},
			Usage:   "Move the pane down",
			Action: func(c *cli.Context) error {
				return callMovePane("cli", "activate-pane-direction", "Down")
			},
		},
		{
			Name:    "left",
			Aliases: []string{"l"},
			Usage:   "Move the pane left",
			Action: func(c *cli.Context) error {
				return callMovePane("cli", "activate-pane-direction", "Left")
			},
		},
		{
			Name:    "right",
			Aliases: []string{"r"},
			Usage:   "Move the pane right",
			Action: func(c *cli.Context) error {
				return callMovePane("cli", "activate-pane-direction", "Right")
			},
		},
	},
}
