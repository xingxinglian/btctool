package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "btc-tool",
	Short: "BTC 工具集合",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
