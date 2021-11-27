package query

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

// ListTown 查询所有乡镇级行政区划信息
func ListTown() (data []base.Division) {
	util.Read(base.TownFile, &data)
	return
}

// GetTownByCode 根据Code查询乡镇级行政区划信息
func GetTownByCode(code string) base.Division {
	data := ListTown()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

// ListTownByName 根据名称模糊查询乡镇级行政区划信息
func ListTownByName(name string) []base.Division {
	var list []base.Division
	data := ListTown()
	for _, datum := range data {
		if strings.Contains(datum.Name, name) {
			list = append(list, datum)
		}
	}
	return list
}
