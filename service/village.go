package service

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

func ListVillage() (data []base.Division) {
	util.Read(base.VillageFile, &data)
	return
}

func GetVillageByCode(code string) base.Division {
	data := ListVillage()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

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
