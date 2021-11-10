package service

import (
	"CHN-Administrative-Divisions/file"
	"CHN-Administrative-Divisions/model"
	"fmt"
)

//省份List

func ListProvince() []model.Division {
	var data []model.Division

	Read(file.ProvinceFile, &data)
	return data
}

func ListCity() []model.Division {
	var data []model.Division

	Read(file.CityFile, &data)
	return data
}

func ListCounty() []model.Division {
	var data []model.Division

	Read(file.CountyFile, &data)
	return data
}

func ListTown() []model.Division {
	var data []model.Division

	Read(file.TownFile, &data)
	return data
}

func ListVillage() []model.Division {
	var data []model.Division

	Read(file.VillageFile, &data)
	return data
}

func GetUp(code string, level int) {

	var d model.Division

	switch level {
	case model.CodeProvince:
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

func GetProvince(code string) *model.Division {
	data := ListProvince()

	for _, datum := range data {
		if datum.Code == code {
			return &datum
		}
	}
	return nil
}
