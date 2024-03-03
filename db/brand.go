package db

import (
	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

const BrandTable = "brand"

func CreateBrand(session db.Session, brand *models.Brand) error {
	_, err := session.SQL().InsertInto(BrandTable).Columns(
		"name",
	).Values(
		brand.Name,
	).Exec()
	return err
}
