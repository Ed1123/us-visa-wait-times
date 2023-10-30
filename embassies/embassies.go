package embassies

import (
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type ConsulateEmbassy struct {
	Name    string
	City    string
	Country string
}

func getCityFromEmbassyName(embassyName string) string {
	nameStrings := strings.SplitN(embassyName, " in ", 2)
	if !(len(nameStrings) > 0) {
		log.Println("Could not split embassy name", embassyName)
		return ""
	}
	return nameStrings[len(nameStrings)-1]
}

func GetConsulatesEmbassies() []ConsulateEmbassy {
	c := colly.NewCollector()
	if os.Getenv("ENV") == "dev" {
		c.CacheDir = "./us_visa_cache"
	}
	embassies := []ConsulateEmbassy{}
	url := "https://www.embassy-worldwide.com/country/united-states"

	c.OnHTML("div.col-md-6", func(e *colly.HTMLElement) {
		countries := []string{}
		e.ForEach("h2", func(_ int, countryEl *colly.HTMLElement) {
			countries = append(countries, countryEl.Text)
		})
		e.ForEach("ul", func(i int, ulEl *colly.HTMLElement) {
			country := &countries[i]
			ulEl.ForEach("li", func(_ int, liEl *colly.HTMLElement) {
				embassyName := strings.Trim(liEl.Text, " \n")
				embassy := ConsulateEmbassy{
					Name:    embassyName,
					City:    getCityFromEmbassyName(embassyName),
					Country: *country,
				}
				embassies = append(embassies, embassy)
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.Visit(url)
	return embassies
}
