package main

import (
	"fmt"
	"github.com/onyas/go-browsercookie"
	"io/ioutil"
	"net/http"
	"strings"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/25


func main() {
	var cookieJar, _ = browsercookie.Chrome("https://account.geekbang.org")
	var client = &http.Client{Jar: cookieJar}

	request, _ := http.NewRequest("POST", "https://time.geekbang.org/serv/v1/my/columns", strings.NewReader("{}"))
	// 设置请求头
	request.Header.Set("Origin", "https://time.geekbang.org")
	request.Header.Set("Referer", "https://time.geekbang.org")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	request.Header.Set("X-Real-IP", "211.161.244.200")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Content-Length", "2")

	res, _ := client.Do(request)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}