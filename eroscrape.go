package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	reportIndexURL := "https://erowid.org/experiences/exp.cgi?S1=15&ShowViews=0&Cellar=0&Start=0&Max=500"

	fmt.Println("Starting erowid-a-scrape\n[erowid experience report text corpus generation]\n\nReport index base URL:", reportIndexURL)

	reportURLs, err := getReportURLS(reportIndexURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Finished getting report URLs")
	ioutil.WriteFile("links.lst", []byte(strings.Join(reportURLs, "\n")), 0666)
	fmt.Println("Report URLs written to file links.lst")

	corpus, err := getReportText(reportURLs)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Finished fetching reports into text corpus")
	ioutil.WriteFile("corpus.txt", []byte(corpus), 0666)
	fmt.Println("Text corpus written to file corpus.txt")

	fmt.Println("Erowid-a-scrape finished")
}

func getReportURLS(reportIndexURL string) ([]string, error) {
	reportURLs := []string{}

	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(`(?m)\Qhttps://erowid.org/experiences/exp.\E.+(ID=|Max=){1}\d{1,6}$`),
		),
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		reportURLs = append(reportURLs, r.URL.String())
	})

	c.Visit(reportIndexURL)

	return reportURLs[1:], nil
}

func getReportText(reportURLs []string) (string, error) {
	corpusText := ""

	c := colly.NewCollector(
		colly.AllowedDomains("erowid.org"),
	)

	c.OnResponse(func(r *colly.Response) {
		bodyString := string(r.Body)

		var re = regexp.MustCompile(`(?m)\Q<!-- Start Body -->\E(\n|\w|\W)*\Q<!-- End Body -->\E`)
		for _, match := range re.FindAllString(bodyString, -1) {
			corpusText = corpusText + strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(match, "<br>", "", -1), "<BR>", "", -1), "<!-- Start Body -->", "", -1), "<!-- End Body -->", "", -1), ". ", ".\n", -1), "\r\n", "\r", -1) + "\n"
		}
	})

	for _, curLink := range reportURLs {
		c.Visit(curLink)
	}

	return corpusText, nil
}
