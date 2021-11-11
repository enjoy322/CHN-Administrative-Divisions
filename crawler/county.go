package crawler

import (
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"fmt"
	"time"
)

func County(fileName string) {
	fmt.Println("-----------------county----------------")
	//	获取所有city
	upLevelList := service.ListCity()
	f, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	var finalList []model.Division
	var failTimes int

	if !f {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 county")
		for _, division := range upLevelList {
			//单线程
			time.Sleep(time.Millisecond * 50)
			doc := CrawlCounty(BaseURL, Latest, division)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := DealCounty(doc, division)
			fmt.Printf("---fail:---%d", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
			break
		}
	} else {
		fmt.Println("爬取未爬取的")
		//	需要爬取未爬取的
		// 读取文件
		service.Read(fileName, &finalList)
		//old去重
		needList := service.FindNeed(model.CodeCity, upLevelList, finalList)
		var new int
		for _, s := range needList {
			doc := CrawlCounty(BaseURL, Latest, s)
			time.Sleep(time.Millisecond * 50)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := DealCounty(doc, s)
			fmt.Printf("---fail:---%d\n", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
				new++
			}
			break
		}
		fmt.Println("new:", new)
	}
	// 写入文件
	fmt.Println("count:", len(finalList))
	service.WriteToJsonFile(fileName, finalList)
	fmt.Println("-----------------county----------------")
}
