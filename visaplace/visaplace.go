package visaplace

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

type ConsulateEmbassy struct {
	Name    string
	City    string
	Country string
}

func GetConsulatesEmbassies() []ConsulateEmbassy {
	c := colly.NewCollector(colly.CacheDir("./us_visa_cache"))
	embassies := []ConsulateEmbassy{}
	url := "https://www.visaplace.com/usa-immigration/resources/embassy/"

	c.OnHTML(".embassy-list", func(e *colly.HTMLElement) {
		e.ForEach("div", func(_ int, countryEl *colly.HTMLElement) {
			country := countryEl.ChildText("h3 > a")
			countryEl.ForEach("ul > li", func(_ int, embassyEl *colly.HTMLElement) {
				embassyName := embassyEl.Text
				embassy := ConsulateEmbassy{
					Name:    embassyName,
					City:    getCityFromEmbassyName(embassyName),
					Country: country,
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

func getCityFromEmbassyName(embassyName string) string {
	re := regexp.MustCompile(`in (.*)`)
	matches := re.FindStringSubmatch(embassyName)
	return matches[len(matches)-1]
}
