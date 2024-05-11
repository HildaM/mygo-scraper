package config

import (
	"MyGo-scraper/common"
	"fmt"
	"os"
	"path/filepath"

	l "MyGo-scraper/common/logger"

	"gopkg.in/yaml.v3"
)

var AppConfig *Config

func Init() {
	AppConfig = &Config{}
	env := common.GetEnv()
	rootDir := common.GetRootDir()

	configPath := filepath.Join(rootDir, "config", fmt.Sprintf("%s.yaml", env))
	l.Logger.Info(fmt.Sprintf("Init config file, filepath: %s", configPath))

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("Init config file error, filepath: %s, err: %v", configPath, err))
		panic(err)
	}

	// 解析配置文件
	configContent := []byte(os.ExpandEnv(string(data)))
	if err := yaml.Unmarshal(configContent, AppConfig); err != nil {
		l.Logger.Error(fmt.Sprintf("Parse config file error, filepath: %s, err: %v", configPath, err))
		panic(err)
	}
}
