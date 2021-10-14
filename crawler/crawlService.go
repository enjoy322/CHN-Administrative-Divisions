package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

//返回统计局提供的年份信息
func CrawlYear() []DivisionYear {
	var yearList []DivisionYear
	reqUrl := baseURL
	content := get(reqUrl)
	fmt.Println(reqUrl)
	fmt.Println(string(content))
	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	matcherCity := matchByClass("class", "cont_tit")
	nodes := TraverseNode(doc, matcherCity)
	for _, node := range nodes {
		docContent, _ := html.Parse(strings.NewReader(renderNode(node)))
		matcherYear := matchByClass("class", "cont_tit03")
		matcherUpdatedAt := matchByClass("class", "cont_tit02")
		nodesYear := TraverseNode(docContent, matcherYear)
		nodesUpdatedAt := TraverseNode(docContent, matcherUpdatedAt)
		year := nodesYear[0].FirstChild.Data
		updatedAt := nodesUpdatedAt[0].FirstChild.Data
		TimeStrToTime("2013-11-06")

		yearList = append(yearList, DivisionYear{
			Year:      year[:4],
			UpdatedAt: TimeStrToTime(updatedAt).Unix(),
		})
	}
	return yearList
}

//省份

func CrawlProvince(year string) []Division {
	var divisionList []Division
	reqUrl := baseURL + year + "/"
	content := get(reqUrl)
	//匹配
	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	matcherCity := matchByClass("class", "provincetr")
	nodes := TraverseNode(doc, matcherCity)
	for _, node := range nodes {
		//获取每个省份的超链接信息
		docPer, err2 := html.Parse(strings.NewReader(renderNode(node)))
		if err2 != nil {
			fmt.Println(err2)
			return nil
		}
		matcherPerProvince := matcherByAtom(atom.A)
		nodesByA := TraverseNode(docPer, matcherPerProvince)
		for _, provinceInfo := range nodesByA {
			url := provinceInfo.Attr[0].Val
			code := url[:len(url)-5]
			if len(code) < 12 {
				b := strings.Builder{}
				b.WriteString(code)
				for i := 0; i < 12-len(code); i++ {
					b.WriteString("0")
				}
				code = b.String()
			}
			divisionList = append(divisionList, Division{
				Url:      url,
				Code:     code,
				Level:    1,
				Name:     provinceInfo.FirstChild.Data,
				FullName: provinceInfo.FirstChild.Data,
			})
		}
	}
	return divisionList
}

//地级市
func CrawlCity(ch chan Division, division Division, year string) {
	var cityList []Division
	url := baseURL + year + "/" + division.Url
	contentCity := get(url)
	//匹配
	docCity, errCity := html.Parse(strings.NewReader(string(contentCity)))
	if errCity != nil {
		fmt.Println(errCity)
		//return nil
	}
	matcherCity := matchByClass("class", "citytr")
	nodeCity := TraverseNode(docCity, matcherCity)
	for _, node := range nodeCity {
		tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		var d = Division{
			Url:          tempUrl,
			Code:         code,
			Name:         name,
			Level:        2,
			ProvinceCode: division.Code,
		}
		ch <- d
		cityList = append(cityList, d)
	}
	//fmt.Println(division.Name)
	//fmt.Println("crawl:",len(cityList))
	//fmt.Println( cityList)
}

//县级市
func CrawlCounty(division Division, year string) []Division {
	var divisionList []Division
	url := baseURL + year + "/" + division.Url
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
func CrawlTown(division Division, year string) []Division {
	var divisionList []Division
	url := baseURL + year + "/" + division.Url
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

//版本信息
