package main

import (
	"log"
	"os"

	commands "github.com/carametal/myterm/commnads"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "myterm",
		Usage: "Integration cli commands",
		Commands: []*cli.Command{
			commands.SplitPane,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
