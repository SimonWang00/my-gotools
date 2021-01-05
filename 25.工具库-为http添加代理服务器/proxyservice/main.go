package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/5

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		parse, err := url.Parse("http://127.0.0.1:8081")
		if err != nil{
			panic(err)
		}
		log.Println("proxy server accept request!")
		proxy := httputil.NewSingleHostReverseProxy(parse)
		proxy.ServeHTTP(writer, request)
	})

	log.Println("proxy server success :", 8080)
	_ = http.ListenAndServe("127.0.0.1:8080", mux)
}
