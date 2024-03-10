package db

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/cafrias/offers-market/models"
	"github.com/upper/db/v4"
)

type OfferResult struct {
	models.Offer
	StoreName string `db:"store_name"`
	BrandName string `db:"brand_name"`
}

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
) (offers []OfferResult, totalPages uint, err error) {
	// TODO: can we count from a query that doesn't have a left join?
	qr := session.SQL().Select(
		"offer.*",
		"st.name AS store_name",
		"br.name AS brand_name",
	).
		From(OfferTable).
		LeftJoin("store AS st").
		On("offer.store_id = st.id").
		LeftJoin("brand AS br").
		On("offer.brand_id = br.id").
		Where("expiration_date > NOW()").
		And("available > 0").
		OrderBy("expiration_date DESC").
		Paginate(limit).
		Page(page)

	totalPages, err = qr.TotalPages()
	if err != nil {
		return nil, 0, err
	}

	err = qr.All(&offers)

	return offers, totalPages, err
}

func SearchAvailableOffers(
	session db.Session,
	term string,
	page uint,
	limit uint,
) (offers []OfferResult, totalPages uint, err error) {
	termBd := strings.Builder{}
	termBd.WriteString("%")
	termBd.WriteString(term)
	termBd.WriteString("%")

	term = termBd.String()

	qr := session.SQL().
		Select(
			"offer.*",
			"st.name AS store_name",
			"br.name as brand_name",
		).
		From(OfferTable).
		LeftJoin("store AS st").
		On("offer.store_id = st.id").
		LeftJoin("brand AS br").
		On("offer.brand_id = br.id").
		Where("offer.expiration_date > NOW()").
		Where("offer.available > 0").
		Where("offer.name ILIKE ? OR st.name ILIKE ? OR br.name ILIKE ?", term, term, term).
		OrderBy("expiration_date ASC").
		Paginate(limit).
		Page(page)

	totalPages, err = qr.TotalPages()
	if err != nil {
		return nil, 0, err
	}

	err = qr.All(&offers)

	return offers, totalPages, err
}
