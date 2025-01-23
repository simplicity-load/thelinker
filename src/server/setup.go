package main

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"exxo.com/database"
	"github.com/jmoiron/sqlx"
	"github.com/lmittmann/tint"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
	// psql_endpoint
	// psql_username
	// psql_password
)

func setupAndConnectDB() (*sqlx.DB, error) {
	flag.Parse()

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return db, err
}
