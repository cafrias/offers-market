package models

import "time"

type Offer struct {
	Id       uint `db:"id"`
	StoreId  uint `db:"store_id"`
	BrandId  uint `db:"brand_id"`
	Quantity uint `db:"quantity"`
	// Available represents the number of items available for the offer
	Available      uint      `db:"available"`
	Name           string    `db:"name"`
	Price          uint      `db:"price"`
	Picture        string    `db:"picture"`
	ExpirationDate time.Time `db:"expiration_date"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
