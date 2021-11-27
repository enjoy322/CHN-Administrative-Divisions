package query

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

// ListProvince 查询所有省级行政区划信息
func ListProvince() (data []base.Division) {
	util.Read(base.ProvinceFile, &data)
	return
}

// GetProvinceByCode 根据Code查询省级行政区划信息
func GetProvinceByCode(code string) base.Division {
	data := ListProvince()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

// ListProvinceByName 根据名称模糊查询省级行政区划信息
func ListProvinceByName(name string) []base.Division {
	var list []base.Division
	data := ListProvince()
	for _, datum := range data {
		if strings.Contains(datum.Name, name) {
			list = append(list, datum)
		}
	}
	return list
}
