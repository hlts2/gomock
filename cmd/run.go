package cmd

import (
	"path/filepath"

	gmk "github.com/hlts2/gomock/pkg/gomock"
	"github.com/urfave/cli"
)

// RunCommand is run command
func RunCommand() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "serve API mock server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "set, s",
				Usage: "configuration file",
				Value: "config.yml",
			},
			cli.StringFlag{
				Name:  "tls-path",
				Usage: "directory to the TLS server.crt/server.key file",
				Value: "",
			},
		},
		Action: func(ctxt *cli.Context) error {
			var config gmk.Config

			err := gmk.LoadConfig(ctxt.String("set"), &config)
			if err != nil {
				return err
			}

			server, err := gmk.NewServer(config.Endpoints)
			if err != nil {
				return err
			}

			dir := ctxt.String("tls-path")
			if len(dir) == 0 {
				return server.Start(":" + config.Port)
			}

			var (
				crt = filepath.Join(dir, "server.crt")
				key = filepath.Join(dir, "server.key")
			)

			return server.StartTLS(":"+config.Port, crt, key)
		},
	}
}
