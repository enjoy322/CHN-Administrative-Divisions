package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"strings"
)

//返回统计局提供的年份信息
func DealYear(doc *html.Node) []DivisionYear {
	var l []DivisionYear
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
		l = append(l, DivisionYear{
			YearStr:    year[:4] + "-01-01",
			Year:       TimeStrToTime(year[:4] + "-01-01").Unix(),
			UpdatedStr: updatedAt,
			UpdatedAt:  TimeStrToTime(updatedAt).Unix(),
		})
	}
	return l
}

//版本信息 写入

func Version() {
	//version := map[string]interface{}{
	//	"URL":         baseURL,
	//	"CreateAt":    BeijingTime().Unix(),
	//	"CreateAtStr": BeijingTime().Format("2006-01-02T15-04-05+08:00"),
	//	"Year":        TimeStrToTime().Unix(),
	//	"YearStr":     latestYear.Year,
	//	"Version":     latestYear.UpdatedAt,
	//	"VersionStr":  StampToTime(latestYear.UpdatedAt).Format("2006-01-02T15-04-05+08:00"),
	//}
	//WriteToJsonFile(dir, "version.json", version)
}

////省份
func DealProvince(doc *html.Node) []Division {
	var dList []Division
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
			simpleCode := url[:len(url)-5]
			var code string
			if len(simpleCode) < 12 {
				b := strings.Builder{}
				b.WriteString(simpleCode)
				for i := 0; i < 12-len(simpleCode); i++ {
					b.WriteString("0")
				}
				code = b.String()
			}
			dList = append(dList, Division{
				Url:        url,
				SimpleCode: simpleCode,
				Code:       code,
				Level:      1,
				Name:       provinceInfo.FirstChild.Data,
				FullName:   provinceInfo.FirstChild.Data,
			})
		}
	}
	return dList
}

//地级市
func DealCity(ch chan Division, doc *html.Node, division Division) {
	var cityList []Division

	matcherCity := matchByClass("class", "citytr")
	nodeCity := TraverseNode(doc, matcherCity)
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

////县级市
//func DealCounty(division Division, year string) []Division {
//	var divisionList []Division
//	url := baseURL + year + "/" + division.Url
//	content := DoRequest(url)
//	//匹配
//	docCity, errCity := html.Parse(strings.NewReader(string(content)))
//	if errCity != nil {
//		fmt.Println(errCity)
//		return nil
//	}
//	matcher := matchByClass("class", "countytr")
//	nodeCity := TraverseNode(docCity, matcher)
//	for _, node := range nodeCity {
//		if node.FirstChild.FirstChild.Data != "a" {
//			continue
//		}
//		tempUrl := node.FirstChild.FirstChild.Attr[0].Val
//		code := node.FirstChild.FirstChild.FirstChild.Data
//		name := node.LastChild.FirstChild.FirstChild.Data
//		divisionList = append(divisionList, Division{
//			Url:          code[:2] + "/" + tempUrl,
//			Code:         code,
//			Name:         name,
//			Level:        3,
//			CityCode:     division.Code,
//			ProvinceCode: division.ProvinceCode,
//		})
//	}
//	return divisionList
//}
//
////镇、街道
//func DealTown(division Division, year string) []Division {
//	var divisionList []Division
//	url := baseURL + year + "/" + division.Url
//	content := DoRequest(url)
//	//fmt.Println(string(content))
//	//匹配
//	docCity, errCity := html.Parse(strings.NewReader(string(content)))
//	if errCity != nil {
//		fmt.Println(errCity)
//		return nil
//	}
//	matcher := matchByClass("class", "towntr")
//	nodeCity := TraverseNode(docCity, matcher)
//	for _, node := range nodeCity {
//		tempUrl := node.FirstChild.FirstChild.Attr[0].Val
//		code := node.FirstChild.FirstChild.FirstChild.Data
//		name := node.LastChild.FirstChild.FirstChild.Data
//		divisionList = append(divisionList, Division{
//			Url:          code[:2] + "/" + code[2:4] + "/" + tempUrl,
//			Code:         code,
//			Name:         name,
//			Level:        4,
//			CountyCode:   division.Code,
//			CityCode:     division.CityCode,
//			ProvinceCode: division.ProvinceCode,
//		})
//	}
//	return divisionList
//}
//
////村级
//func DealVillage(division Division) []Division {
//	var divisionList []Division
//	url := baseURL + division.Url
//	content := DoRequest(url)
//	//匹配
//	docCity, errCity := html.Parse(strings.NewReader(string(content)))
//	if errCity != nil {
//		fmt.Println(errCity)
//		return nil
//	}
//	matcher := matchByClass("class", "villagetr")
//	nodeCity := TraverseNode(docCity, matcher)
//	for _, node := range nodeCity {
//		code := node.FirstChild.FirstChild.Data
//		name := node.LastChild.FirstChild.Data
//		divisionList = append(divisionList, Division{
//			Url:          "",
//			Code:         code,
//			Name:         name,
//			Level:        5,
//			TownCode:     division.Code,
//			CountyCode:   division.CountyCode,
//			CityCode:     division.CityCode,
//			ProvinceCode: division.ProvinceCode,
//		})
//	}
//	return divisionList
//}
//
////
// TraverseNode 收集与给定功能匹配的节点
func TraverseNode(doc *html.Node, matcher func(node *html.Node) (bool, bool)) (nodes []*html.Node) {
	var keep, exit bool
	var f func(*html.Node)
	f = func(n *html.Node) {
		keep, exit = matcher(n)
		if keep {
			nodes = append(nodes, n)
		}
		if exit {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
