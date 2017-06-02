package models

import (
	"log"
	"time"
)

type Price struct {
	ID        int
	Price     string
	CreatedAt time.Time
}

func AddPrice(amount string) (bool, error) {
	sqlStatement := `INSERT INTO prices (price, created_at) VALUES($1, $2)`

	_, err := db.Exec(sqlStatement, amount, time.Now())

	if err != nil {
		log.Panic(err)
	}

	return true, nil
}
