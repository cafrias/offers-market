package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/cafrias/offers-market/db"
	"github.com/cafrias/offers-market/models"
)

const NumberOfBrands = 100
const NumberOfStores = 100
const NumberOfOffers = 1000

func main() {
	gofakeit.Seed(0)

	session, err := db.Connect()
	if err != nil {
		panic(err)
	}

	brands := make([]models.Brand, 0, NumberOfBrands)
	for range NumberOfBrands {
		brand := createBrand()
		brands = append(brands, brand)
	}

	brandIds := make([]uint, 0, NumberOfBrands)
	for _, brand := range brands {
		rId, err := db.CreateBrand(session, &brand)
		if err != nil {
			panic(err)
		}

		brandIds = append(brandIds, rId)
	}

	stores := make([]models.Store, 0, NumberOfStores)
	for range NumberOfStores {
		store := createStore()
		stores = append(stores, store)
	}

	storeIds := make([]uint, 0, NumberOfStores)
	for _, store := range stores {
		rId, err := db.CreateStore(session, &store)
		if err != nil {
			panic(err)
		}
		storeIds = append(storeIds, rId)
	}

	offers := make([]models.Offer, 0, NumberOfOffers)
	for range NumberOfOffers {
		offer := createOffer(brandIds, storeIds)
		offers = append(offers, offer)
	}

	for _, offer := range offers {
		_, err := db.CreateOffer(session, &offer)
		if err != nil {
			panic(err)
		}
	}
}

func createBrand() models.Brand {
	return models.Brand{
		Name: gofakeit.Company(),
	}
}

func createStore() models.Store {
	return models.Store{
		Name:     gofakeit.Company(),
		Address:  gofakeit.Address().Address,
		Phone:    gofakeit.Phone(),
		Email:    gofakeit.Email(),
		Website:  gofakeit.URL(),
		Location: [2]float64{gofakeit.Latitude(), gofakeit.Longitude()},
	}
}

func createOffer(brandIds []uint, storeIds []uint) models.Offer {
	picSeed := gofakeit.UUID()
	return models.Offer{
		BrandId:        brandIds[gofakeit.Number(0, len(brandIds)-1)],
		StoreId:        uint(storeIds[gofakeit.Number(0, len(storeIds)-1)]),
		Price:          uint(gofakeit.Number(10_000, 10_000_000)),
		Quantity:       uint(gofakeit.Number(1, 12)),
		Available:      uint(gofakeit.Number(1, 200)),
		Name:           gofakeit.ProductName(),
		Picture:        fmt.Sprintf("https://picsum.photos/seed/%v/200/300", picSeed),
		ExpirationDate: time.Now().Add(time.Duration(gofakeit.Number(1, 30)) * 24 * time.Hour),
	}
}
