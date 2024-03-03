package db

import (
	"errors"
	"fmt"

	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

const StoreTable = "store"

func CreateStore(session db.Session, store *models.Store) error {
	name := store.Name
	if len(name) == 0 {
		return errors.New("store name is required")
	}

	_, err := session.SQL().InsertInto(StoreTable).Columns(
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
	).Exec()

	return err
}
