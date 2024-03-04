package db

import (
	"errors"

	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

const BrandTable = "brand"

func CreateBrand(session db.Session, brand *models.Brand) (uint, error) {
	var id uint = 0
	if brand.Name == "" {
		return id, errors.New("brand name is required")
	}

	err := session.SQL().InsertInto(BrandTable).Columns(
		"name",
	).Values(
		brand.Name,
	).Returning("id").Iterator().ScanOne(&id)

	return id, err
}
