package query

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

// ListCounty 查询所有县级行政区划信息
func ListCounty() (data []base.Division) {
	util.Read(base.CountyFile, &data)
	return
}

// GetCountyByCode 根据Code查询县级政区划信息
func GetCountyByCode(code string) base.Division {
	data := ListCounty()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

// ListCountyByName 根据名称模糊查询县级行政区划信息
func ListCountyByName(name string) []base.Division {
	var list []base.Division
	data := ListCounty()
	for _, datum := range data {
		if strings.Contains(datum.Name, name) {
			list = append(list, datum)
		}
	}
	return list
}
