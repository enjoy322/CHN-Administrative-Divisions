package crawler

import (
	"CHN-Administrative-Divisions/file"
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func Town() {
	fmt.Println("-----------------town----------------")
	//	获取所有city
	fileName := file.TownFile
	countyList := service.ListCounty()
	fmt.Println("county:", len(countyList))

	existsDivision, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsDivision {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 town")
		var finalList []model.Division
		var failTimes = 0
		for _, division := range countyList {
			//单线程
			doc := CrawlTown(BaseURL, Latest, division)
			time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 500)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			fmt.Printf("---fail:---%d", failTimes)
			tempList := DealTown(doc, division)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}

		//写入
		dataJson, _ := json.Marshal(finalList)
		fmt.Println("town:", len(finalList))
		_ = ioutil.WriteFile(fileName, dataJson, 0755)

	} else {
		fmt.Println("检查并爬取")
		//	需要爬取未爬取的
		var old []model.Division
		// 读取文件
		service.Read(fileName, &old)

		//old去重
		var result []model.Division
		tempMap := map[string]model.Division{} // 存放不重复主键
		for _, e := range old {
			l := len(tempMap)
			tempMap[e.CountyCode] = e
			if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
				result = append(result, e)
			}
		}

		var needCrawl []model.Division
		ss := time.Now().UnixNano() / 1e6
		var doneCrawl = 0
		for _, division := range countyList {
			f := true
			for _, done := range result {
				if done.CountyCode == division.Code {
					f = false
					doneCrawl++
				}

			}
			if f {
				needCrawl = append(needCrawl, division)
			}
		}
		fmt.Println("need crawl deal cost:", time.Now().UnixNano()/1e6-ss, "ms")

		var newList []model.Division
		var failTimes = 0
		for _, s := range needCrawl {
			doc := CrawlTown(BaseURL, Latest, s)
			time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 500)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			fmt.Printf("---fail:---%d", failTimes)
			tempList := DealTown(doc, s)

			for _, m := range tempList {
				newList = append(newList, m)
			}
			break

		}

		fmt.Println("old：", len(old))
		//重新写入
		if len(newList) > 0 {
			//	新内容
			fmt.Println("new：", len(newList))
			for _, division := range newList {
				old = append(old, division)
			}
			fmt.Println("total:", len(old))
		}

		// 写入文件
		out, _ := json.Marshal(old)
		_ = ioutil.WriteFile(fileName, out, 0755)
	}
	//}
	fmt.Println("-----------------town----------------")
}
