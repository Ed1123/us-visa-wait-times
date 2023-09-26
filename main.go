package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func getCountryForCity(cityName string) string {
	return cityName
}

type Days int16
type Message string

type WaitingTime struct {
	Days    *Days    `json:"days,omitempty"`
	Message *Message `json:"message,omitempty"`
}

type CityWaitingTime struct {
	CityName                string
	Country                 string
	StudentExchangeVisitor  WaitingTime
	PetitionBasedTempWorker WaitingTime
	CrewTransit             WaitingTime
	BusinessTourismVisitor  WaitingTime
}

func parseWaitingTime(str string) WaitingTime {
	if v, ok := strconv.Atoi(strings.Split(str, " ")[0]); ok == nil {
		days := Days(int16(v))
		return WaitingTime{&days, nil}
	} else if str != "" {
		message := Message(str)
		return WaitingTime{nil, &message}
	}
	return WaitingTime{nil, nil}
}

func main() {
	c := colly.NewCollector(colly.CacheDir("./us_visa_cache"))

	cities := []CityWaitingTime{}

	// Find and visit all links
	c.OnHTML(
		"div.tsg-rwd-body-frame-row div div.tsg-rwd-main-copy-frame.dataCSIpage div.tsg-rwd-main-copy-body-frame.no-rail.dataCSIpage div.tsg-rwd-content-page-parsysxxx.parsys div.tsg-rwd-text.parbase.section div p table tbody",
		func(e *colly.HTMLElement) {
			e.ForEach("tr:nth-child(n+2)", func(_ int, el *colly.HTMLElement) {
				cityName := el.ChildText("td:nth-child(1)")
				cityWaitTime := CityWaitingTime{
					cityName,
					getCountryForCity(cityName),
					parseWaitingTime(el.ChildText("td:nth-child(2)")),
					parseWaitingTime(el.ChildText("td:nth-child(3)")),
					parseWaitingTime(el.ChildText("td:nth-child(4)")),
					parseWaitingTime(el.ChildText("td:nth-child(5)")),
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
