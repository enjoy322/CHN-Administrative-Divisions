package crawler

import (
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"strings"
)

//返回统计局提供的年份信息
func DealYear(doc *html.Node) []model.DivisionYear {
	var l []model.DivisionYear
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
		l = append(l, model.DivisionYear{
			YearStr:    year[:4] + "-01-01",
			Year:       service.TimeStrToTime(year[:4] + "-01-01").Unix(),
			UpdatedStr: updatedAt,
			UpdatedAt:  service.TimeStrToTime(updatedAt).Unix(),
		})
	}
	return l
}

//版本信息 写入

func WriteVersion() {
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
func DealProvince(doc *html.Node) []model.Division {
	var dList []model.Division
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
			dList = append(dList, model.Division{
				Branch:       true,
				Code:         code,
				Level:        model.CodeProvince,
				ProvinceCode: code,
				Name:         provinceInfo.FirstChild.Data,
			})
		}
	}
	return dList
}

//地级市
func DealCity(doc *html.Node, division model.Division) []model.Division {
	var tempList []model.Division
	matcher := matchByClass("class", "citytr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		//tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		var d = model.Division{
			Code:         code,
			Name:         name,
			Level:        model.CodeCity,
			Branch:       true,
			ProvinceCode: division.Code,
			CityCode:     code,
		}
		tempList = append(tempList, d)
	}
	return tempList
}

func DealCounty(doc *html.Node, division model.Division) []model.Division {
	var data []model.Division
	matcher := matchByClass("class", "countytr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		if node.FirstChild.FirstChild.Data != "a" {
			code := node.FirstChild.FirstChild.Data
			name := node.LastChild.FirstChild.Data
			var d = model.Division{
				Code:         code,
				Name:         name,
				Branch:       false,
				Level:        model.CodeCounty,
				CountyCode:   code,
				CityCode:     division.Code,
				ProvinceCode: division.ProvinceCode,
			}
			data = append(data, d)

		} else {
			//tempUrl := node.FirstChild.FirstChild.Attr[0].Val
			code := node.FirstChild.FirstChild.FirstChild.Data
			name := node.LastChild.FirstChild.FirstChild.Data
			var d = model.Division{
				Code:         code,
				Name:         name,
				Level:        model.CodeCounty,
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

func DealTown(doc *html.Node, division model.Division) []model.Division {
	var data []model.Division
	matcher := matchByClass("class", "towntr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		//tempUrl := node.FirstChild.FirstChild.Attr[0].Val
		code := node.FirstChild.FirstChild.FirstChild.Data
		name := node.LastChild.FirstChild.FirstChild.Data
		var d = model.Division{
			Code:         code,
			Name:         name,
			Level:        model.CodeTown,
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

func DealVillage(doc *html.Node, division model.Division) []model.Division {
	var data []model.Division
	matcher := matchByClass("class", "villagetr")
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		code := node.FirstChild.FirstChild.Data
		//vType := node.FirstChild.NextSibling.FirstChild.Data
		name := node.LastChild.FirstChild.Data
		var d = model.Division{
			Code:         code,
			Name:         name,
			Level:        model.CodeVillage,
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
	html.Render(w, n)
	return buf.String()
}
