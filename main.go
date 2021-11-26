package main

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/crawler"
)

func main() {
	//爬取省份
	crawler.Province(base.ProvinceFile)
	//crawler.City(file.CityFile)
	//crawler.County(file.CountyFile)
	//crawler.Town(file.TownFile)
	//crawler.Village(file.VillageFile)

}
