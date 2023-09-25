package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type WaitTimeDetails struct {
	days    int16
	hasDays bool
	message string
}

type CityWaitingTime struct {
	cityName                string
	studentExchangeVisitor  WaitTimeDetails
	petitionBasedTempWorker WaitTimeDetails
	crewTransit             WaitTimeDetails
	businessTourismVisitor  WaitTimeDetails
	// Country                 string
}

func parseWaitingTime(str string) WaitTimeDetails {
	if v, ok := strconv.Atoi(strings.Split(str, " ")[0]); ok == nil {
		return WaitTimeDetails{int16(v), true, ""}
	} else {
		return WaitTimeDetails{0, false, str}
	}
}

func main() {
	c := colly.NewCollector(colly.CacheDir("./us_visa_cache"))

	cities := []CityWaitingTime{}

	// Find and visit all links
	c.OnHTML(
		"div.tsg-rwd-body-frame-row div div.tsg-rwd-main-copy-frame.dataCSIpage div.tsg-rwd-main-copy-body-frame.no-rail.dataCSIpage div.tsg-rwd-content-page-parsysxxx.parsys div.tsg-rwd-text.parbase.section div p table tbody",
		func(e *colly.HTMLElement) {
			e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
				cityWaitTime := CityWaitingTime{
					e.ChildText("td:nth-child(1)"),
					parseWaitingTime(e.ChildText("td:nth-child(2)")),
					parseWaitingTime(e.ChildText("td:nth-child(3)")),
					parseWaitingTime(e.ChildText("td:nth-child(4)")),
					parseWaitingTime(e.ChildText("td:nth-child(5)")),
				}
				cities = append(cities, cityWaitTime)
			})
		})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://travel.state.gov/content/travel/en/us-visas/visa-information-resources/global-visa-wait-times.html")

	jsonBytes, err := json.Marshal(cities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}
