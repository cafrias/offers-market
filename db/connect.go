package db

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{
	// TODO: move settings to secrets file
	Database: `postgres`,
	Host:     `localhost`,
	User:     `postgres`,
	Password: `example`,
}

func Connect() (db.Session, error) {
	return postgresql.Open(settings)
}
