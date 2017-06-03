package models

import (
	"log"
	"time"
)

type Price struct {
	ID        int       `json:"id"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

func AddPrice(amount string) (bool, error) {
	sqlStatement := `INSERT INTO prices (price, created_at) VALUES($1, $2)`

	_, err := db.Exec(sqlStatement, amount, time.Now())

	if err != nil {
		log.Panic(err)
	}

	return true, nil
}

func AllPrices() ([]*Price, error) {
	rows, err := db.Query("SELECT * FROM prices")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	prices := make([]*Price, 0)

	for rows.Next() {
		price := new(Price)
		err = rows.Scan(&price.ID, &price.Price, &price.CreatedAt)

		if err != nil {
			return nil, err
		}

		prices = append(prices, price)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}
