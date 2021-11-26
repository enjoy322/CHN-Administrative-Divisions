package service

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

func ListCounty() (data []base.Division) {
	util.Read(base.CountyFile, &data)
	return
}
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

func GetCountyByCode(code string) base.Division {
	data := ListCounty()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}
