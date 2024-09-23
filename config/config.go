package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// NodeConfig 包含节点的配置
type NodeConfig struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Pass         string `yaml:"pass"`
	HTTPPostMode bool   `yaml:"http_post_mode"`
	DisableTLS   bool   `yaml:"disable_tls"`
}

// APIConfig 包含 API 的配置
type APIConfig struct {
	URL string `yaml:"url"`
}

// BitcoinConfig 包含整体的比特币配置
type BitcoinConfig struct {
	Mode string     `yaml:"mode"`
	Node NodeConfig `yaml:"node"`
	API  APIConfig  `yaml:"api"`
}

var config BitcoinConfig

// LoadConfig 加载配置文件
func LoadConfig() error {
	configFile := "conf.yaml" // 默认配置文件

	// 根据环境变量选择配置文件
	if os.Getenv("BTC_ENV") == "test" {
		configFile = "conf_test.yaml"
	}

	// 读取配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("无法读取配置文件 %s: %v", configFile, err)
	}

	// 解析配置文件
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("无法解析配置文件 %s: %v", configFile, err)
	}

	return nil
}

// GetBitcoinConfig 返回当前的比特币配置
func GetBitcoinConfig() BitcoinConfig {
	return config
}
