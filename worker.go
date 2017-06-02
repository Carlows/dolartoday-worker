package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"

	"github.com/Carlows/dolartoday-worker/models"
	"github.com/robfig/cron"

	"github.com/PuerkitoBio/goquery"
)

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
	models.InitDB("user=postgres dbname=dolartoday_worker sslmode=disable")

	c := cron.New()

	c.AddFunc("0 0 8-16 ? * 1-5", func() {
		models.AddPrice(requestDolarCotization())
		fmt.Println("Adding dolar price: ", requestDolarCotization())
	})

	go c.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
