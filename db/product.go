package db

import (
	"errors"

	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

const ProductTable = "product"

func CreateProduct(session db.Session, product *models.Product) (id uint, err error) {
	name := product.Name
	if len(name) == 0 {
		return 0, errors.New("product name is required")
	}

	bId := product.BrandId
	if bId == 0 {
		return 0, errors.New("product brand_id is required")
	}

	picture := product.Picture
	if len(picture) == 0 {
		return 0, errors.New("product picture is required")
	}

	err = session.SQL().InsertInto(ProductTable).Columns(
		"name",
		"brand_id",
		"picture",
		"description",
	).Values(
		name,
		bId,
		picture,
		product.Description,
	).Returning("id").Iterator().ScanOne(&id)

	return id, err
}
