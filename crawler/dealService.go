package crawler

import (
	"CHN-Administrative-Divisions/base"
	"CHN-Administrative-Divisions/util"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"log"
	"strings"
)

// DealProvince 省份
func DealProvince(doc *html.Node) []base.Division {
	var dList []base.Division
	matcherCity := util.MatchByClass("class", "provincetr")
	nodes := TraverseNode(doc, matcherCity)
	for _, node := range nodes {
		//获取每个省份的超链接信息
		docPer, err2 := html.Parse(strings.NewReader(renderNode(node)))
		if err2 != nil {
			fmt.Println(err2)
			return nil
		}
		matcherPerProvince := util.MatcherByAtom(atom.A)
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
			dList = append(dList, base.Division{
				Branch:       true,
				Code:         code,
				Level:        base.CodeProvince,
				ProvinceCode: code,
				Name:         provinceInfo.FirstChild.Data,
			})
		}
	}
	return dList
}

// DealCity 地级市
func DealCity(doc *html.Node, division base.Division) []base.Division {
	var tempList []base.Division
	matcher := util.MatchByClass("class", "citytr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		//tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		var d = base.Division{
			Code:         code,
			Name:         name,
			Level:        base.CodeCity,
			Branch:       true,
			ProvinceCode: division.Code,
			CityCode:     code,
		}
		tempList = append(tempList, d)
	}
	return tempList
}

func DealCounty(doc *html.Node, division base.Division) []base.Division {
	var data []base.Division
	matcher := util.MatchByClass("class", "countytr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		if node.FirstChild.FirstChild.Data != "a" {
			code := node.FirstChild.FirstChild.Data
			name := node.LastChild.FirstChild.Data
			var d = base.Division{
				Code:         code,
				Name:         name,
				Branch:       false,
				Level:        base.CodeCounty,
				CountyCode:   code,
				CityCode:     division.Code,
				ProvinceCode: division.ProvinceCode,
			}
			data = append(data, d)

		} else {
			//tempUrl := node.FirstChild.FirstChild.Attr[0].Val
			code := node.FirstChild.FirstChild.FirstChild.Data
			name := node.LastChild.FirstChild.FirstChild.Data
			var d = base.Division{
				Code:         code,
				Name:         name,
				Level:        base.CodeCounty,
				Branch:       true,
				CountyCode:   code,
				CityCode:     division.Code,
				ProvinceCode: division.ProvinceCode,
			}
			data = append(data, d)
		}
	}
	return data
}

func DealTown(doc *html.Node, division base.Division) []base.Division {
	var data []base.Division
	matcher := util.MatchByClass("class", "towntr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		//tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		var d = base.Division{
			Code:         code,
			Name:         name,
			Level:        base.CodeTown,
			Branch:       true,
			TownCode:     code,
			CountyCode:   division.Code,
			CityCode:     division.CityCode,
			ProvinceCode: division.ProvinceCode,
		}
		data = append(data, d)
	}
	return data
}

func DealVillage(doc *html.Node, division base.Division) []base.Division {
	var data []base.Division
	matcher := util.MatchByClass("class", "villagetr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		code := node.FirstChild.FirstChild.Data
		//vType := node.FirstChild.NextSibling.FirstChild.Data
		name := node.LastChild.FirstChild.Data
		var d = base.Division{
			Code:         code,
			Name:         name,
			Level:        base.CodeVillage,
			TownCode:     division.Code,
			CountyCode:   division.CountyCode,
			CityCode:     division.CityCode,
			ProvinceCode: division.ProvinceCode,
		}
		data = append(data, d)
	}
	return data
}

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
	err := html.Render(w, n)
	if err != nil {
		log.Println(err)
		return ""
	}
	return buf.String()
}
