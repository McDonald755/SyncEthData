package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"os"
	"path"
)

var (
	APPVIPER *viper.Viper
	DB       *gorm.DB
	CLIENT   *ethclient.Client
)

func init() {
	APPVIPER = InitAppConfig()
	DB = InitDB()
	CLIENT = InitClient()
}

func InitAppConfig() *viper.Viper {
	workDir, _ := os.Getwd()
	appViper := viper.New()
	appViper.SetConfigName("application")
	appViper.SetConfigType("yml")
	appViper.AddConfigPath(path.Join(workDir, "config"))
	err := appViper.ReadInConfig()
	if err != nil {

	}
	return appViper
}

func InitDB() *gorm.DB {
	host := APPVIPER.GetString("database.host")
	port := APPVIPER.GetString("database.port")
	database := APPVIPER.GetString("database.databaseName")
	username := APPVIPER.GetString("database.username")
	password := APPVIPER.GetString("database.password")
	charset := APPVIPER.GetString("database.charset")
	loc := APPVIPER.GetString("database.loc")

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		panic("connected error" + err.Error())
	} else {
		fmt.Println("connected db")
	}
	return db
}

func InitClient() *ethclient.Client {
	url := APPVIPER.GetString("url")
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Error(err)
	} else {
		fmt.Println("client success")
	}
	return client

}
