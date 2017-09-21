package main

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"
)

var RootCmd = &cli.Command{
	Use:   "gomock",
	Short: "A CLI tool for api mock server",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
