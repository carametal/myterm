package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "myterm",
		Usage: "Integration cli commands",
		Commands: []*cli.Command{
			{
				Name:    "split-pane",
				Aliases: []string{"split", "sp"},
				Usage:   "Split pane",
				Action: func(cCtx *cli.Context) error {
					err := exec.Command("wezterm", "cli", "split-pane", "--horizontal").Run()
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
