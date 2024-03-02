package models

type Store struct {
	Id      uint
	Name    string
	Address string
	// location represents the latitude and longitude of the store
	Location [2]float64
	Website  string
	Phone    string
	Email    string
}
