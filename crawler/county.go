package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/query"
	"CHN-Administrative-Divisions/util"
	"fmt"
	"os"
	"time"
)

func County(fileName string) {
	fmt.Println("-----------------county----------------")
	//	获取所有city
	upLevelList := query.ListCity()
	f, err := util.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	var finalList []base.Division
	failTimes := 0

	//todo
	//地级市直到镇，街道

	if !f {
		//	不存在 直接爬取
		fmt.Println("不存在 直接爬取")
		for _, division := range upLevelList {
			//单线程
			time.Sleep(time.Millisecond * 50)
			doc := CrawlCounty(base.URL, division)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := DealCounty(doc, division)
			fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
	} else {
		fmt.Println("爬取未爬取的")
		//	需要爬取未爬取的
		// 读取文件
		util.Read(fileName, &finalList)
		//old去重
		needList := util.FindNeed(base.CodeCity, upLevelList, finalList)
		fmt.Println("needCrawl:", len(needList))
		fmt.Println(needList)
		//	[{441900000000 东莞市   441900000000 440000000000 2 true} {442000000000 中山市   442000000000 440000000000 2 true} {460400000000 儋州市   460400000000 460000000000 2 true}]
		var newCrawl int
		for _, s := range needList {
			doc := CrawlCounty(base.URL, s)
			time.Sleep(time.Millisecond * 50)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 5 {
					break
				}
				continue
			}
			tempList := DealCounty(doc, s)
			fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
				newCrawl++
			}
		}
		fmt.Println("newCrawl:", newCrawl)
	}
	// 写入文件
	fmt.Println("count:", len(finalList))
	util.WriteToJsonFile(fileName, finalList)
	fmt.Println("-----------------county----------------")
}
