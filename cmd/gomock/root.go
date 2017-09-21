package main

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"
)

var rootCmd = &cli.Command{
	Use:   "gomock",
	Short: "A CLI tool for api mock server",
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
