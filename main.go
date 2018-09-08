package main

import (
	"log"
	"os"

	"github.com/hlts2/gomock/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomock"
	app.Usage = "API mock server"
	app.Version = "0.0.1"
	app.Commands = cli.Commands{
		cmd.RunCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
