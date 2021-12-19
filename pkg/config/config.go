package config

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
)

var RELEASE bool = false

func Init() {
	RELEASE = true
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// log.Errorf("read config failed: %v", err)
		// viper.Set("app_name", "File Master Client")
		// viper.Set("log.level", "DEBUG")
		// viper.Set("mysql.ip", "127.0.0.1")
		// viper.Set("mysql.port", 3306)
		// viper.Set("mysql.user", "root")
		// viper.Set("mysql.password", "123456")
		// viper.Set("mysql.database", "fmc")

		if err != viper.SafeWriteConfig() {
			log.Error(err)
		}
	}

	// fmt.Println()
	fmt.Println(viper.Get("app_name"))
	fmt.Println("server.host: ", viper.Get("server.host"))
	fmt.Println("server.port: ", viper.Get("server.port"))
	fmt.Println("control path: ", viper.Get("control_path"))

}
