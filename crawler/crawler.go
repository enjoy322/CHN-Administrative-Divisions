package main

import (
	"fmt"
	"golang.org/x/net/html"
	"sync"
	"time"
)

const baseURL = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/"

func main() {
	start := time.Now().UnixNano() / 1e6
	//获取年份
	yearHtml := CrawlYear(baseURL)
	yearList := DealYear(yearHtml)
	fmt.Println("yearList:", len(yearList), yearList)
	//取第一个最新
	latestYear := yearList[0]
	latest := latestYear.YearStr[:4]
	//省份
	provinceHtml := CrawlProvince(baseURL, latest)
	provinceList := DealProvince(provinceHtml)
	fmt.Println("province:", len(provinceList), provinceList)

	//爬取地级市
	var cityList []Division
	ch := make(chan Division, 500)
	var wg sync.WaitGroup

	for _, division := range provinceList {
		//单线程
		doc := CrawlCity(baseURL, latest, division)

		//	多协程处理
		go func(ch chan Division, doc *html.Node, division2 Division, w *sync.WaitGroup) {
			w.Add(1)
			defer w.Done()
			DealCity(ch, doc, division2)
		}(ch, doc, division, &wg)
	}
	//03
	//var i = 1
	//	for{
	//		select {
	//		case s := <- ch:
	//			cityList = append(cityList, s)
	//			i++
	//		case <- time.After(time.Second*3):
	//			fmt.Printf("time out")
	//			goto out
	//		}
	//	}
	//out:
	//	fmt.Printf("end")
	//	fmt.Println(i)

	//01
	//var i = 1
	//go func() {
	//	wg.Wait()
	//	close(ch)
	//}()
	//
	//for division := range ch {
	//	fmt.Println(division)
	//	i++
	//}
	//fmt.Println(i)

	//end := make(chan int)
	//go func(w *sync.WaitGroup) {
	//	w.Wait()
	//	end <- 1
	//}(&wg)
	//
	//L:
	//	for {
	//		select {
	//		case data := <-ch:
	//
	//			cityList = append(cityList, data)
	//		case <-end:
	//			break L
	//		}
	//	}
	fmt.Println("cityList:", len(cityList), cityList)

	//	//地级市
	//	var cityList []Division
	//	//for _, province := range provinceList {
	//	//	tempList := CrawlCity(province, latestYear.Year)
	//	//	for _, division := range tempList {
	//	//		cityList = append(cityList, division)
	//	//	}
	//	//}
	//	ch := make(chan Division, 100)
	//	var wg sync.WaitGroup
	//
	//	for i, division := range provinceList {
	//		wg.Add(1)
	//
	//		go func(ch chan Division, dd Division, i int) {
	//			defer wg.Done()
	//			CrawlCity(ch, dd, "2020")
	//		}(ch, division, i)
	//	}
	//	end := make(chan int)
	//	go func() {
	//		wg.Wait()
	//		end <- 1
	//	}()
	//
	//L:
	//	for {
	//		select {
	//		case data := <-ch:
	//
	//			cityList = append(cityList, data)
	//		case <-end:
	//			break L
	//		}
	//
	//	}
	//
	//	fmt.Println("city:", len(cityList))
	//	fmt.Println(cityList)
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
