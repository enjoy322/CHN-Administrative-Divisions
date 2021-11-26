package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/service"
	"fmt"
	"os"
	"time"
)

func Village(fileName string) {
	fmt.Println("-----------------village----------------")
	//	获取上一级
	upLevelList := service.ListTown()
	f, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	var finalList []base.Division
	var failTimes = 0
	if !f {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 town获取village")

		for _, division := range upLevelList {
			//单线程
			doc := CrawlVillage(base.URL, division)
			time.Sleep(time.Millisecond * 20)
			if doc == nil {
				time.Sleep(time.Millisecond * 50)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := DealVillage(doc, division)
			fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
	} else {
		fmt.Println("检查并爬取")
		// 读取文件中的数据
		service.Read(fileName, &finalList)
		//从保存的village中提取已经爬取得城镇
		needList := service.FindNeed(base.CodeTown, upLevelList, finalList)

		fmt.Println("needCrawl:", len(needList))
		var newCrawl int
		for _, s := range needList {
			doc := CrawlVillage(base.URL, s)
			time.Sleep(time.Millisecond * 20)
			if doc == nil {
				time.Sleep(time.Millisecond * 50)
				failTimes++
				if failTimes > 50 {
					break
				}
				continue
			}
			tempList := DealVillage(doc, s)
			fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
				newCrawl++
			}
		}
		fmt.Println("newCrawl：", newCrawl)
	}
	// 写入文件
	fmt.Println("count:", len(finalList))
	service.WriteToJsonFile(fileName, finalList)
	fmt.Println("-----------------village----------------")
}
