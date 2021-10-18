package main

import (
	"CHN-Administrative-Divisions/crawler"
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var dList []model.Division
	service.Read("./file/test.json", &dList)
	fmt.Println(len(dList), dList)

	//写入错误

	//var failTest = []model.Fail{
	//	{
	//		Type: model.FailCrawlCity,
	//		Divisions: []model.Division{
	//			{Url: "11.html", Code: "1100000000"},
	//			{Url: "12.html", Code: "1200000000"},
	//		},
	//	},
	//}
	//service.WriteToJsonFile("./file", "fail1.json", failTest)

	existsFail, err := service.PathExists("./file/fail.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsFail {
		//	不存在，不用爬取
		fmt.Println("不存在，不用爬取")
	} else {
		//	检查fail 中有误爬取city失败
		urlList := service.CheckFail(model.CodeFailCrawlCity)
		fmt.Println(urlList)
		if len(urlList) < 1 {
			fmt.Println("//无错误，不爬取")
			//	无错误，不怕取
		} else {
			fmt.Println("需要爬取未爬取的")
			//	需要爬取未爬取的
			var old []model.Division
			// 读取文件中的数据,保存为map格式
			data, _ := ioutil.ReadFile("./file/test.json")
			err := json.Unmarshal(data, &old)
			if err != nil {
				log.Fatal(err)
			}
			var newCityList []model.Division
			var failList []model.Fail
			var faild []model.Division
			for _, s := range urlList {
				doc := crawler.CrawlCity(crawler.BaseURL, "2020", s)
				if doc == nil {
					faild = append(faild, s)
					continue
				}
				tempList := crawler.DealCity(doc, s)
				for _, m := range tempList {
					newCityList = append(newCityList, m)
				}
			}
			if len(faild) > 0 {
				failList = append(failList, model.Fail{Type: model.CodeFailCrawlCity, Divisions: faild})
				//	写入json

			}

			fmt.Println("old：", len(old))
			fmt.Println("new：", len(newCityList))

			for _, division := range newCityList {
				old = append(old, division)
			}
			fmt.Println("count:", len(old))
			// 写入文件
			out, _ := json.Marshal(old)
			_ = ioutil.WriteFile("./file/index.json", out, 0755)
		}
	}

}
