package db

import (
	"fmt"
	gomysql "github.com/go-sql-driver/mysql"
	"github.com/siriusol/plantscale_db/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	myDBClient *gorm.DB
)

func Init() {
	var err error
	mysqlConf := pkg.GetDBConfig()

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cfg := gomysql.Config{
		User:      mysqlConf.Username,
		Passwd:    mysqlConf.Password,
		Addr:      fmt.Sprintf("%s:%d", mysqlConf.Host, mysqlConf.Port),
		Loc:       loc,
		ParseTime: true,
		DBName:    mysqlConf.DBName,
		Net:       "tcp",
	}
	myDBClient, err = gorm.Open(mysql.Open(cfg.FormatDSN()))
	if err != nil {
		panic(err)
	}
}

func Client() *gorm.DB {
	return myDBClient
}
