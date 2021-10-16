package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
)

func DoRequest(url string) (io.Reader, error, bool) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("[error] 请求失败" + err.Error())
		return nil, err, false
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[error] 请求失败" + err.Error())
		return nil, err, false
	}
	if response.StatusCode != 200 {
		return nil, nil, false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("[error] body解析错误！" + err.Error())
		return nil, err, false
	}

	bodyStr := bytes.NewReader(body)
	toGBStr, err := charset.NewReader(bodyStr, "gb2312")
	if err != nil {
		fmt.Println("[error] 转换gb2312失败" + err.Error())
		return nil, err, false
	}
	return toGBStr, nil, true
}
