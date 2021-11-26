package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/service"
	"golang.org/x/net/html"
	"strings"
)

//省份
func CrawlProvince(url string) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString("/")
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlCity(url string, division base.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString("/")
	builder.WriteString(division.Code[:2])
	builder.WriteString(".html")
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53.html
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlCounty(url string, division base.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString("/")
	builder.WriteString(division.ProvinceCode[:2])
	builder.WriteString("/")
	builder.WriteString(division.CityCode[:4])
	builder.WriteString(".html")
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53/5301.html
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlTown(url string, division base.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString("/")
	builder.WriteString(division.ProvinceCode[:2])
	builder.WriteString("/")
	builder.WriteString(division.CityCode[2:4])
	builder.WriteString("/")
	builder.WriteString(division.CountyCode[:6])
	builder.WriteString(".html")
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/11/01/110101.html
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlVillage(url string, division base.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString("/")
	builder.WriteString(division.ProvinceCode[:2])
	builder.WriteString("/")
	builder.WriteString(division.CityCode[2:4])
	builder.WriteString("/")
	builder.WriteString(division.CountyCode[4:6])
	builder.WriteString("/")
	builder.WriteString(division.TownCode[:9])
	builder.WriteString(".html")
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/11/01/01/110101001.html
	doc := crawlHtml(builder.String())
	return doc
}

//爬取页面
func crawlHtml(url string) *html.Node {
	content, _, f := service.DoRequest(url)
	if !f {
		return nil
	}

	doc, err := html.Parse(content)
	if err != nil {
		panic("[error] html解析失败")
	}
	return doc
}
