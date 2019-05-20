package mysql_conn

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	User string
	Pass string
	Host string
	Port string
	DB   string
}

func GetDBConn() *gorm.DB {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Mysql.User,
		config.Mysql.Pass,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.DB)

	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
