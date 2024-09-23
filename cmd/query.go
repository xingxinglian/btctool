package cmd

import (
	"btctool/utils" // 导入新的 utils 模块
	"fmt"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query [address]",
	Short: "查询地址的 UTXO 数量",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		utxoCount, err := utils.GetUTXOCount(address) // 调用 utils 中的函数
		if err != nil {
			fmt.Printf("查询失败: %v\n", err)
			return
		}
		fmt.Printf("地址 %s 的 UTXO 数量为: %d\n", address, utxoCount)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
