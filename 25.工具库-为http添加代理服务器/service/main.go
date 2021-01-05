package main

import (
	"log"
	"net/http"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/5

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/do", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("service is running")
		writer.Write([]byte("service done!"))
	})
	log.Println("server sucess:", 8080)
	_ = http.ListenAndServe("127.0.0.1:8081", mux)
}
