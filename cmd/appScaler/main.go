package main

import (
	"github.com/olisajc/appScaler/pkg/cli"
)

func main() {

	rootCmd := cli.Root()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
