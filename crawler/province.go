package crawler

import (
	"CHN-Administrative-Divisions/service"
	"fmt"
)

// Province 爬取Province
func Province(fileName string) {
	f, err := service.PathExists(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !f {
		fmt.Println("不存在")
		//省份
		doc := CrawlProvince(BaseURL, Latest)
		if doc == nil {
			fmt.Println("获取省份失败")
			return
		}
		provinceList := DealProvince(doc)
		// 写入文件
		service.WriteToJsonFile(fileName, provinceList)
	}
}
