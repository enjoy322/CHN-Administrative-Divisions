package query

import (
	"CHN-Administrative-Divisions/base"
)

// GetByCode 根据行政区划Code查询唯一行政区划信息
func GetByCode(code string) (data base.Division) {
	switch level(code) {
	case base.CodeProvince:
		data = GetProvinceByCode(code)
		break
	case base.CodeCity:
		data = GetCityByCode(code)
		break
	case base.CodeCounty:
		return GetCountyByCode(code)
	case base.CodeTown:
		return GetTownByCode(code)
	case base.CodeVillage:
		return GetVillageByCode(code)
	default:
		panic("err")
	}
	return
}

// ListBelongingsByCode 查询某行政区划的下一级所有行政区划信息
func ListBelongingsByCode(code string) (data []base.Division) {
	levelInt := level(code)
	switch levelInt {
	case base.CodeProvince:
		tempList := ListCity()
		for _, division := range tempList {
			if division.ProvinceCode == code {
				data = append(data, division)
			}
		}
		return
	case base.CodeCity:
		tempList := ListCounty()
		for _, division := range tempList {
			if division.CityCode == code {
				data = append(data, division)
			}
		}
		return
	case base.CodeCounty:
		tempList := ListTown()
		for _, division := range tempList {
			if division.CountyCode == code {
				data = append(data, division)
			}
		}
		return
	case base.CodeTown:
		tempList := ListVillage()
		for _, division := range tempList {
			if division.TownCode == code {
				data = append(data, division)
			}
		}
		return
	case base.CodeVillage:
		return
	default:
		panic("err")
	}
	return
}

// GetUpLevelDivisionByCode 查询某行政区划的上一级行政区划
func GetUpLevelDivisionByCode(code string) (data base.Division, f bool) {
	levelInt := level(code)
	switch levelInt {
	case base.CodeProvince:
		return base.Division{}, false
	case base.CodeCity:
		temp := GetCityByCode(code)
		data = GetProvinceByCode(temp.ProvinceCode)
		return
	case base.CodeCounty:
		temp := GetCountyByCode(code)
		data = GetCityByCode(temp.CityCode)
		return
	case base.CodeTown:
		temp := GetTownByCode(code)
		data = GetCountyByCode(temp.CountyCode)
		return
	case base.CodeVillage:
		temp := GetVillageByCode(code)
		data = GetTownByCode(temp.TownCode)
		return
	default:
		panic("err")
	}
	return
}

// ListByName 根据名称模糊查询
func ListByName(name string) map[int][]base.Division {
	data := make(map[int][]base.Division)
	dataProvince := ListProvinceByName(name)
	dataCity := ListCityByName(name)
	dataCounty := ListCountyByName(name)
	dataTown := ListTownByName(name)
	dataVillage := ListVillageByName(name)
	data[base.CodeProvince] = dataProvince
	data[base.CodeCity] = dataCity
	data[base.CodeCounty] = dataCounty
	data[base.CodeTown] = dataTown
	data[base.CodeVillage] = dataVillage
	return data
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
	switch 12 - count {
	case 2:
		return base.CodeProvince
	case 4:
		return base.CodeCity
	case 6:
		return base.CodeCounty
	case 9:
		return base.CodeTown
	case 12:
		return base.CodeVillage
	default:
		panic("err")
	}
}
