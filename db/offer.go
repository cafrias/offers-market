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
	Name               string `db:"product_name"`
	ProductDescription string `db:"product_description"`
	Picture            string `db:"product_picture"`
	StoreName          string `db:"store_name"`
	BrandName          string `db:"brand_name"`
}

const OfferTable = "offer"

func CreateOffer(session db.Session, offer *models.Offer) (sql.Result, error) {
	sId := offer.StoreId
	if sId == 0 {
		return nil, errors.New("offer store_id is required")
	}

	pId := offer.ProductId
	if pId == 0 {
		return nil, errors.New("offer product_id is required")
	}

	price := offer.Price
	if price == 0 {
		return nil, errors.New("offer price is required")
	}

	available := offer.Available
	if available == 0 {
		return nil, errors.New("offer available is required")
	}

	quantity := offer.Quantity
	if quantity == 0 {
		return nil, errors.New("offer quantity is required")
	}

	expirationDate := offer.ExpirationDate
	if expirationDate.IsZero() {
		return nil, errors.New("offer expiration_date is required")
	}

	return session.SQL().InsertInto(OfferTable).Columns(
		"store_id",
		"product_id",
		"price",
		"quantity",
		"available",
		"expiration_date",
		"description",
	).Values(
		sId,
		pId,
		price,
		quantity,
		available,
		expirationDate,
		offer.Description,
	).Returning("id").Exec()
}

func GetAvailableOffers(
	session db.Session,
	page uint,
	limit uint,
) (offers []OfferResult, totalPages uint, err error) {
	// TODO: can we count from a query that doesn't have a left join?
	qr := selectOffers(session).
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

	qr := selectOffers(session).
		Where("offer.expiration_date > NOW()").
		Where("offer.available > 0").
		Where("product.name ILIKE ? OR st.name ILIKE ? OR br.name ILIKE ?", term, term, term).
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

func selectOffers(session db.Session) db.Selector {
	return session.SQL().Select(
		"offer.*",
		"product.name AS product_name",
		"product.description AS product_description",
		"product.picture AS product_picture",
		"st.name AS store_name",
		"br.name AS brand_name",
	).From(OfferTable).
		LeftJoin("product").
		On("offer.product_id = product.id").
		LeftJoin("store AS st").
		On("offer.store_id = st.id").
		LeftJoin("brand AS br").
		On("product.brand_id = br.id")
}
