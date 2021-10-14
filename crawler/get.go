package main

import (
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func get(url string) []byte {
	resp := cGet(url)
	//if err != nil {
	//	panic("[error] 获取网页内容错误！" + err.Error())
	//}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp)

	if err != nil {
		panic("[error] 解析错误！" + err.Error())
	}

	r2 := strings.NewReader(string(body))
	d2, err := charset.NewReader(r2, "gb2312")
	content, err := ioutil.ReadAll(d2)
	return content
}

func cGet(url string) io.Reader {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, _ := client.Do(reqest)
	var b []byte
	_, err = response.Body.Read(b)
	//defer response.Body.Close()
	if err != nil {
		return nil
	}
	return response.Body
}
