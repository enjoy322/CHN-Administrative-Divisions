package service

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

func ListCity() (data []base.Division) {
	util.Read(base.CityFile, &data)
	return
}

func GetCityByCode(code string) base.Division {
	data := ListCity()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

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
