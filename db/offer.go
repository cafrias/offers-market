package db

import (
	"database/sql"
	"errors"

	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

const OfferTable = "offer"

func CreateOffer(session db.Session, offer *models.Offer) (sql.Result, error) {
	name := offer.Name
	if len(name) == 0 {
		return nil, errors.New("offer name is required")
	}

	bId := offer.BrandId
	if bId == 0 {
		return nil, errors.New("offer brand_id is required")
	}

	sId := offer.StoreId
	if sId == 0 {
		return nil, errors.New("offer store_id is required")
	}

	price := offer.Price
	if price == 0 {
		return nil, errors.New("offer price is required")
	}

	quantity := offer.Quantity
	if quantity == 0 {
		return nil, errors.New("offer quantity is required")
	}

	available := offer.Available
	if available == 0 {
		return nil, errors.New("offer available is required")
	}

	expirationDate := offer.ExpirationDate
	if expirationDate.IsZero() {
		return nil, errors.New("offer expiration_date is required")
	}

	picture := offer.Picture
	if len(picture) == 0 {
		return nil, errors.New("offer picture is required")
	}

	return session.SQL().InsertInto(OfferTable).Columns(
		"name",
		"brand_id",
		"store_id",
		"price",
		"quantity",
		"available",
		"expiration_date",
		"picture",
	).Values(
		name,
		bId,
		sId,
		price,
		quantity,
		available,
		expirationDate,
		picture,
	).Returning("id").Exec()
}

func GetAvailableOffers(
	session db.Session,
	page uint,
	limit uint,
) (offers []models.Offer, err error) {
	err = session.SQL().SelectFrom(OfferTable).
		Where("expiration_date > NOW()").
		And("available > 0").
		OrderBy("expiration_date DESC").
		Paginate(limit).
		Page(page).
		All(&offers)

	return offers, err
}
