package main

import (
	"net/http"
	"testing"
)

func TestLimit(t *testing.T) {
	for i := 0; i < 2; i++ {
		go func() {
			for {
				_, err := http.Get("http://127.0.0.1:8099/limit_api")
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	select {}
}
