package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/service"
	"CHN-Administrative-Divisions/util"
	"fmt"
	"os"
	"time"
)

func Town(fileName string) {
	fmt.Println("-----------------town----------------")
	//	获取city
	upLevelList := service.ListCounty()
	f, err := util.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	var finalList []base.Division
	var failTimes = 0

	if !f {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 town")
		for _, division := range upLevelList {
			//单线程
			doc := CrawlTown(base.URL, division)
			time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 500)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
			tempList := DealTown(doc, division)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
	} else {
		fmt.Println("检查并爬取")
		//	需要爬取未爬取的
		// 读取文件
		util.Read(fileName, &finalList)
		fmt.Println("doneList:", len(finalList))
		needList := util.FindNeed(base.CodeCounty, upLevelList, finalList)
		fmt.Println("needCrawl:", len(needList))
		var newCrawl int
		for _, s := range needList {
			if !s.Branch {
				continue
			}
			doc := CrawlTown(base.URL, s)
			time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 500)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
			tempList := DealTown(doc, s)

			for _, m := range tempList {
				finalList = append(finalList, m)
				newCrawl++
			}
		}
		fmt.Println("newCrawl：", newCrawl)
	}
	// 写入文件
	fmt.Println("count:", len(finalList))
	util.WriteToJsonFile(fileName, finalList)
	fmt.Println("-----------------town----------------")
}
