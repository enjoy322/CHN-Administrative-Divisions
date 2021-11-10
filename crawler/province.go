package crawler

import (
	"CHN-Administrative-Divisions/file"
	"CHN-Administrative-Divisions/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Province 爬取Province
func Province() {
	existsPro, err := service.PathExists(file.ProvinceFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !existsPro {
		fmt.Println("不存在")
		//省份
		doc := CrawlProvince(BaseURL, Latest)
		if doc == nil {
			fmt.Println("获取省份失败")
			return
		}
		provinceList := DealProvince(doc)
		fmt.Println(provinceList)
		// 写入文件
		out, _ := json.Marshal(provinceList)
		_ = ioutil.WriteFile(file.ProvinceFile, out, 0755)
	}
}
