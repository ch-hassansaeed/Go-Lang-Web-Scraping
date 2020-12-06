package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print url link which exist in visiting site
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only visit those links which are in AllowedDomains list
		//c.Visit(e.Request.AbsoluteURL(link)) //use it if you again want to visit found links
	})

	// for each request print "Visiting then url of site"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// for example send scraping request on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}