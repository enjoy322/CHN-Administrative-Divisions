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

// City 爬取City
func City() {
	fileName := file.CityFile
	fmt.Println("-----------------city----------------")
	//	获取所有省份
	provinceList := service.ListProvince()
	existsCity, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsCity {
		fmt.Println("//不存在 直接爬取 city")
		var cityList []model.Division
		var failTimes int
		var i = 0
		for _, division := range provinceList {
			//单线程
			doc := CrawlCity(BaseURL, Latest, division)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 10 {
					//错误10次结束本次爬取
					break
				}
				continue
			}
			tempList := DealCity(doc, division)
			for _, m := range tempList {
				cityList = append(cityList, m)
			}
			i++
			if i > 3 {
				break
			}
		}
		//city 写入
		city, _ := json.Marshal(cityList)
		_ = ioutil.WriteFile(file.CityFile, city, 0755)

	} else {
		fmt.Println("需要爬取未爬取的")
		//	需要爬取未爬取的
		var old []model.Division
		// 读取文件中的数据,保存为map格式
		data, _ := ioutil.ReadFile(fileName)
		err := json.Unmarshal(data, &old)
		if err != nil {
			return
		}
		var needCrawlList []model.Division
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
		for _, division := range provinceList {
			if _, ok := tempMap[division.Code]; !ok {
				needCrawlList = append(needCrawlList, division)
			}
		}

		var newList []model.Division
		var failTimes int
		for _, s := range needCrawlList {
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
				newList = append(newList, m)
			}
		}
		//重新写入
		if len(newList) > 0 {
			//	新内容
			fmt.Println("newList:", len(newList))
			for _, division := range newList {
				old = append(old, division)
			}
			fmt.Println("count:", len(old))
		}

		// 写入文件
		out, _ := json.Marshal(old)
		_ = ioutil.WriteFile("./file/city.json", out, 0755)
	}
	fmt.Println("-----------------city----------------")
}
