package service

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

func ListTown() (data []base.Division) {
	util.Read(base.TownFile, &data)
	return
}

func GetTownByCode(code string) base.Division {
	data := ListTown()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

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
