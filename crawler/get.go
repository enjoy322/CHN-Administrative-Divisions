package main

import (
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
	"strings"
)

func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic("[error] 获取网页内容错误！" + err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("[error] 解析错误！" + err.Error())
	}

	r2 := strings.NewReader(string(body))
	d2, err := charset.NewReader(r2, "gb2312")
	content, err := ioutil.ReadAll(d2)
	return content
}
