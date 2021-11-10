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

func County() {
	fmt.Println("-----------------county----------------")
	fileName := file.CountyFile
	//	获取所有city
	cityList := service.ListCity()
	existsDivision, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsDivision {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 county")
		var finalList []model.Division
		var failTimes = 0
		var done = 0
		for _, division := range cityList {
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
			done++
			fmt.Printf("---fail:---%d", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
		//	city 写入
		dataJson, _ := json.Marshal(finalList)
		_ = ioutil.WriteFile(fileName, dataJson, 0755)

	} else {
		fmt.Println("爬取未爬取的")
		//	需要爬取未爬取的
		var old []model.Division
		// 读取文件
		service.Read("./file/county.json", &old)
		//old去重
		var result []model.Division
		tempMap := map[string]model.Division{} // 存放不重复主键
		for _, e := range old {
			l := len(tempMap)
			tempMap[e.CityCode] = e
			if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
				result = append(result, e)
			}
		}
		var needCrawl []model.Division
		var doneCrawl = 0
		for _, division := range cityList {
			f := true
			for _, done := range result {
				if done.CityCode == division.Code {
					f = false
					doneCrawl++
				}

			}
			if f {
				needCrawl = append(needCrawl, division)
			}
		}

		var newList []model.Division
		var failTimes = 0
		fmt.Println("need:", needCrawl)
		var done = 0
		for _, s := range needCrawl {
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
			done++
			fmt.Printf("---fail:---%d", failTimes)
			for _, m := range tempList {
				newList = append(newList, m)
			}
		}
		//重新写入
		if len(newList) > 0 {
			//	新内容
			fmt.Println("old：", len(old))
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
	fmt.Println("-----------------county----------------")
}
