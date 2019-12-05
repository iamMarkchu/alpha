package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Instance *viper.Viper
	err      error
)

func init() {
	Instance = viper.New()
	Instance.SetConfigName("config")
	Instance.AddConfigPath("./etc/")
	if err = Instance.ReadInConfig(); err != nil {
		fmt.Println("读取配置出现错误：", err.Error())
	}
}

func GetConfigInstance() *viper.Viper {
	return Instance
}
