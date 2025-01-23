package database

import (
	_ "embed"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:embed schema.sql
var schema string

func Connect() (*sqlx.DB, error) {
	psqlendpoint := "user=socialplus dbname=socialplus password=socialplus123 host=localhost sslmode=disable"

	var db *sqlx.DB
	db, err := sqlx.Open("postgres", psqlendpoint)
	if err != nil {
		slog.Error("Connect:Failed to connect to db", "error", err, "endpoint", psqlendpoint)
		return nil, err
	}

	slog.Debug("Connect", "schema", schema)

	_, err = db.Exec(schema)
	if err != nil {
		slog.Error("Connect:Couldn't execute schema", "error", err, "schema", schema)
		return nil, err
	}

	return db, nil
}
