package service

import (
	"CHN-Administrative-Divisions/base"
	"fmt"
)

//省份List

func ListProvince() []base.Division {
	var data []base.Division

	Read(base.ProvinceFile, &data)
	return data
}

func ListCity() []base.Division {
	var data []base.Division

	Read(base.CityFile, &data)
	return data
}

func ListCounty() []base.Division {
	var data []base.Division

	Read(base.CountyFile, &data)
	return data
}

func ListTown() []base.Division {
	var data []base.Division

	Read(base.TownFile, &data)
	return data
}

func ListVillage() []base.Division {
	var data []base.Division

	Read(base.VillageFile, &data)
	return data
}

func GetUp(code string, level int) {

	var d base.Division

	switch level {
	case base.CodeProvince:
		p := GetProvince(d.ProvinceCode)
		fmt.Println(p)

	}

}

// Get 根据code获取唯一的区划信息-针对json文件
func Get(code string) {
	//	对code初步处理，减少筛选
	var count = 0
	for i := len(code); i > 0; i-- {
		if code[i-1:i] == "0" {
			count++
		} else {
			break
		}
	}
	level := 12 - count
	switch level {

	case 2:
		//return GetProvince(code)
	}
}

func GetProvince(code string) *base.Division {
	data := ListProvince()

	for _, datum := range data {
		if datum.Code == code {
			return &datum
		}
	}
	return nil
}
