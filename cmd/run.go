package cmd

import (
	"fmt"
	"os"

	"github.com/hlts2/gomock/pkg/gomock"

	cli "github.com/spf13/cobra"
)

var runCmd = &cli.Command{
	Use:   "run",
	Short: "Start API mock server",
	Run: func(cmd *cli.Command, args []string) {
		if err := run(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

var configPath string

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&configPath, "set", "s", "config.yml", "set config file")
}

func run(cmd *cli.Command, args []string) error {
	var config gomock.Config

	err := gomock.LoadConfig(configPath, &config)
	if err != nil {
		return err
	}

	mockServer := gomock.NewServer(config)

	return mockServer.Launch()
}
