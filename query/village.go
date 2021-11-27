package query

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

// ListVillage 查询所有村级行政区划信息
func ListVillage() (data []base.Division) {
	util.Read(base.VillageFile, &data)
	return
}

// GetVillageByCode 根据Code查询村级行政区划信息
func GetVillageByCode(code string) base.Division {
	data := ListVillage()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

// ListVillageByName 根据名称模糊查询村级行政区划信息
func ListVillageByName(name string) []base.Division {
	var list []base.Division
	data := ListVillage()
	for _, datum := range data {
		if strings.Contains(datum.Name, name) {
			list = append(list, datum)
		}
	}
	return list
}
