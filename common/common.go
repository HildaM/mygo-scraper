package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const (
	CONFIG_ENV = "CONFIG_ENV"
	defaultEnv = "dev"
)

func GetRootDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		// 此处不能使用logger包的Logger，因为尚未初始化
		logrus.Error(fmt.Sprintf("Get Root Dir error, err: %v", err))
		panic(err)
	}
	return dir
}

func GetEnv() string {
	env := os.Getenv(CONFIG_ENV)
	if env == "" {
		logrus.Warn(fmt.Sprintf("CONFIG_ENV is NOT set, default to [%s].\n", defaultEnv))
		env = defaultEnv
	}
	return env
}
