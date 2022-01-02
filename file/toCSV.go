package file

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ToCSV() {
	fmt.Println("to csv")

	file, err := os.OpenFile("csv/division.csv", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"code", "province_code", "city_code", "county_code", "town_code", "name", "level", "branch"})
	w.Flush()

	for _, v := range base.File {
		var list []base.Division
		util.Read(v, &list)
		var data [][]string
		for i, division := range list {
			if i%100 == 0 {
				w.WriteAll(data)
				w.Flush()
				data = make([][]string, 0)
			}
			var f = 0
			if division.Branch {
				f = 1
			}
			data = append(data, []string{division.Code, division.ProvinceCode, division.CityCode,
				division.CountyCode, division.TownCode, division.Name,
				strconv.Itoa(int(division.Level)), strconv.Itoa(f)})
			if i == len(list)-1 {
				w.WriteAll(data)
				w.Flush()
			}
		}
	}
}
