package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Ed1123/us-visa-wait-times/usvisa"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	r.HandleFunc("/table-test", tableTest)
	r.HandleFunc("/wait-times", waitTimes)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
