package cityAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type response struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64
		}
		Properties struct {
			Components struct {
				City    string
				Country string
			}
			Annotations struct {
				Flag string
			}
		}
	}
}

type CityInfo struct {
	CityName string
	Country  string
	//Flag	  string
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
	// q.Add("no_annotations", "1") // uncomment to include flag
	q.Add("limit", "1")

	apiKey, isApiKey := os.LookupEnv("OPENCAGE_API_KEY")
	if !isApiKey {
		log.Fatalln("Couldn't find OPENCAGE_API_KEY in environment variables")
	}
	q.Add("key", apiKey)

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
	if len(r.Features) == 0 {
		log.Println("No city info found for", city)
		return CityInfo{CityName: city}
	}
	return CityInfo{
		CityName:  city,
		Country:   r.Features[0].Properties.Components.Country,
		Latitude:  r.Features[0].Geometry.Coordinates[0],
		Longitude: r.Features[0].Geometry.Coordinates[1],
	}
}

type CityInfoCache struct {
	cache map[string]CityInfo
	file  string
}

func NewCityInfoCache() CityInfoCache {
	cityCache := CityInfoCache{
		cache: make(map[string]CityInfo),
		file:  "./.cache/city_info_cache.json",
	}
	cityCache.readFromDisk()
	return cityCache
}

func (c CityInfoCache) getCityInfo(city string) (CityInfo, bool) {
	ci, ok := c.cache[city]
	return ci, ok
}

func (c CityInfoCache) setCityInfo(city string, ci CityInfo) {
	c.cache[city] = ci
	c.writeToDisk()
}

func (c CityInfoCache) readFromDisk() error {
	f, err := os.Open(c.file)
	if os.IsNotExist(err) {
		dir := filepath.Dir(c.file)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalln("Couldn't create cache directory", err)
		}
		c.writeToDisk()
		return nil
	}
	if err != nil {
		log.Fatalln("Couldn't open cache file.", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&c.cache)
	if err != nil {
		log.Fatalln("Couldn't decode cache file", err)
	}
	return nil
}

func (c CityInfoCache) writeToDisk() {
	f, err := os.Create(c.file)
	if err != nil {
		log.Fatalln("Couldn't open cache file", err)
	}
	err = json.NewEncoder(f).Encode(c.cache)
	if err != nil {
		log.Fatalln("Couldn't encode cache file", err)
	}
}
func (c CityInfoCache) GetCityInfo(city string) CityInfo {
	cityInfo, inCache := c.getCityInfo(city)
	if inCache {
		log.Println("Getting city info from cache.")
		return cityInfo
	} else {
		log.Println("Getting city info from API for", city, "and caching it.")
		cityInfo := GetCityInfo(city)
		c.setCityInfo(city, cityInfo)
		return cityInfo
	}
}
