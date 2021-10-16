package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"time"
)

func CrawlYear(url string) *html.Node {
	doc := crawl(url)
	return doc
}

//省份
func CrawlProvince(url string, yearStr string) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	doc := crawl(builder.String())
	return doc
}

//省份
func CrawlCity(url string, yearStr string, division Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.Url)
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53.html
	doc := crawl(builder.String())
	return doc
}

//爬取页面
func crawl(url string) *html.Node {
	content, _, f := DoRequest(url)
	if !f {
		//	本次请求失败
		timer1 := time.NewTimer(time.Second * 3)
		<-timer1.C //阻塞，3秒以后重新执行
		fmt.Println("重新执行")
		CrawlYear(url)
	}

	doc, err := html.Parse(content)
	if err != nil {
		panic("[error] html解析失败")
	}
	return doc
}
