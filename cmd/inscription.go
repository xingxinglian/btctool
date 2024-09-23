package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inscriptionCmd = &cobra.Command{
	Use:   "inscribe [data]",
	Short: "打铭文",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := args[0]
		// 这里应该是实际的铭文逻辑
		fmt.Printf("铭文数据: %s\n", data)
	},
}

func init() {
	rootCmd.AddCommand(inscriptionCmd)
}
