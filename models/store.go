package models

type Store struct {
	Id      uint    `db:"id"`
	Name    string  `db:"name"`
	Address string  `db:"address"`
	Lat     float64 `db:"lat"`
	Lng     float64 `db:"lng"`
	Website string  `db:"website"`
	Phone   string  `db:"phone"`
	Email   string  `db:"email"`
}
