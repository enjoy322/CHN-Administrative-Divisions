package service

import (
	"CHN-Administrative-Divisions/base"
	"fmt"
)

func GetByCode(code string) (data base.Division) {
	fmt.Println("code:", code)
	switch level(code) {
	case 2:
		data = GetProvinceByCode(code)
		break
	//case 4:
	//	data = GetCityByCode(code)
	//	break
	//case 6:
	//	return GetCountyByCode(code)
	//case 9:
	//	return GetTownByCode(code)
	//case 12:
	//	return GetVillageByCode(code)
	default:
		panic("err")
	}
	return
}

func level(code string) int {
	var count = 0
	for i := len(code); i > 0; i-- {
		if code[i-1:i] == "0" {
			count++
		} else {
			break
		}
	}
	fmt.Println(12 - count)
	return 12 - count
}
