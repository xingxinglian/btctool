package main

import (
	"btctool/cmd"
	"btctool/config"
	"log"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	cmd.Execute()
}
