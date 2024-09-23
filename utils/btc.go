package utils

import (
	"btctool/config" // 导入配置模块
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

func GetUTXOCount(address string) (int, error) {
	// 从配置模块获取节点配置
	cfg := config.GetBitcoinConfig()

	// 使用配置创建 RPC 客户端
	connCfg := &rpcclient.ConnConfig{
		Host:         cfg.Node.Host,
		User:         cfg.Node.User,
		Pass:         cfg.Node.Pass,
		HTTPPostMode: cfg.Node.HTTPPostMode,
		DisableTLS:   cfg.Node.DisableTLS,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return 0, fmt.Errorf("创建 RPC 客户端失败: %v", err)
	}
	defer client.Shutdown()

	// 解析地址
	addr, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err != nil {
		return 0, fmt.Errorf("解析地址失败: %v", err)
	}

	// 获取未花费的交易输出
	unspent, err := client.ListUnspentMinMaxAddresses(0, 9999999, []btcutil.Address{addr})
	if err != nil {
		return 0, fmt.Errorf("获取未花费交易输出失败: %v", err)
	}

	return len(unspent), nil
}
