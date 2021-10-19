package service

import "CHN-Administrative-Divisions/model"

//省份List

func ListProvince() []model.Division {
	var data []model.Division

	Read("./file/province.json", &data)
	return data
}

func ListCity() []model.Division {
	var data []model.Division

	Read("./file/city.json", &data)
	return data
}

func ListCounty() []model.Division {
	var data []model.Division

	Read("./file/county.json", &data)
	return data
}

func ListTown() []model.Division {
	var data []model.Division

	Read("./file/town.json", &data)
	return data
}
