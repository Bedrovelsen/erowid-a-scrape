package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func main() {
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(`(?m)\Qhttps://erowid.org/experiences/exp.\E.+(ID=|Max=){1}\d{1,6}$`),
		),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("ExperienceVaultURL: ", r.URL.String())
	})

	c.Visit("https://erowid.org/experiences/exp.cgi?S1=15&ShowViews=0&Cellar=0&Start=0&Max=500")
}
