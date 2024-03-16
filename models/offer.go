package models

import "time"

type Offer struct {
	Id        uint `db:"id"`
	StoreId   uint `db:"store_id"`
	ProductId uint `db:"product_id"`
	Quantity  uint `db:"quantity"`
	// Available represents the number of items available for the offer
	Available      uint      `db:"available"`
	Price          uint      `db:"price"`
	Description    string    `db:"description"`
	ExpirationDate time.Time `db:"expiration_date"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
