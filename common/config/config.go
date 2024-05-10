package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	CONFIG_ENV = "CONFIG_ENV"
	defaultEnv = "dev"
)

var AppConfig *Config

func GetEnv() string {
	env := os.Getenv(CONFIG_ENV)
	if env == "" {
		logrus.Warn(fmt.Sprintf("CONFIG_ENV is NOT set, default to [%s].\n", defaultEnv))
		env = defaultEnv
	}
	return env
}

func Init() {
	AppConfig = &Config{}
	env := GetEnv()
	rootDir := GetRootDir()

	configPath := filepath.Join(rootDir, "config", fmt.Sprintf("%s.yaml", env))
	logrus.Info(fmt.Sprintf("Init config file, filepath: %s", configPath))

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		logrus.Error(fmt.Sprintf("Init config file error, filepath: %s, err: %v", configPath, err))
		panic(err)
	}

	// 解析配置文件
	configContent := []byte(os.ExpandEnv(string(data)))
	if err := yaml.Unmarshal(configContent, AppConfig); err != nil {
		logrus.Error(fmt.Sprintf("Parse config file error, filepath: %s, err: %v", configPath, err))
		panic(err)
	}
}

func GetRootDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Error(fmt.Sprintf("Get Root Dir error, err: %v", err))
		panic(err)
	}
	return dir
}
