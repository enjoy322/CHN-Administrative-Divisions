package main

import (
	"CHN-Administrative-Divisions/query"
	"fmt"
)

func main() {
	//爬取省份
	//crawler.Province(base.ProvinceFile)
	//crawler.City(file.CityFile)
	//crawler.County(file.CountyFile)
	//crawler.Town(file.TownFile)
	//crawler.Village(file.VillageFile)
	fmt.Println(query.GetByCode("530000000000"))
	fmt.Println(query.GetProvinceByCode("530000000000"))
	fmt.Println(query.ListBelongingsByCode("530000000000"))
	fmt.Println(query.ListBelongingsByCode("532600000000"))
	fmt.Println(query.ListBelongingsByCode("532601001000"))
	fmt.Println(query.ListBelongingsByCode("532601001001"))
	fmt.Println(query.ListByName("昆明"))

}
