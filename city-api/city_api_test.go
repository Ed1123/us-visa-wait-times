package cityAPI

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestCityInfo(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cityInfo := GetCityInfo("New York")

	if cityInfo.Country != "United States" {
		t.Error(cityInfo)
		t.Errorf("Expected country to be United States, got %s", cityInfo.Country)
	}
}
