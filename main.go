package main

import (
	"CHN-Administrative-Divisions/crawler"
	"CHN-Administrative-Divisions/file"
)

func main() {
	//爬取省份
	//crawler.Province(file.ProvinceFile)
	//crawler.City(file.CityFile)
	//crawler.County(file.CountyFile)
	crawler.Town(file.TownFile)
	//crawler.Village()

}
