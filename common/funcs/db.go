package funcs

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Db       string
	Username string
	Password string
	Params   string
}

func MustGetDb(logger logger.Interface) *gorm.DB {
	var config Config
	viper.SetConfigName(".env")
	viper.SetConfigType("toml")
	viper.AddConfigPath(ProjectPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic("read db file err: " + err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic("Unmarshal db conf err: " + err.Error())
	}
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Db,
		config.MySQL.Params,
	)
	db, err := gorm.Open(mysql.Open(connectStr), &gorm.Config{Logger: logger})
	if err != nil {
		panic("connect db err: " + err.Error())
	}
	return db
}
