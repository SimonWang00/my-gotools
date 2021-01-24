package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/7

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

// StartPprof start http pprof
func StartPprof(addrs []string) {
	pprofServeMux := http.NewServeMux()
	pprofServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	for _, addr := range addrs {
		go func() {
			if err := http.ListenAndServe(addr, pprofServeMux); err != nil {
				fmt.Printf("http.ListenAndServe(\"%s\", pprofServeMux) error(%v)", addr, err)
				panic(err)
			}
		}()
	}
}

func main()  {
	StartPprof([]string{"127.0.0.1:8099"})
	select {
	}
}
