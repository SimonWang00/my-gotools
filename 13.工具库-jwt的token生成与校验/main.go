package main

import (
	"log"
	"my-gotools/13.工具库-jwt的token生成与校验/controller"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	token := &controller.TokenController{}
	mux.HandleFunc("/create_token", token.CreateToken)
	mux.HandleFunc("/create_rsa_token", token.CreateTokenByRsa)
	mux.HandleFunc("/test_token", token.TestToken)
	mux.HandleFunc("/test_rsa_token", token.TestRsaToken)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("http server bind:8080")
	err := s.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
