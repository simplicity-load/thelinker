package database

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func GetShortlink(db *sqlx.DB, shortlink string) (Shortlink, error) {
	link := Shortlink{}
	err := db.Get(&link, "SELECT * FROM shortlinks WHERE short_url = $1", shortlink)
	if err != nil {
		slog.Error("GetShortlink", "error", err, "shortlink", shortlink)
		return Shortlink{}, err
	}
	return link, nil
}

func QueryShortlinks(db *sqlx.DB) ([]Shortlink, error) {
	links := []Shortlink{}
	err := db.Select(&links, "SELECT * FROM shortlinks;")
	if err != nil {
		slog.Error("QueryShortlinks", "error", err)
		return nil, err
	}
	slog.Debug("QueryShortlinks", "links", links)
	return links, nil
}
