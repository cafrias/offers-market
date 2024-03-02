package models

import "time"

type Offer struct {
	Id       uint
	StoreId  uint
	BrandId  uint
	Quantity uint
	// Available represents the number of items available for the offer
	Available      uint
	Name           string
	Price          uint
	Picture        string
	ExpirationDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
