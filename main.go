package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/carlows/dolartoday-worker/models"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)

func GetPricesEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Printf(
		"%s\t%s",
		req.Method,
		req.RequestURI,
	)

	prices, err := models.AllPricesChart()

	if err != nil {
		log.Panic(err)
	}

	json.NewEncoder(w).Encode(prices)
}

func requestDolarCotization() string {
	doc, err := goquery.NewDocument("https://twitter.com/DolarToday")

	if err != nil {
		log.Fatal(err)
	}

	content := doc.Find(".ProfileHeaderCard-bio.u-dir").Contents().Text()

	return parseDolarCotization(content)
}

func parseDolarCotization(text string) string {
	r, _ := regexp.Compile("\\d+,\\d+")
	amount := r.FindString(text)

	return amount
}

func main() {
	fmt.Println("Init dolartoday-api and worker")

	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		dbUrl = "user=postgres dbname=dolartoday_worker sslmode=disable"
	}

	models.InitDB(dbUrl)

	c := cron.New()

	c.AddFunc("0 0 6-18 ? * 1-5", func() {
		models.AddPrice(requestDolarCotization())
		fmt.Println("Adding dolar price: ", requestDolarCotization())
	})

	go c.Start()

	router := mux.NewRouter()
	router.HandleFunc("/prices", GetPricesEndpoint).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
