package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/28

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包"
)

const (
	HOST 	 string = "127.0.0.1"
	PORT 	 string = "3306"
	DATABASE string = "blog"
	USER 	 string = "root"
	PASSWORD string = "000000"
)

// user表
type User struct {
	Id   int	`json:"id"`
	Name string	`json:"name"`
}


func main() {
	dbUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", USER, PASSWORD, HOST, PORT, DATABASE)
	db, _ := sql.Open("mysql", dbUrl)
	err := db.Ping()
	if err != nil{
		fmt.Println("mysql connect failed...")
	}
	fmt.Println("mysql connect sucess!")
	var user []User
	sqlString := "select * from  user where name=?"
	rows, err := db.Query(sqlString, "SimonWang00")
	var Id int
	var Name string
	for rows != nil && rows.Next() {
		if err := rows.Scan(&Id, &Name); err != nil {
			continue
		}
		user = append(user, User{Id: Id, Name: Name})
	}
	fmt.Println(user)
}