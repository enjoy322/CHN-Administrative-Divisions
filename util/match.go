package util

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

func MatcherByAtom(param atom.Atom) func(node *html.Node) (keep bool, exit bool) {
	matcher := func(node *html.Node) (keep bool, exit bool) {
		if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
			exit = true
		}
		if node.DataAtom == param {
			keep = true
		}
		return
	}
	return matcher
}

func MatchByClass(findAttrKey string, findAttrVal string) func(node *html.Node) (keep bool, exit bool) {
	matcher := func(node *html.Node) (keep bool, exit bool) {
		if node.Type == html.ElementNode {
			var s string
			var ok bool
			for _, attr := range node.Attr {
				if attr.Key == findAttrKey {
					s = attr.Val
					ok = true
					break
				}
			}

			if ok && s == findAttrVal {
				keep = true
			}

		}
		return
	}
	return matcher
}
