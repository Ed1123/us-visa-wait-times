package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.CacheDir("./us_visa_cache"))

	// Find and visit all links
	c.OnHTML(
		"div.tsg-rwd-body-frame-row div div.tsg-rwd-main-copy-frame.dataCSIpage div.tsg-rwd-main-copy-body-frame.no-rail.dataCSIpage div.tsg-rwd-content-page-parsysxxx.parsys div.tsg-rwd-text.parbase.section div p table tbody",
		func(e *colly.HTMLElement) {
			e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
				fmt.Println(e.ChildText("td:nth-child(1)"))
				fmt.Println(e.ChildText("td:nth-child(2)"))
				fmt.Println(e.ChildText("td:nth-child(3)"))
				fmt.Println(e.ChildText("td:nth-child(4)"))
				fmt.Println(e.ChildText("td:nth-child(5)"))
				fmt.Println("end of row")
			})
		})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://travel.state.gov/content/travel/en/us-visas/visa-information-resources/global-visa-wait-times.html")
}
