package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/service"
	"CHN-Administrative-Divisions/util"
	"fmt"
	"os"
	"time"
)

// City 爬取City
func City(fileName string) {
	fmt.Println("-----------------city----------------")
	//	获取省份
	upLevelList := service.ListProvince()
	f, err := util.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//错误10次结束本次爬取
	var finalList []base.Division

	if !f {
		fmt.Println("不存在 直接爬取")
		finalList = crawlCity(upLevelList)
	} else {
		fmt.Println("检查未爬取")
		// 读取文件中的数据,保存为map格式
		util.Read(fileName, &finalList)
		//需要爬取
		needList := util.FindNeed(base.CodeProvince, upLevelList, finalList)
		fmt.Println("needCrawl:", len(needList))
		newList := crawlCity(needList)
		fmt.Println("newCrawl:", len(newList))
		finalList = append(finalList, newList...)
	}
	// 写入文件
	fmt.Println("count:", len(finalList))
	util.WriteToJsonFile(fileName, finalList)
	fmt.Println("-----------------city----------------")
}

func crawlCity(crawlList []base.Division) (newList []base.Division) {
	var failTimes int
	for _, s := range crawlList {
		doc := CrawlCity(base.URL, s)
		if doc == nil {
			time.Sleep(time.Millisecond * 200)
			failTimes++
			if failTimes > 10 {
				break
			}
		}
		tempList := DealCity(doc, s)
		fmt.Fprintf(os.Stdout, "---fail:---%d\r", failTimes)
		for _, m := range tempList {
			newList = append(newList, m)
		}
	}
	return
}
