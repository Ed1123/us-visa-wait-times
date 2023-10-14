package cityAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type response struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			Components struct {
				City    string
				Country string
			} `json:"components"`
			Annotations struct {
				Flag string
			} `json:"annotations"`
		} `json:"properties"`
	} `json:"features"`
}

type CityInfo struct {
	CityName  string
	Country   string
	Latitude  float64
	Longitude float64
}

func GetCityInfo(city string) CityInfo {
	// Retrieves information about a city using the OpenCage Geocoding API.
	// It takes a city name as input and returns a CityInfo struct containing the city's name, country, latitude, and longitude.
	version := 1
	format := "geojson"
	u, err := url.Parse(fmt.Sprintf("https://api.opencagedata.com/geocode/v%d/%s", version, format))
	if err != nil {
		log.Fatalln("Couldn't parse url", err)
	}
	q := u.Query()
	q.Add("q", city)
	q.Add("no_annotations", "1")
	q.Add("limit", "1")
	q.Add("key", os.Getenv("OPENCAGE_API_KEY"))

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln("Couldn't get response", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Couldn't read response body", err)
	}
	var r response
	if err := json.Unmarshal(body, &r); err != nil {
		log.Fatalln("Couldn't unmarshal response", err)
	}

	// fmt.Println(r.Features[0].Properties.Annotations.Flag)

	return CityInfo{
		CityName:  city,
		Country:   r.Features[0].Properties.Components.Country,
		Latitude:  r.Features[0].Geometry.Coordinates[0],
		Longitude: r.Features[0].Geometry.Coordinates[1],
	}
}
