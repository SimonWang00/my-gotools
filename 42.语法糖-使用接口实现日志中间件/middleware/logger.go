package middleware

//File  : logger.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/11

import (
	"fmt"
	"time"
)

type ServiceMiddleware func(Service) Service


type logServer struct {
	next Service
	log  string
}


func LogMiddleware(name string) ServiceMiddleware {
	return func(next Service) Service {
		return logServer{
			next: next,
			log:  "日志中间件V1",
		}
	}
}


func (mw logServer) Add(a, b int) (res int) {
	defer func(start time.Time) {
		fmt.Println("log", mw.log, "a - ", a, "\tb - ", b, "\tres - ", res, "\ttime ", time.Since(start))
	}(time.Now())
	return mw.next.Add(a, b)
}


type logV2Server struct {
	next Service
	log  string
}


func LogV2Middleware(name string) ServiceMiddleware {
	return func(next Service) Service {
		return logV2Server{
			next: next,
			log:  "日志中间件V2",
		}
	}
}


func (mw logV2Server) Add(a, b int) (res int) {
	defer func(start time.Time) {
		fmt.Println("log", mw.log, "a - ", a, "\tb - ", b, "\tres - ", res, "\ttime ", time.Since(start))
	}(time.Now())
	return mw.next.Add(a, b)
}
