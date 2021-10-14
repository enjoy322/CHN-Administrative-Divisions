package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"sync"
	"time"
)

const baseURL = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/"

func main() {
	start := time.Now().UnixNano() / 1e6
	//yearList := CrawlYear()
	//fmt.Println(yearList)
	//latestYear := yearList[0]
	//写入文件
	dir := "./file"
	err := MkDir(dir)
	if err != nil {
		panic("创建失败")
		return
	}
	//version := map[string]interface{}{
	//	"URL":         baseURL,
	//	"CreateAt":    BeijingTime().Unix(),
	//	"CreateAtStr": BeijingTime().Format("2006-01-02T15-04-05+08:00"),
	//	"Year":        TimeStrToTime(latestYear.Year + "-01-01").Unix(),
	//	"YearStr":     latestYear.Year,
	//	"Version":     latestYear.UpdatedAt,
	//	"VersionStr":  StampToTime(latestYear.UpdatedAt).Format("2006-01-02T15-04-05+08:00"),
	//}

	provinceList := CrawlProvince("2020")
	fmt.Println("province:", len(provinceList))
	fmt.Println(provinceList)

	//地级市
	var cityList []Division
	//for _, province := range provinceList {
	//	tempList := CrawlCity(province, latestYear.Year)
	//	for _, division := range tempList {
	//		cityList = append(cityList, division)
	//	}
	//}
	ch := make(chan Division, 100)
	var wg sync.WaitGroup

	for i, division := range provinceList {
		wg.Add(1)

		go func(ch chan Division, dd Division, i int) {
			defer wg.Done()
			CrawlCity(ch, dd, "2020")
		}(ch, division, i)
	}
	end := make(chan int)
	go func() {
		wg.Wait()
		end <- 1
	}()

L:
	for {
		select {
		case data := <-ch:

			cityList = append(cityList, data)
		case <-end:
			break L
		}

	}

	fmt.Println("city:", len(cityList))
	fmt.Println(cityList)
	//1s

	//县级市
	//var countyList []Division
	//for _, city := range cityList {
	//	tempList := CrawlCounty(city, latestYear.Year)
	//	for _, division := range tempList {
	//		countyList = append(countyList, division)
	//	}
	//}
	//fmt.Println("county:", len(countyList))
	//2990
	//5s

	////镇、街道
	//var townList []Division
	//for _, county := range countyList {
	//	tempList := CrawlTown(county, latestYear.Year)
	//	for _, division := range tempList {
	//		townList = append(townList, division)
	//	}
	//}
	//fmt.Println("town:", len(townList))
	//
	////村级
	//var villageList []Division
	//for _, town := range townList {
	//	tempList := CrawlTown(town, latestYear.Year)
	//	for _, division := range tempList {
	//		villageList = append(villageList, division)
	//	}
	//}
	//fmt.Println("village:", len(villageList))

	//for _, province := range provinceList {
	//	cityList := CrawlCity(province)
	//	for _, city := range cityList {
	//		countyList := CrawlCounty(city)
	//		for _, county := range countyList {
	//			townList := CrawlTown(county)
	//			for _, _ = range townList {
	//				//villageList := CrawlVillage(town)
	//				//fmt.Println(villageList)
	//			}
	//		}
	//	}
	//}

	endTime := time.Now().UnixNano() / 1e6
	fmt.Println("crawl done,cost:", endTime-start, "ms")

	//写入
	//WriteToJsonFile(dir, "version.json", version)
	//WriteToJsonFile(dir, "province.json", provinceList)

}

// TraverseNode 收集与给定功能匹配的节点
func TraverseNode(doc *html.Node, matcher func(node *html.Node) (bool, bool)) (nodes []*html.Node) {
	var keep, exit bool
	var f func(*html.Node)
	f = func(n *html.Node) {
		keep, exit = matcher(n)
		if keep {
			nodes = append(nodes, n)
		}
		if exit {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
