package visaplace

import "testing"

func TestGetConsulatesEmbassies(t *testing.T) {
	embassies := GetConsulatesEmbassies()

	if len(embassies) == 0 {
		t.Error("Expected non-empty list of embassies")
	}

	for _, embassy := range embassies {
		if embassy.Name == "" {
			t.Error("Expected non-empty embassy name")
		}
		if embassy.City == "" {
			t.Error("Expected non-empty embassy city")
		}
		if embassy.Country == "" {
			t.Error("Expected non-empty embassy country")
		}
	}
}
