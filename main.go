package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/cczyWyc/crawler-geektime/collect"
	"time"
)

func main() {
	url := "https://book.douban.com/subject/1007305/"
	var f collect.FetCher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read contect faield: %v", err)
		return
	}
	fmt.Println(body)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("read content failed: %v", err)
	}
	doc.Find("div.news_li h2 a[target=_blank]").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Review %d %s\n", i, title)
	})
}
