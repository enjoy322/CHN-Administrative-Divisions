package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/service"
	"fmt"
	"log"
)

// Province 爬取Province
func Province(fileName string) {
	f, err := service.PathExists(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	if !f {
		fmt.Println("不存在")
		//省份
		doc := CrawlProvince(base.URL)
		if doc == nil {
			fmt.Println("获取省份失败")
			return
		}
		provinceList := DealProvince(doc)
		// 写入文件
		service.WriteToJsonFile(fileName, provinceList)
	}
}
