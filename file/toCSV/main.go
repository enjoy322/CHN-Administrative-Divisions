package main

import (
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	s1 := time.Now().UnixNano() / 1e6
	var villageList []model.Division
	// 读取文件中的数据
	service.Read("./file/village.json", &villageList)
	fmt.Println(len(villageList))
	file, err := os.OpenFile("./file/v1.csv", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	//file.WriteString("tom,163,appName")
	w := csv.NewWriter(file)
	w.Write([]string{"Url", "Code", "SimpleCode", "Name", "FullName", "Level", "VillageType", "TownCode", "CountyCode", "CityCode", "ProvinceCode"})

	con := make([][]string, 0)
	for _, division := range villageList {
		con = append(con, []string{division.Url, division.Code, division.SimpleCode, division.Name, division.FullName, division.VillageType,
			division.VillageType, division.TownCode, division.CountyCode, division.CityCode, division.ProvinceCode})
	}
	w.WriteAll(con)
	w.Flush()
	fmt.Println(time.Now().UnixNano()/1e6-s1, "ms")

}
