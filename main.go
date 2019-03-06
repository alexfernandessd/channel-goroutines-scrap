package main

import (
	"fmt"
	"time"
)

// Response ...
type Response struct {
	Title  string
	Author string
	Claps  string
}

func main() {
	urls := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		// "https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}

	ch := make(chan Response)

	ini := time.Now()

	go scrapList(urls, ch)

	for resp := range ch {
		fmt.Println(resp)
	}

	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")
}

func scrapList(urls []string, in chan Response) {
	defer close(in)

	var outs = []chan Response{}

	for i, url := range urls {
		outs = append(outs, make(chan Response))
		go processURL(url, outs[i])
	}

	for i := range outs {
		for response := range outs[i] {
			in <- response
		}
	}
}

func processURL(url string, in chan Response) {
	defer close(in)
	var resp Response

	htmlParsed, _ := GetHTMLParsed(url)

	a := GetFirstElementByClass(htmlParsed, "a", "link link--primary u-accentColor--hoverTextNormal")
	resp.Author = GetFirstTextNode(a).Data

	div := GetFirstElementByClass(htmlParsed, "div", "section-content")
	h1 := GetFirstElementByClass(div, "h1", "graf--title")
	resp.Title = GetFirstTextNode(h1).Data

	footer := GetFirstElementByClass(htmlParsed, "footer", "u-paddingTop10")
	buttonLikes := GetFirstElementByClass(footer, "button", "js-multirecommendCountButton")
	resp.Claps = GetFirstTextNode(buttonLikes).Data

	in <- resp
}
