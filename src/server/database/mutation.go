package database

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func InsertShortlink(db *sqlx.DB, shortlink Shortlink) error {
	if _, err := db.NamedExec(`
		INSERT INTO shortlinks
		(original_url, short_url, date)
		VALUES (:original_url, :short_url, now())`,
		shortlink,
	); err != nil {
		slog.Error("InsertShortlink", "error", err, "shortlink", shortlink)
		return err
	}
	return nil
}
