package service

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"strings"
)

func ListProvince() (data []base.Division) {
	util.Read(base.ProvinceFile, &data)
	return
}

func GetProvinceByCode(code string) base.Division {
	data := ListProvince()
	for _, datum := range data {
		if datum.Code == code {
			return datum
		}
	}
	return base.Division{}
}

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
