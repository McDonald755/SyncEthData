package db

import (
	"SyncEthData/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

var DB *gorm.DB

func init() {
	DB = InitDB()
}

func InitDB() *gorm.DB {
	host := config.APPVIPER.GetString("database.host")
	port := config.APPVIPER.GetString("database.port")
	database := config.APPVIPER.GetString("database.databaseName")
	username := config.APPVIPER.GetString("database.username")
	password := config.APPVIPER.GetString("database.password")
	charset := config.APPVIPER.GetString("database.charset")
	loc := config.APPVIPER.GetString("database.loc")

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	//db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		panic("connected error" + err.Error())
	} else {
		fmt.Println("connected db")
	}
	return db
}
