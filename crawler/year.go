package crawler

import (
	"CHN-Administrative-Divisions/service"
	"fmt"
)

var Latest = "2020"

//爬取年
func Year() {
	//获取年份
	yearHtml := CrawlYear(BaseURL)
	if yearHtml == nil {
		fmt.Println("获取年份失败")
		return
	}
	yearList := DealYear(yearHtml)
	//写入数据库
	service.WriteToJsonFile("./file", "year.json", yearList)

	fmt.Println("yearList:", len(yearList), yearList)
	//取第一个最新
	latestYear := yearList[0]
	_ = latestYear.YearStr[:4]
}
