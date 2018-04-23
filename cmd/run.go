package cmd

import (
	"fmt"
	"os"

	"github.com/hlts2/gomock/pkg/gomock"
	"gopkg.in/yaml.v2"

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

var configFileName string

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&configFileName, "set", "s", "config.yml", "set config file")
}

func run(cmd *cli.Command, args []string) error {
	reader, err := os.Open(configFileName)
	if err != nil {
		return err
	}

	var config gomock.Config

	err = yaml.NewDecoder(reader).Decode(&config)
	if err != nil {
		return err
	}

	mockServer := gomock.NewServer(config)

	return mockServer.Launch()
}
