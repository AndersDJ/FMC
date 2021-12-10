/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"

	cmd "github.com/AndersDJ/FMC/cmd"
	"github.com/AndersDJ/FMC/pkg/logger"
	log "github.com/cihub/seelog"

	"github.com/spf13/viper"
)

func main() {

	logger.Init()

	defer log.Flush()

	logger.TestLog()
	// 设置配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// viper.SetConfigType("json")
	viper.AddConfigPath(".")

	// 配置文件初始化
	// viper.Set("app_name", "File Master Client")
	// viper.Set("log.level", "DEBUG")
	// viper.Set("mysql.ip", "127.0.0.1")
	// viper.Set("mysql.port", 3306)
	// viper.Set("mysql.user", "root")
	// viper.Set("mysql.password", "123456")
	// viper.Set("mysql.database", "fmc")

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("read config failed: %v", err)
		viper.Set("app_name", "File Master Client")
		viper.Set("log.level", "DEBUG")
		viper.Set("mysql.ip", "127.0.0.1")
		viper.Set("mysql.port", 3306)
		viper.Set("mysql.user", "root")
		viper.Set("mysql.password", "123456")
		viper.Set("mysql.database", "fmc")

		if err != viper.SafeWriteConfig() {
			log.Error(err)
		}
	}

	// fmt.Println()
	fmt.Println(viper.Get("app_name"))
	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))
	fmt.Println("mysql databasee: ", viper.Get("mysql.databasee"))

	// viper.Set("redis.ip", "127.0.0.1")
	// viper.Set("redis.port", 6381)

	cmd.Execute()
}
