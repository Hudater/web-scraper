package main

import (
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	headings := doc.Find("h1").First().Text()
	if len(headings) > 0 {
		return headings
	}
	headings = doc.Find("h2").First().Text()
	if len(headings) > 0 {
		return headings
	}
	return ""
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	paras := doc.Find("main p").First().Text()
	if len(paras) > 0 {
		return paras
	}
	paras = doc.Find("p").First().Text()
	if len(paras) > 0 {
		return paras
	}
	return ""
}
