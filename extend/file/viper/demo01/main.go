package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// 读取配置文件config
type Config struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

// viper:
//  1. 读取/修改、监听 本地或远程 json，toml，ini，yaml，hcl，env 等格式的文件内容
//  2. 配置key是不区分大小写的
func main() {
	// 把配置文件读取到结构体上
	var config Config

	// configName 指定了后缀的话需要单独设置 SetConfigType
	//  文件名相同时，会直接按顺序找，要想找其他必须写明后缀和type
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	viper.Unmarshal(&config) //将配置文件绑定到config上
	fmt.Printf("config: %+v, redis: %v \n", config, config.Redis)
	fmt.Printf("all data: %+v \n", viper.AllSettings())

	// 读取值
	portsSlice := viper.GetIntSlice("mysql.ports")
	fmt.Printf("portsSlice: %v \n", portsSlice)

	// 写入
	viper.Set("mysql.ports", []int{1201, 1611, 1235})
	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)
	}
}
