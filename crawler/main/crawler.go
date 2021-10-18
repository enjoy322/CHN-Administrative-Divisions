package main

import (
	"CHN-Administrative-Divisions/crawler"
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"fmt"
	"time"
)

func main() {
	//年份和31个省份 爬取失败 直接结束
	failMap := make(map[string][]string)

	start := time.Now().UnixNano() / 1e6
	//获取年份
	yearHtml := crawler.CrawlYear(crawler.BaseURL)
	if yearHtml == nil {
		fmt.Println("获取年份失败")
		return
	}
	yearList := crawler.DealYear(yearHtml)
	//写入数据库
	service.WriteToJsonFile("./file", "year.json", yearList)

	fmt.Println("yearList:", len(yearList), yearList)
	//取第一个最新
	latestYear := yearList[0]
	latest := latestYear.YearStr[:4]

	//检查是否需要读取省份
	//检查是否存在province.json,不存在就爬取；存在：查看fail.json是否存在和fail.json中是否存在省份失败
	existsPro, err := service.PathExists("./file/province.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsPro {
		fmt.Println("不存在")
		//省份
		provinceHtml := crawler.CrawlProvince(crawler.BaseURL, latest)
		provinceList := crawler.DealProvince(provinceHtml)
		if len(provinceList) < 31 {
			fmt.Println("获取省份失败")
			return
		}
		if _, ok := failMap["Province"]; !ok {
			service.WriteToJsonFile("./file", "province.json", provinceList)
		}
	}

	//获取省份
	provinceList := service.ListProvince()
	fmt.Println(provinceList)

	existsCity, err := service.PathExists("./file/city.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsCity {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 city")
		var cityList []model.Division
		for _, division := range provinceList {
			//单线程
			doc := crawler.CrawlCity(crawler.BaseURL, latest, division)
			if doc == nil {
				failMap["CityOfProvince"] = append(failMap["CityOfProvince"], division.Url)
				continue
			}
			tempList := crawler.DealCity(doc, division)
			for _, m := range tempList {
				cityList = append(cityList, m)
			}
		}

		if _, ok := failMap["CityOfProvince"]; !ok {
			//	无失败
			service.WriteToJsonFile("./file", "city.json", cityList)
			fmt.Println("city.json write done")
		}
	} else {
		//检查fail.json
		fmt.Println("//存在city 检查fail")
		existsFail, err := service.PathExists("./file/fail1.json")
		if err != nil {
			fmt.Println(err.Error())
		}
		if !existsFail {
			//	不存在，不用爬取
		} else {
			//	检查fail 中有误爬取city失败
			urlList := service.CheckFail(model.FailCrawlCity)
			fmt.Println(urlList)
			if len(urlList) < 0 {
				//	无错误，不怕取
			} else {
				//	需要爬取未爬取的

			}
		}
	}

	cityList := service.ListCity()
	fmt.Println(len(cityList))

	//县级市
	//var countyList []Division
	//chanCounty := make(chan Division, 1000)
	//var wgCounty sync.WaitGroup
	//var cCount int64=0
	//for _, division := range cityList {
	//	//单线程
	//	c1 := time.Now().UnixNano() / 1e6
	//	doc := CrawlCounty(baseURL, latest, division)
	//	c2 := time.Now().UnixNano() / 1e6
	//	cCount = (c2 - c1) + cCount
	//	if doc == nil {
	//		fmt.Println("失败一次")
	//		failMap["County"] = append(failMap["County"], division.Url)
	//		continue
	//	}
	//	//	协程处理
	//	go func(ch chan Division, doc *html.Node, division2 Division, w *sync.WaitGroup) {
	//		w.Add(1)
	//		defer w.Done()
	//		DealCounty(ch, doc, division2)
	//	}(chanCounty, doc, division, &wgCounty)
	//}
	//
	//fmt.Println("cCounty cost:",cCount)
	////01
	//go func() {
	//	wgCounty.Wait()
	//	close(chanCounty)
	//}()
	//for division := range chanCounty {
	//	countyList = append(countyList, division)
	//}
	//
	//fmt.Println("county:", len(countyList))
	//2990--3271
	//45490ms

	//县级市
	var countyList []model.Division
	var ccc chan model.Division
	for _, city := range cityList {
		doc := crawler.CrawlCounty(crawler.BaseURL, latest, city)
		if doc == nil {
			fmt.Println("失败一次")
			failMap["County"] = append(failMap["County"], city.Url)
		}
		tList := crawler.DealCounty(ccc, doc, city)
		for _, division := range tList {
			countyList = append(countyList, division)
		}
	}

	fmt.Println("county:", len(countyList))
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
	fmt.Println("fail:", failMap)
	fmt.Println("crawl done,cost:", endTime-start, "ms")

	//写入
	//WriteToJsonFile(dir, "version.json", version)
	//WriteToJsonFile(dir, "province.json", provinceList)

}

func toFailFile() {
	service.WriteToJsonFile("./file", "city.json", nil)
}
