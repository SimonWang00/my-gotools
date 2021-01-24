package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/28

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

// DB gorm
var (
	db *gorm.DB
	Url string = "127.0.0.1:3306"
	Connection string = "mysql"
	isDebug bool = true
)

// 文章表
type Article struct {
	Id 			int64	`db:"id"`
	Title 		string	`db:"title"`
	Summary 	string	`db:"summary"`
	Content 	string	`db:"content"`
	Classify 	string	`db:"classify"`
	Tag 		string	`db:"tag"`
	CreateTime 	string	`db:"create_time"`
}

// init 初始化数据库
func init() {
	var err error
	db, err = gorm.Open(Connection, Url)
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+ Url +" error: ", err)
	} else {
		log.Print("\n\n------------------------------------------ GORM OPEN SUCCESS! -----------------------------------------------\n\n")
	}
	db.LogMode(isDebug)
	// 连接池最大连接数100
	db.DB().SetMaxOpenConns(100)
	// 最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 创建所有表
	createTable()
}

// 创建表
func createTable()  {
	if !db.HasTable(&Article{}){
		err := db.Set("gorm:table_options","ENGINE=InnoDB DEFAULT CHARSET=utf8mb4" ).
			CreateTable(&Article{}).Error
		if err != nil{
			log.Fatalf("create table AriticleTable error(%v)!", err.Error())
		}
		log.Println("AriticleTable create suecess!")
	}
	log.Println("all tables are ready!")
}

// QueryLatestAriticle 主页拉取最新的五篇文章
func QueryLatestAriticle()  *[]Article {
	var ariticles []Article
	// SELECT title, summary, content, classify, tag, create_time FROM `ariticles`   ORDER BY create_time desc
	db.Select([]string{"id","title","summary","content","classify","tag","create_time"}).Order("create_time desc").Limit(5).Find(&ariticles)
	return &ariticles
}

// QueryAllAriticle 查询所有文章
func QueryAllAriticle(page int, pagesize int) (*[]Article, int){
	var ariticles []Article
	var totalblogs int
	db.Model(Article{}).Select([]string{"id","title","summary","content","classify","tag","create_time"}).
		Order("create_time desc").Limit(pagesize).Count(&totalblogs).
		Offset((page-1)*pagesize).Find(&ariticles)
	return &ariticles, totalblogs
}

// QueryAriticleByclassify 通过分类查询所有文章
func QueryAriticleByclassify(classify string, page int, pagesize int) (*[]Article, int) {
	var ariticles []Article
	var totalblogs int
	db.Model(Article{}).Where("classify = ?", classify).Count(&totalblogs).
		Select([]string{"id","title","summary","content","classify","tag","create_time"}).
		Order("create_time desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&ariticles)
	return &ariticles, totalblogs
}

// QueryAriticleById 通过id 查询文章
func QueryAriticleById(id int) *Article {
	var ariticle Article
	db.Where("id = ?", id).Select([]string{"id","title","summary","content","classify","tag","create_time"}).Find(&ariticle)
	return &ariticle
}

