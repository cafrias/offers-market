package db

import (
	"errors"
	"fmt"

	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

const StoreTable = "store"

func CreateStore(session db.Session, store *models.Store) (uint, error) {
	var id uint = 0
	name := store.Name
	if len(name) == 0 {
		return id, errors.New("store name is required")
	}

	err := session.SQL().InsertInto(StoreTable).Columns(
		"name",
		"address",
		"phone",
		"email",
		"website",
		"location",
	).Values(
		store.Name,
		store.Address,
		store.Phone,
		store.Email,
		store.Website,
		fmt.Sprintf("(%v, %v)", store.Lng, store.Lat),
	).Returning("id").Iterator().ScanOne(&id)

	return id, err
}

func GetStore(session db.Session, id uint) (*models.Store, error) {
	store := &models.Store{}
	err := session.SQL().
		Select(
			"*",
			db.Raw("(location[0]) as lng"),
			db.Raw("(location[1]) as lat"),
		).
		From(StoreTable).
		Where("id", id).
		One(store)
	return store, err
}
