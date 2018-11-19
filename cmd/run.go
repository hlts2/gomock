package cmd

import (
	"path/filepath"

	gmk "github.com/hlts2/gomock/pkg/gomock"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// RunCommand is run command
func RunCommand() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "start API mock server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "set, s",
				Usage: "config file",
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
				return errors.Wrap(err, "faild to load configuration file")
			}

			server := gmk.NewServer(&config)

			dir := ctxt.String("tls-path")
			if len(dir) == 0 {
				err := server.Serve()
				if err != nil {
					return errors.Wrap(err, "faild to run server")
				}
				return nil
			}

			var (
				crtPath = filepath.Join(dir, "server.crt")
				keyPath = filepath.Join(dir, "server.key")
			)

			err = server.ServeTLS(crtPath, keyPath)
			if err != nil {
				return errors.Wrap(err, "faild to run server")
			}
			return nil
		},
	}
}
