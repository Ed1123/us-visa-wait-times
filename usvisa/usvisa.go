package usvisa

import (
	"log"
	"strconv"
	"strings"

	"github.com/Ed1123/us-visa-wait-times/embassies"
	"github.com/gocolly/colly"
)

type Days int16
type Message string

type WaitTime struct {
	Days    *Days    `json:"days,omitempty"`
	Message *Message `json:"message,omitempty"`
}

type CityWaitTime struct {
	CityName string
	// Country                 string
	StudentExchangeVisitor  WaitTime
	PetitionBasedTempWorker WaitTime
	CrewTransit             WaitTime
	BusinessTourismVisitor  WaitTime
}

func parseWaitTime(str string) WaitTime {
	if v, ok := strconv.Atoi(strings.Split(str, " ")[0]); ok == nil {
		days := Days(int16(v))
		return WaitTime{&days, nil}
	} else if str != "" {
		message := Message(str)
		return WaitTime{nil, &message}
	}
	return WaitTime{nil, nil}
}

func GetWaitData() []CityWaitTime {
	c := colly.NewCollector(colly.CacheDir("./us_visa_cache"))

	cities := []CityWaitTime{}

	// Find and visit all links
	c.OnHTML(
		"p table tbody",
		func(e *colly.HTMLElement) {
			e.ForEach("tr:nth-child(n+2)", func(_ int, el *colly.HTMLElement) {
				cityName := el.ChildText("td:nth-child(1)")
				cityWaitTime := CityWaitTime{
					cityName,
					// getCountryForCity(cityName),
					parseWaitTime(el.ChildText("td:nth-child(2)")),
					parseWaitTime(el.ChildText("td:nth-child(3)")),
					parseWaitTime(el.ChildText("td:nth-child(4)")),
					parseWaitTime(el.ChildText("td:nth-child(5)")),
				}
				cities = append(cities, cityWaitTime)
			})
		})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("https://travel.state.gov/content/travel/en/us-visas/visa-information-resources/global-visa-wait-times.html")

	return cities
}

type EmbassyInfo struct {
	cityEmbassyMap map[string]embassies.ConsulateEmbassy
}

func newEmbassyInfo() EmbassyInfo {
	embassiesArray := embassies.GetConsulatesEmbassies()
	cityEmbassyMap := make(map[string]embassies.ConsulateEmbassy)
	for _, embassy := range embassiesArray {
		cityEmbassyMap[embassy.City] = embassy
	}
	return EmbassyInfo{cityEmbassyMap}
}

func (e EmbassyInfo) getEmbassy(cityName string) embassies.ConsulateEmbassy {
	return e.cityEmbassyMap[cityName]
}

type CityWaitTimeWithCountry struct {
	CityWaitTime
	Country string
}

func GetWaitDataWithCountry() []CityWaitTimeWithCountry {
	cites := GetWaitData()
	citiesWithCountry := []CityWaitTimeWithCountry{}
	embassyInfo := newEmbassyInfo()
	for _, city := range cites {
		embassy := embassyInfo.getEmbassy(city.CityName)
		citiesWithCountry = append(
			citiesWithCountry,
			CityWaitTimeWithCountry{city, embassy.Country},
		)
	}
	return citiesWithCountry
}
