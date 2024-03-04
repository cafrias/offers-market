package main

import (
	"log"

	"github.com/cafrias/offers-market/db"
)

func main() {
	session, err := db.Connect()
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer session.Close()

	offers, err := db.GetAvailableOffers(session, 1, 15)
	if err != nil {
		log.Fatalf("db.GetAvailableOffers(): %q\n", err)
	}

	for _, offer := range offers {
		log.Printf("Offer: %+v\n", offer)
	}

	// store := models.Store{
	// 	Name:    "Store 2",
	// 	Address: "Address 1",
	// }
	// err = db.CreateStore(session, &store)
	// if err != nil {
	// 	log.Fatalf("db.CreateStore(): %q\n", err)
	// }

	// brand := models.Brand{
	// 	Name: "Brand 1",
	// }

	// err = db.CreateBrand(session, &brand)
	// if err != nil {
	// 	log.Fatalf("db.CreateBrand(): %q\n", err)
	// }

	// offer := models.Offer{
	// 	Name:           "Offer 1",
	// 	BrandId:        2,
	// 	StoreId:        1,
	// 	Price:          100,
	// 	Quantity:       10,
	// 	Available:      10,
	// 	ExpirationDate: time.Now().Add(time.Hour * 24),
	// 	Picture:        "https://www.google.com",
	// }

	// err = db.CreateOffer(session, &offer)
	// if err != nil {
	// 	log.Fatalf("db.CreateOffer(): %q\n", err)
	// }
}
