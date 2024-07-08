package utils

import (
	"database/sql"
	"webpaygo/api/config"

	_ "github.com/lib/pq"
)

// OpenDB abre la conexión con la base de datos utilizando la URL generada por config.DBURL()
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBURL())
	if err != nil {
		return nil, err
	}
	return db, nil
}
