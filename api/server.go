package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlows/dolartoday-worker/models"
	"github.com/gorilla/mux"
)

func GetPricesEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Printf(
		"%s\t%s",
		req.Method,
		req.RequestURI,
	)

	prices, err := models.AllPrices()

	if err != nil {
		log.Panic(err)
	}

	json.NewEncoder(w).Encode(prices)
}

func main() {
	models.InitDB("user=postgres dbname=dolartoday_worker sslmode=disable")

	router := mux.NewRouter()

	router.HandleFunc("/prices", GetPricesEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
