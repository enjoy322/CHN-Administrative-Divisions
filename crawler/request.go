package crawler

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
)

const BaseURL = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/"

func DoRequest(url string) (io.Reader, error, bool) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("[error] 请求失败" + err.Error())
		return nil, err, false
	}
	request.Header.Set("allower_redirection", "False")
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[error] 请求失败" + err.Error())
		return nil, err, false
	}
	if response.StatusCode != 200 {
		//fmt.Println(response.Cookies())
		//uu := response.Header.Get("Location")
		//fmt.Println(uu)
		//body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
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
