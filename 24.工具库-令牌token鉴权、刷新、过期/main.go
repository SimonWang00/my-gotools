package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/5

/*
使用client_id去申请授权码
使用client_id client_secret 申请回的授权码去申请令牌（包含令牌ID, 令牌Token，令牌Token过期时间，刷新令牌Token）
使用令牌token去访问需要鉴权的网址
http://127.0.0.1:8080/credentials
http://127.0.0.1:8080/token?grant_type=client_credentials&client_id=1217a6e7&client_secret=88e97031&scope=all
http://127.0.0.1:8080/protected?access_token=AB7ODWY5P-YBI3AKAYDJ5Q
*/


func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		f.ServeHTTP(w, r)
	})
}


func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: "0.0.0.0:6379",
		DB:   0,
	}))
	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)		// 定时刷新
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(r *errors.Response) {
		log.Println("Response Error:", r.Error.Error())
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	//http://127.0.0.1:8080/credentials
	// {"CLIENT_ID":"4a9c083c","CLIENT_SECRET":"6cae7395"}
	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
		clientId := uuid.New().String()[:8]
		clientSecret := uuid.New().String()[:8]
		err := clientStore.Set(clientId, &models.Client{
			ID:     clientId,
			Secret: clientSecret,
			Domain: "http://localhost:8080",
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
	})

	//http://127.0.0.1:8080/protected
	http.HandleFunc("/protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}, srv))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
