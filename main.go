package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Ed1123/us-visa-wait-times/table"
	"github.com/Ed1123/us-visa-wait-times/usvisa"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func waitTimes(w http.ResponseWriter, r *http.Request) {
	cities := usvisa.GetWaitData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func tableJS(w http.ResponseWriter, r *http.Request) {
	cities := usvisa.GetWaitDataWithCountry()
	w.Header().Set("Content-Type", "text/html")
	tmplFile := "templates/table.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatal(err)
	}
	tmplErr := tmpl.Execute(
		w, map[string]interface{}{"Data": cities, "Title": "US Visa Wait Times"},
	)
	if tmplErr != nil {
		log.Fatal(tmplErr)
	}
}

func templTable(w http.ResponseWriter, r *http.Request) {
	cities := usvisa.GetWaitData()
	table.Table(cities).Render(r.Context(), w)
}

func waitTimesWithCountry(w http.ResponseWriter, r *http.Request) {
	cities := usvisa.GetWaitDataWithCountry()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()

	r.HandleFunc("/table-js", tableJS)
	r.HandleFunc("/wait-times", waitTimes)
	r.HandleFunc("/wait-times-with-country", waitTimesWithCountry)

	r.HandleFunc("/table", templTable)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
