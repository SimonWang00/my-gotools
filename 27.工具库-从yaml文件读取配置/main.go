package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/5

import (
	easyconfig "github.com/spf13/viper"
	"log"
)

var (
	AppConfig   *appConfig
	MysqlConfig *mysqlConfig
)

// 应用程序配置
type appConfig struct {
	// 应用名称
	Name string
	// 运行模式: debug, release, gotest
	RunMode string
	// 运行 addr
	Addr string
	// 完整 url
	URL string
}

// Mysql配置
type mysqlConfig struct {
	Host 	 string
	Port 	 int
	Database string
	Username string
	Password string
}


func newAppConfig() *appConfig {
	// 默认配置
	easyconfig.SetDefault("APP.NAME", "gin_bbs")
	easyconfig.SetDefault("APP.RUNMODE", "release")
	easyconfig.SetDefault("APP.ADDR", ":8080")
	easyconfig.SetDefault("APP.KEY", "Atg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5")
	easyconfig.SetDefault("APP.ENABLE_CSRF", true)

	return &appConfig{
		Name:    easyconfig.GetString("APP.NAME"),
		RunMode: easyconfig.GetString("APP.RUNMODE"),
		Addr:    easyconfig.GetString("APP.ADDR"),
		URL:     easyconfig.GetString("APP.URL"),
	}
}

// newDBConfig 是数据库默认配置
func newDBConfig() *mysqlConfig {
	// 默认配置
	easyconfig.SetDefault("DB.HOST", "127.0.0.1")
	easyconfig.SetDefault("DB.PORT", 3306)
	easyconfig.SetDefault("DB.DATABASE", easyconfig.GetString("APP.NAME"))
	easyconfig.SetDefault("DB.USERNAME", "43.工具库-gin常用组件总结")
	easyconfig.SetDefault("DB.PASSWORD", "")

	username := easyconfig.GetString("DB.USERNAME")
	password := easyconfig.GetString("DB.PASSWORD")
	host := easyconfig.GetString("DB.HOST")
	port := easyconfig.GetInt("DB.PORT")
	database := easyconfig.GetString("DB.DATABASE")
	database = database + "_" + AppConfig.RunMode

	return &mysqlConfig{
		Host:       host,
		Port:       port,
		Database:   database,
		Username:   username,
		Password:   password,
	}
}


func main() {
	fpath := "./27.工具库-从yaml文件读取配置/config.yaml"
	configFileType := "yaml"
	// 初始化 viper 配置
	easyconfig.SetConfigFile(fpath)
	easyconfig.SetConfigType(configFileType)

	if err := easyconfig.ReadInConfig(); err != nil {
		log.Printf("读取配置文件失败，请检查 config.yaml 配置文件是否存在: %v \n", err)
		panic(err)
	}
	// 初始化 apps 配置
	AppConfig = newAppConfig()
	log.Println("config:", AppConfig)
	MysqlConfig = newDBConfig()
	log.Println("config:", MysqlConfig)

}
