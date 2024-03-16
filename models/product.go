package models

type Product struct {
	Id          uint   `db:"id"`
	Name        string `db:"name"`
	Picture     string `db:"picture"`
	BrandId     uint   `db:"brand_id"`
	Description string `db:"description"`
}
