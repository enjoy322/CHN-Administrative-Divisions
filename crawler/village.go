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

func Village() {
	fmt.Println("-----------------village----------------")
	//	获取上一级
	fileName := file.VillageFile
	townList := service.ListTown()
	fmt.Println("town:", len(townList))

	existsDivision, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsDivision {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 town获取village")
		var finalList []model.Division
		var failTimes = 0
		for _, division := range townList {
			//单线程
			doc := CrawlVillage(service.BaseURL, Latest, division)
			time.Sleep(time.Millisecond * 20)
			if doc == nil {
				time.Sleep(time.Millisecond * 50)
				failTimes++
				if failTimes > 5 {
					break
				}
				continue
			}
			tempList := DealVillage(doc, division)
			fmt.Printf("---fail:---%d", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
		fmt.Println(finalList)

		//写入
		dataJson, _ := json.Marshal(finalList)
		fmt.Println("village:", len(finalList))
		_ = ioutil.WriteFile(fileName, dataJson, 0755)

	} else {
		fmt.Println("检查并爬取")
		//	需要爬取未爬取的
		s1 := time.Now().UnixNano() / 1e6
		var old []model.Division
		// 读取文件中的数据
		service.Read(fileName, &old)
		//待爬取
		//从保存的village中提取已经爬取得城镇
		var needCrawlTown []model.Division
		ss := time.Now().UnixNano() / 1e6
		fmt.Println("读取json:", ss-s1, "ms")

		//old去重
		var result []model.Division
		tempMap := map[string]model.Division{} // 存放不重复主键
		for _, e := range old {
			l := len(tempMap)
			tempMap[e.TownCode] = e
			if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
				result = append(result, e)
			}
		}
		s3 := time.Now().UnixNano() / 1e6
		fmt.Println("unique:", s3-ss, "ms")

		fmt.Println("result:", len(result))
		fmt.Println("old:", len(old))

		for _, division := range townList {
			if _, ok := tempMap[division.Code]; !ok {
				needCrawlTown = append(needCrawlTown, division)
			}
		}
		fmt.Println("need crawl deal cost:", time.Now().UnixNano()/1e6-ss, "ms")
		var newList []model.Division
		var failTimes = 0
		fmt.Println("needCrawlTown:", len(needCrawlTown))
		for _, s := range needCrawlTown {
			doc := CrawlVillage(service.BaseURL, Latest, s)
			time.Sleep(time.Millisecond * 20)
			if doc == nil {
				time.Sleep(time.Millisecond * 50)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := DealVillage(doc, s)
			fmt.Printf("---fail:---%d", failTimes)
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
	fmt.Println("-----------------village----------------")
}
