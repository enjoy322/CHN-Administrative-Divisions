package query

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

// ListCity 查询所有地市级行政区划信息
func ListCity() (data []base.Division) {
	util.Read(base.CityFile, &data)
	return
}

// GetCityByCode 根据Code查询地市级行政区划信息
func GetCityByCode(code string) base.Division {
	data := ListCity()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

// ListCityByName 根据名称模糊查询地市级行政区划信息
func ListCityByName(name string) []base.Division {
	var list []base.Division
	data := ListCity()
	for _, datum := range data {
		if strings.Contains(datum.Name, name) {
			list = append(list, datum)
		}
	}
	return list
}
