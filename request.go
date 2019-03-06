package main

import (
	"net/http"

	"golang.org/x/net/html"
)

// GetHTMLParsed ...
func GetHTMLParsed(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	htmlParsed, pErr := html.Parse(resp.Body)
	if pErr != nil {
		return nil, pErr
	}

	return htmlParsed, nil
}
