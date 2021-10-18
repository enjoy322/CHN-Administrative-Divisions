package service

import "CHN-Administrative-Divisions/model"

//省份List

func ListProvince() []model.Division {
	var provinceList []model.Division

	Read("./file/province.json", &provinceList)
	return provinceList
}

func ListCity() []model.Division {
	var cityList []model.Division

	Read("./file/city.json", &cityList)
	return cityList
}
