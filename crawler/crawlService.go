package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

//地级市
func CrawlCity(division Division) []Division {
	var cityList []Division
	url := baseURL + division.Url
	contentCity := get(url)
	//匹配
	docCity, errCity := html.Parse(strings.NewReader(string(contentCity)))
	if errCity != nil {
		fmt.Println(errCity)
		return nil
	}
	matcherCity := matchByClass("class", "citytr")
	nodeCity := TraverseNode(docCity, matcherCity)
	for _, node := range nodeCity {
		tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		cityList = append(cityList, Division{
			Url:          tempUrl,
			Code:         code,
			Name:         name,
			Level:        2,
			ProvinceCode: division.Code,
		})
	}
	return cityList
}

//县级市
func CrawlCounty(division Division) []Division {
	var divisionList []Division
	url := baseURL + division.Url
	content := get(url)
	//匹配
	docCity, errCity := html.Parse(strings.NewReader(string(content)))
	if errCity != nil {
		fmt.Println(errCity)
		return nil
	}
	matcher := matchByClass("class", "countytr")
	nodeCity := TraverseNode(docCity, matcher)
	for _, node := range nodeCity {
		if node.FirstChild.FirstChild.Data != "a" {
			continue
		}
		tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		divisionList = append(divisionList, Division{
			Url:          code[:2] + "/" + tempUrl,
			Code:         code,
			Name:         name,
			Level:        3,
			CityCode:     division.Code,
			ProvinceCode: division.ProvinceCode,
		})
	}
	return divisionList
}

//镇、街道
func CrawlTown(division Division) []Division {
	var divisionList []Division
	url := baseURL + division.Url
	content := get(url)
	//fmt.Println(string(content))
	//匹配
	docCity, errCity := html.Parse(strings.NewReader(string(content)))
	if errCity != nil {
		fmt.Println(errCity)
		return nil
	}
	matcher := matchByClass("class", "towntr")
	nodeCity := TraverseNode(docCity, matcher)
	for _, node := range nodeCity {
		tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		divisionList = append(divisionList, Division{
			Url:          code[:2] + "/" + code[2:4] + "/" + tempUrl,
			Code:         code,
			Name:         name,
			Level:        4,
			CountyCode:   division.Code,
			CityCode:     division.CityCode,
			ProvinceCode: division.ProvinceCode,
		})
	}
	return divisionList
}

//村级
func CrawlVillage(division Division) []Division {
	var divisionList []Division
	url := baseURL + division.Url
	content := get(url)
	//匹配
	docCity, errCity := html.Parse(strings.NewReader(string(content)))
	if errCity != nil {
		fmt.Println(errCity)
		return nil
	}
	matcher := matchByClass("class", "villagetr")
	nodeCity := TraverseNode(docCity, matcher)
	for _, node := range nodeCity {
		code := node.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.Data
		divisionList = append(divisionList, Division{
			Url:          "",
			Code:         code,
			Name:         name,
			Level:        5,
			TownCode:     division.Code,
			CountyCode:   division.CountyCode,
			CityCode:     division.CityCode,
			ProvinceCode: division.ProvinceCode,
		})
	}
	return divisionList
}
