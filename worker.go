package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
)

func requestDolarCotization() float64 {
	doc, err := goquery.NewDocument("https://twitter.com/DolarToday")

	if err != nil {
		log.Fatal(err)
	}

	content := doc.Find(".ProfileHeaderCard-bio.u-dir").Contents().Text()

	return parseDolarCotization(content)
}

func parseDolarCotization(text string) float64 {
	r, _ := regexp.Compile("\\d+,\\d+")
	amountToParse := r.FindString(text)
	validAmount := strings.Replace(amountToParse, ",", ".", -1)
	amount, err := strconv.ParseFloat(validAmount, 64)

	if err != nil {
		log.Fatal(err)
	}

	return amount
}

func main() {
	c := cron.New()

	c.AddFunc("0 0 8-16 ? * 1-5", func() {
		fmt.Println(requestDolarCotization())
	})

	go c.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
