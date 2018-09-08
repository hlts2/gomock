package cmd

import (
	gmk "github.com/hlts2/gomock/pkg/gomock"
	"github.com/urfave/cli"
)

func RunCommand() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "serve API mock server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "set, s",
				Usage: "set config file",
				Value: "config.yml",
			},
		},
		Action: func(ctxt *cli.Context) error {
			var config gmk.Config

			err := gmk.LoadConfig(ctxt.String("set"), &config)
			if err != nil {
				return err
			}

			return gmk.NewServer(config).Serve()
		},
	}
}
