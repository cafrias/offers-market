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
		fmt.Sprintf("(%v, %v)", store.Location[0], store.Location[1]),
	).Returning("id").Iterator().ScanOne(&id)

	return id, err
}
