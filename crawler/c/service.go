package c

import (
	"CHN-Administrative-Divisions/crawler"
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/process"
	"CHN-Administrative-Divisions/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

var Latest = "2020"

//爬取年

//爬取Province

//爬取City
func CCity() {
	fmt.Println("-----------------city----------------")
	//	获取所有省份

	provinceList := service.ListProvince()
	//fail
	//读取
	fails := service.CheckFail()

	existsCity, err := service.PathExists("./file/city.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var newFail model.Fail
	if !existsCity {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 city")
		var cityList []model.Division
		for _, division := range provinceList {
			//单线程
			doc := crawler.CrawlCity(crawler.BaseURL, Latest, division)
			if doc == nil {
				newFail.City = append(newFail.City, division)
				continue
			}
			tempList := crawler.DealCity(doc, division)
			for _, m := range tempList {
				cityList = append(cityList, m)
			}
		}

		//处理fail.json 文件
		if len(newFail.City) < 1 {
			//无失败
			fails.City = fails.City[0:0]
			//	写入json
			fs, _ := json.Marshal(fails)
			_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
		} else {
			//	有失败
			fails.City = newFail.City
			//	写入json
			fs, _ := json.Marshal(fails)
			_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
		}

		//	city 写入
		city, _ := json.Marshal(cityList)
		fmt.Println("city:", len(cityList))
		_ = ioutil.WriteFile("./file/city.json", city, 0755)

	} else {
		//检查fail.json
		fmt.Println("//存在city 检查fail")
		fails := service.CheckFail()
		fmt.Println("fails:", fails)
		cityFailList := fails.City
		fmt.Println("cityFailList:", cityFailList)
		if len(cityFailList) < 1 {
			fmt.Println("//无错误，不爬取")
			//	无错误，不怕取
		} else {
			fmt.Println("需要爬取未爬取的")
			//	需要爬取未爬取的
			var old []model.Division
			// 读取文件中的数据,保存为map格式
			data, _ := ioutil.ReadFile("./file/city.json")
			err := json.Unmarshal(data, &old)
			if err != nil {
				log.Fatal(err)
			}
			var newCityList []model.Division
			//var failList []model.Fail
			var faild []model.Division
			for _, s := range cityFailList {
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
			//处理fail.json 文件
			if len(faild) < 1 {
				//无失败
				fails.City = fails.City[0:0]
				//	写入json
				fs, _ := json.Marshal(fails)
				_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
			} else {
				//	有失败
				fails.City = faild
				//	写入json
				fs, _ := json.Marshal(fails)
				_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
			}

			fmt.Println("old：", len(old))
			fmt.Println("new：", len(newCityList))
			//重新写入
			if len(newCityList) > 0 {
				//	新内容
				fmt.Println("newCityList:", newCityList)
				for _, division := range newCityList {
					old = append(old, division)
				}
				fmt.Println("count:", len(old))
			}

			// 写入文件
			out, _ := json.Marshal(old)
			_ = ioutil.WriteFile("./file/city.json", out, 0755)
		}
	}
	fmt.Println("-----------------city----------------")
}

func CCounty() {
	fmt.Println("-----------------county----------------")
	//	获取所有city
	upList := service.ListCity()

	fails := service.CheckFail()

	existsDivision, err := service.PathExists("./file/county.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var newFail model.Fail
	var total float64 = 3271
	var hasTotal float64 = 0
	if !existsDivision {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 county")
		var finalList []model.Division
		for _, division := range upList {
			//单线程
			doc := crawler.CrawlCounty(crawler.BaseURL, Latest, division)
			if doc == nil {
				newFail.County = append(newFail.County, division)
				continue
			}
			tempList := crawler.DealCounty(doc, division)
			hasTotal += float64(len(tempList))
			percent := hasTotal / total * 100
			percentStr := strconv.FormatFloat(percent, 'E', 2, 64)
			percentJson, _ := json.Marshal(map[string]string{"County": percentStr + "%"})
			_ = ioutil.WriteFile("./file/crawlProcess.json", percentJson, 0755)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}

		}

		//处理fail.json 文件
		if len(newFail.County) < 1 {
			//无失败
			fails.County = fails.County[0:0]
			//	写入json
			fs, _ := json.Marshal(fails)
			_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
		} else {
			//	有失败
			fails.County = newFail.County
			//	写入json
			fs, _ := json.Marshal(fails)
			_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
		}

		//	city 写入
		dataJson, _ := json.Marshal(finalList)
		fmt.Println("county:", len(finalList))
		_ = ioutil.WriteFile("./file/county.json", dataJson, 0755)

	} else {
		//检查fail.json
		fmt.Println("//存在 检查fail")
		failList := fails.County
		fmt.Println("FailList:", failList)
		if len(failList) < 1 {
			fmt.Println("//无错误，不爬取")
			//	无错误，不怕取
		} else {
			fmt.Println("爬取未爬取的")
			//	需要爬取未爬取的
			var old []model.Division
			// 读取文件中的数据,保存为map格式
			data, _ := ioutil.ReadFile("./file/county.json")
			err := json.Unmarshal(data, &old)
			if err != nil {
				log.Fatal(err)
			}
			var newList []model.Division
			//var failList []model.Fail
			//var newFail []model.Division
			for _, s := range failList {
				doc := crawler.CrawlCounty(crawler.BaseURL, Latest, s)
				if doc == nil {
					newFail.County = append(newFail.County, s)
					continue
				}
				tempList := crawler.DealCounty(doc, s)
				for _, m := range tempList {
					newList = append(newList, m)
				}
			}
			//处理fail.json 文件
			if len(newFail.County) < 1 {
				//无失败
				fails.County = fails.County[0:0]
				//	写入json
				fs, _ := json.Marshal(fails)
				_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
			} else {
				//	有失败
				fails.County = newFail.County
				//	写入json
				fs, _ := json.Marshal(fails)
				_ = ioutil.WriteFile("./file/fail.json", fs, 0755)
			}

			fmt.Println("old：", len(old))
			fmt.Println("new：", len(newList))
			//重新写入
			if len(newList) > 0 {
				//	新内容
				fmt.Println("newList:", newList)
				for _, division := range newList {
					old = append(old, division)
				}
				fmt.Println("count:", len(old))
			}

			// 写入文件
			out, _ := json.Marshal(old)
			_ = ioutil.WriteFile("./file/county.json", out, 0755)
		}
	}
	fmt.Println("-----------------county----------------")
}

func CTown() {
	fmt.Println("-----------------town----------------")
	//	获取所有city
	upList := service.ListCounty()
	fmt.Println("county:", len(upList))

	var bar process.Bar
	existsDivision, err := service.PathExists("./file/town.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsDivision {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 town")
		var finalList []model.Division
		bar.NewOption(0, int64(len(upList)))
		var i = 0
		var failI = 0
		for _, division := range upList {
			if division.Url == "" {
				continue
			}
			//单线程
			doc := crawler.CrawlTown(crawler.BaseURL, Latest, division)
			time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 500)
				failI++
				if failI > 10 {
					break
				}
				continue
			}
			i++
			bar.Play(int64(i))
			tempList := crawler.DealTown(doc, division)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
		bar.Finish()

		//写入
		dataJson, _ := json.Marshal(finalList)
		fmt.Println("town:", len(finalList))
		_ = ioutil.WriteFile("./file/town.json", dataJson, 0755)

	} else {
		fmt.Println("检查并爬取")
		//	需要爬取未爬取的
		var old []model.Division
		// 读取文件
		service.Read("./file/town.json", &old)
		var needCrawl []model.Division
		ss := time.Now().UnixNano() / 1e6
		var doneCrawl = 0
		for _, division := range upList {
			f := true
			for _, done := range old {
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
		var failI = 0
		bar.NewOption(int64(len(old)), int64(len(upList)))
		var j = 0
		for _, s := range needCrawl {
			if s.Url == "" {
				continue
			}
			doc := crawler.CrawlTown(crawler.BaseURL, Latest, s)
			time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 500)
				failI++
				if failI > 10 {
					break
				}
				continue
			}
			j++
			bar.Play(int64(doneCrawl + j))
			tempList := crawler.DealTown(doc, s)

			for _, m := range tempList {
				newList = append(newList, m)
			}
		}
		bar.Finish()

		fmt.Println("old：", len(old))
		fmt.Println("new：", len(newList))
		//重新写入
		if len(newList) > 0 {
			//	新内容
			fmt.Println("newList:", newList)
			for _, division := range newList {
				old = append(old, division)
			}
			fmt.Println("count:", len(old))
		}

		// 写入文件
		out, _ := json.Marshal(old)
		_ = ioutil.WriteFile("./file/town.json", out, 0755)
	}
	//}
	fmt.Println("-----------------town----------------")
}

func CVillage() {
	fmt.Println("-----------------village----------------")
	//	获取上一级
	upList := service.ListTown()
	fmt.Println("town:", len(upList))

	existsDivision, err := service.PathExists("./file/village.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var bar process.Bar
	var i = 0
	if !existsDivision {
		//	不存在 直接爬取
		fmt.Println("//不存在 直接爬取 town获取village")
		var finalList []model.Division
		bar.NewOption(0, int64(len(upList)))
		var failTimes = 0
		for _, division := range upList {
			//单线程
			doc := crawler.CrawlVillage(crawler.BaseURL, Latest, division)
			//time.Sleep(time.Millisecond * 100)
			if doc == nil {
				time.Sleep(time.Millisecond * 300)
				failTimes++
				if failTimes > 10 {
					break
				}
				continue
			}
			tempList := crawler.DealVillage(doc, division)
			i++
			bar.Play(int64(i))
			fmt.Printf("---fail:---%d", failTimes)
			for _, m := range tempList {
				finalList = append(finalList, m)
			}
		}
		bar.Finish()

		//写入
		dataJson, _ := json.Marshal(finalList)
		fmt.Println("village:", len(finalList))
		_ = ioutil.WriteFile("./file/village.json", dataJson, 0755)

	} else {
		fmt.Println("检查并爬取")
		//	需要爬取未爬取的
		s1 := time.Now().UnixNano() / 1e6
		var old []model.Division
		// 读取文件中的数据
		service.Read("./file/village.json", &old)

		//待爬取
		//从保存的village中提取已经爬取得城镇
		var needCrawlTown []model.Division
		ss := time.Now().UnixNano() / 1e6
		fmt.Println(ss-s1, "ms")

		for _, division := range upList {
			f := true
			for _, done := range old {
				if done.TownCode == division.Code {
					f = false
				}

			}
			if f {
				needCrawlTown = append(needCrawlTown, division)
			}
		}
		all := len(upList)
		need := len(needCrawlTown)
		fmt.Println("need crawl deal cost:", time.Now().UnixNano()/1e6-ss, "ms")

		var newList []model.Division
		var failTimes = 0
		bar.NewOption(0, int64(len(upList)))
		var j = 0
		fmt.Println("needCrawlTown:", len(needCrawlTown))
		for _, s := range needCrawlTown {
			time.Sleep(time.Millisecond * 50)
			doc := crawler.CrawlVillage(crawler.BaseURL, Latest, s)
			if doc == nil {
				time.Sleep(time.Millisecond * 200)
				failTimes++
				if failTimes > 2 {
					break
				}
				continue
			}
			j++
			tempList := crawler.DealVillage(doc, s)
			bar.Play(int64(all - need + j))
			fmt.Printf("---fail:---%d", failTimes)
			for _, m := range tempList {
				newList = append(newList, m)
			}
		}
		bar.Finish()

		fmt.Println("old：", len(old))
		fmt.Println("new：", len(newList))
		//重新写入
		if len(newList) > 0 {
			//	新内容
			for _, division := range newList {
				old = append(old, division)
			}
			fmt.Println("count:", len(old))
		}

		// 写入文件
		out, _ := json.Marshal(old)
		_ = ioutil.WriteFile("./file/village.json", out, 0755)
	}
	fmt.Println("-----------------village----------------")
}
