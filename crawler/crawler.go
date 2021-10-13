package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const baseURL = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/"

func main() {
	contentProvince := get(baseURL)

	doc, err := html.Parse(strings.NewReader(string(contentProvince)))
	if err != nil {
		fmt.Println(err)
		return
	}
	matcher := matchByClass("class", "provincetr")

	var dataWithHtml []string
	nodes := TraverseNode(doc, matcher)
	for _, node := range nodes {
		dataWithHtml = append(dataWithHtml, renderNode(node))
	}

	matcherA := matcherByAtom(atom.A)

	var provinceList []Division

	for _, str := range dataWithHtml {
		doc2, _ := html.Parse(strings.NewReader(str))
		nodes2 := TraverseNode(doc2, matcherA)
		for _, node := range nodes2 {
			url := node.Attr[0].Val
			code := url[:len(url)-5]
			if len(code) < 12 {
				b := strings.Builder{}
				b.WriteString(code)
				for i := 0; i < 12-len(code); i++ {
					b.WriteString("0")
				}
				code = b.String()
			}
			provinceList = append(provinceList, Division{
				Url:   url,
				Code:  code,
				Level: 1,
				Name:  node.FirstChild.Data,
			})
		}
	}
	fmt.Println(provinceList)
	//写入文件
	dir := "./file"
	err = MkDir(dir)
	if err != nil {
		panic("创建失败")
		return
	}
	marshal, err := json.Marshal(provinceList)
	if err != nil {
		return
	}
	//生成json文件
	err = ioutil.WriteFile(dir+"/province.json", marshal, os.ModeDir)
	if err != nil {
		panic("写入失败"+err.Error())
	}

	for _, province := range provinceList {
		if province.Code != "530000000000" {
			continue
		}
		cityList := CrawlCity(province)
		for _, city := range cityList {
			countyList := CrawlCounty(city)
			for _, county := range countyList {
				townList := CrawlTown(county)
				for _, _ = range townList {
					//villageList := CrawlVillage(town)
					//fmt.Println(villageList)
				}
			}
		}
	}

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
