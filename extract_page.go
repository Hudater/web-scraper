package main

import (
	"fmt"
	"log"
	"net/url"
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

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}
	// THIS signature is copy pasta from boot.dev ch1-l5
	// For each '<a href>' it finds, it will run this function.
	urlSlice := []string{}
	var docErr error
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		parsedHref, err := url.Parse(href)
		if err != nil {
			docErr = fmt.Errorf("unable to parse href %q: %w", href, err)
			return
		}
		fullURL := baseURL.ResolveReference(parsedHref).String()
		urlSlice = append(urlSlice, fullURL)
	})
	if docErr != nil {
		return nil, docErr
	}
	return urlSlice, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}
	// For each '<img src>' it finds, it will run this function.
	imageURLSlice := []string{}
	var docErr error
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		parsedSrc, err := url.Parse(src)
		if err != nil {
			docErr = fmt.Errorf("unable to parse image src %q: %w", src, err)
			return
		}
		fullURL := baseURL.ResolveReference(parsedSrc).String()
		imageURLSlice = append(imageURLSlice, fullURL)
	})
	if docErr != nil {
		return nil, docErr
	}
	return imageURLSlice, nil
}

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		log.Printf("unable to parse URL")
		return PageData{}
	}
	
	linksSlice, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		log.Printf("unable to get links")
	}
	 
	imagesSlice, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		log.Printf("unable to get images")
	}
	
	pageDataStruct := PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  linksSlice,
		ImageURLs:      imagesSlice,
	}
	return pageDataStruct
}
