package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

var APPVIPER *viper.Viper

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

func init() {
	APPVIPER = InitAppConfig()
}
