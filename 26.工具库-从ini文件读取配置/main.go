package main

import (
	"github.com/go-ini/ini"
	"log"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/5

// MyCfg ini配置项
type MyCfg struct{
	AppModel string		`ini:"app_model"`
	Type	 int		`ini:"type"`
	MysqlCfg MysqlCfg	`ini:"mysql"`
}

// MysqlCfg mysql数据库配置
type MysqlCfg struct {
	Name string 		`ini:"username"`
	Pass string 		`ini:"passwd"`
}

func main() {
	fpath := "./26.工具库-从ini文件读取配置/default.ini"
	cfg, err := ini.Load(fpath)
	if err != nil{
		log.Println(err)
	}
	log.Println(cfg.Section("").Key("app_model").Value())
	log.Println(cfg.Section("mysql").Key("username").Value())
	log.Println(cfg.Section("mysql").Key("passwd").Value())
	p := new(MyCfg)
	err = cfg.MapTo(p)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(p)
}
