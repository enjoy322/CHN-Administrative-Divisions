package crawler

import (
	"CHN-Administrative-Divisions/model"
	"golang.org/x/net/html"
	"strings"
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

func CrawlCity(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.Url)
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53.html
	doc := crawl(builder.String())
	return doc
}

func CrawlCounty(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.Url)
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53/5301.html
	doc := crawl(builder.String())
	return doc
}

func CrawlTown(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.Url)
	//fmt.Println(builder.String())
	doc := crawl(builder.String())
	return doc
}

func CrawlVillage(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.Url)
	//fmt.Println(builder.String())
	doc := crawl(builder.String())
	return doc
}

//爬取页面
func crawl(url string) *html.Node {
	content, _, f := DoRequest(url)
	if !f {
		return nil
		//	本次请求失败
		//timer1 := time.NewTimer(time.Second*5)
		//<-timer1.C
		//fmt.Println("重新执行",url)
		//CrawlYear(url)
	}

	doc, err := html.Parse(content)
	if err != nil {
		panic("[error] html解析失败")
	}
	return doc
}
