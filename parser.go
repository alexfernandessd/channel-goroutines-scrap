package main

import (
	"strings"

	"golang.org/x/net/html"
)

// GetFirstTextNode ...
func GetFirstTextNode(htmlParsed *html.Node) *html.Node {
	if htmlParsed == nil {
		return nil
	}

	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Type == html.TextNode {
			return m
		}
		r := GetFirstTextNode(m)
		if r != nil {
			return r
		}
	}
	return nil
}

// GetFirstElementByClass ...
func GetFirstElementByClass(htmlParsed *html.Node, elm, className string) *html.Node {
	for m := htmlParsed.FirstChild; m != nil; m = m.NextSibling {
		if m.Data == elm && HasClass(m.Attr, className) {
			return m
		}
		r := GetFirstElementByClass(m, elm, className)
		if r != nil {
			return r
		}
	}
	return nil
}

// HasClass ...
func HasClass(attribs []html.Attribute, className string) bool {
	for _, attr := range attribs {
		if attr.Key == "class" && strings.Contains(attr.Val, className) {
			return true
		}
	}
	return false
}
