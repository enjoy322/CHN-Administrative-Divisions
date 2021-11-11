package crawler

import (
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"golang.org/x/net/html"
	"strings"
)

func CrawlYear(url string) *html.Node {
	doc := crawlHtml(url)
	return doc
}

//省份
func CrawlProvince(url string, yearStr string) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlCity(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.Code[:2])
	builder.WriteString(".html")
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53.html
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlCounty(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
	builder.WriteString("/")
	builder.WriteString(division.ProvinceCode[:2])
	builder.WriteString("/")
	builder.WriteString(division.CityCode[:4])
	builder.WriteString(".html")
	//http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/53/5301.html
	doc := crawlHtml(builder.String())
	return doc
}

func CrawlTown(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
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

func CrawlVillage(url string, yearStr string, division model.Division) *html.Node {
	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString(yearStr)
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
