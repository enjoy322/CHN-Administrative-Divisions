package crawler

import (
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"fmt"
	"time"
)

// City 爬取City
func City(fileName string) {
	fmt.Println("-----------------city----------------")
	//	获取省份
	upLevelList := service.ListProvince()
	f, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//错误10次结束本次爬取
	var failTimes int
	var finalList []model.Division

	if !f {
		fmt.Println("不存在 直接爬取")
		for _, division := range upLevelList {
			doc := CrawlCity(BaseURL, Latest, division)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := DealCity(doc, division)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
	} else {
		fmt.Println("检查未爬取")
		// 读取文件中的数据,保存为map格式
		service.Read(fileName, &finalList)

		//需要爬取
		needList := service.FindNeed(model.CodeProvince, upLevelList, finalList)

		fmt.Println("needCrawl:", len(needList))
		var newCrawl int
		for _, s := range needList {
			doc := CrawlCity(BaseURL, Latest, s)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 10 {
					break
				}
			}
			tempList := DealCity(doc, s)
			for _, m := range tempList {
				finalList = append(finalList, m)
				newCrawl++
			}
		}
		//重新写入
		fmt.Println("newCrawl:", newCrawl)
	}
	// 写入文件
	fmt.Println("count:", len(finalList))
	service.WriteToJsonFile(fileName, finalList)
	fmt.Println("-----------------city----------------")
}
