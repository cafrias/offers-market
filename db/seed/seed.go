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
const NumberOfProducts = 100

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

	products := [NumberOfProducts]models.Product{}
	for i := range products {
		product := createProduct(brandIds)
		products[i] = product
	}
	productIds := [NumberOfProducts]uint{}
	for i, product := range products {
		rId, err := db.CreateProduct(session, &product)
		if err != nil {
			panic(err)
		}

		productIds[i] = rId
	}

	offers := make([]models.Offer, 0, NumberOfOffers)
	for range NumberOfOffers {
		offer := createOffer(productIds, storeIds)
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
		Name:    gofakeit.Company(),
		Address: gofakeit.Address().Address,
		Phone:   gofakeit.Phone(),
		Email:   gofakeit.Email(),
		Website: gofakeit.URL(),
		Lat:     gofakeit.Latitude(),
		Lng:     gofakeit.Longitude(),
	}
}

func createOffer(productIds [NumberOfProducts]uint, storeIds []uint) models.Offer {
	return models.Offer{
		StoreId:        uint(storeIds[gofakeit.Number(0, len(storeIds)-1)]),
		Price:          uint(gofakeit.Number(10_000, 10_000_000)),
		Quantity:       uint(gofakeit.Number(1, 12)),
		Available:      uint(gofakeit.Number(1, 200)),
		ProductId:      uint(productIds[gofakeit.Number(0, len(productIds)-1)]),
		Description:    gofakeit.Sentence(10),
		ExpirationDate: time.Now().Add(time.Duration(gofakeit.Number(1, 30)) * 24 * time.Hour),
	}
}

func createProduct(brandIds []uint) models.Product {
	picSeed := gofakeit.UUID()
	return models.Product{
		Name:        gofakeit.ProductName(),
		BrandId:     brandIds[gofakeit.Number(0, len(brandIds)-1)],
		Picture:     fmt.Sprintf("https://picsum.photos/seed/%v/200/300", picSeed),
		Description: gofakeit.Sentence(10),
	}
}
