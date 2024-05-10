package initializer

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Initializer struct {
	Name string
	Func func() error
}

var initList []Initializer

func init() {
	initList = make([]Initializer, 0)
}

// Init 对外暴露的初始化函数
func Init() {
	for _, v := range initList {
		if err := v.Func(); err != nil {
			logrus.Error(fmt.Sprintf("Initializer failed, name: %v, err: %v", v.Name, err))
			panic(err)
		}
		logrus.Info(fmt.Sprintf("Initializer success, name: %v", v.Name))
	}
}

// Register 注册初始化函数
func Register(name string, fn func() error) {
	initList = append(initList, Initializer{
		Name: name,
		Func: fn,
	})
}
