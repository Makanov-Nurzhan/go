package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresDB(host, port, user, password, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)

	return sql.Open("postgres", dsn)
}
