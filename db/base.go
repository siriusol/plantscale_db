package db

import (
	"fmt"
	"github.com/siriusol/plantscale_db/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	myDBClient *gorm.DB
)

func Init() {
	var err error
	mysqlConf := pkg.GetDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&tls=true", mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.DBName)
	dsn += "&loc=Asia%2fShanghai"
	myDBClient, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}

func Client() *gorm.DB {
	return myDBClient
}
