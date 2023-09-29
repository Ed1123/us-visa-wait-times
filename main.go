package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/Ed1123/us-visa-wait-times/usvisa"
)

func waitTimes(w http.ResponseWriter, r *http.Request) {
	cities := usvisa.GetWaitData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func tableTest(w http.ResponseWriter, r *http.Request) {
	cities := usvisa.GetWaitData()
	w.Header().Set("Content-Type", "text/html")
	tmplFile := "templates/table.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatal(err)
	}
	tmplErr := tmpl.Execute(
		w, struct {
			Data  []usvisa.CityWaitTime
			Title string
		}{cities, "US Visa Wait Times"},
	)
	if tmplErr != nil {
		log.Fatal(tmplErr)
	}
}

func main() {
	http.HandleFunc("/table-test", tableTest)
	http.HandleFunc("/wait-times", waitTimes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
