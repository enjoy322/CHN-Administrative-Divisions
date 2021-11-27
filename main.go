package main

import (
	"CHN-Administrative-Divisions/service"
	"fmt"
)

func main() {
	//爬取省份
	//crawler.Province(base.ProvinceFile)
	//crawler.City(file.CityFile)
	//crawler.County(file.CountyFile)
	//crawler.Town(file.TownFile)
	//crawler.Village(file.VillageFile)
	fmt.Println(service.Service{}.GetByCode("530000000000"))
	fmt.Println(service.GetProvinceByCode("530000000000"))
	fmt.Println(service.Service{}.ListBelongingsByCode("530000000000"))
	fmt.Println(service.Service{}.ListBelongingsByCode("532600000000"))
	fmt.Println(service.Service{}.ListBelongingsByCode("532601001000"))
	fmt.Println(service.Service{}.ListBelongingsByCode("532601001001"))

}
