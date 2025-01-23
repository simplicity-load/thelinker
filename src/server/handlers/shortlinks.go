package handlers

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"exxo.com/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const ENDPOINT = "http://localhost:8800"

func Shortlinks(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortlinks, err := database.QueryShortlinks(db)
		if err != nil {
			return SendError(c, fiber.ErrInternalServerError.Message)
		}
		shortlinks, err = database.QueryShortlinks(db)
		if err != nil {
			return SendError(c, fiber.ErrInternalServerError.Message)
		}

		for i, link := range shortlinks {
			shortlinks[i] = database.Shortlink{
				OriginalURL: link.OriginalURL,
				ShortURL:    fmt.Sprintf("%s/%s", ENDPOINT, link.ShortURL),
				Date:        link.Date,
			}
		}
		return SendMessage(c, shortlinks)
	}
}

func genRandomStr() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var sb strings.Builder
	for range 6 {
		i := rand.Intn(len(charset))
		sb.WriteByte(charset[i])
	}
	return sb.String()
}

const MAX_RETRIES = 5

func SubmitLink(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userInput := SubmitLinkInput{}
		if err := ParseData(c, &userInput); err != nil {
			return SendError(c, fiber.ErrBadRequest.Message)
		}

		var genStr string
		for range MAX_RETRIES {
			genStr = genRandomStr()
			_, err := database.GetShortlink(db, genStr)
			if err != nil {
				// Link not found so genStr can be used as new link
				break
			}
		}
		slog.Debug("genshortlink", "link", genRandomStr())

		shortlink := database.Shortlink{
			OriginalURL: userInput.OriginalURL,
			ShortURL:    genStr,
			Date:        time.Now(),
		}

		if err := database.InsertShortlink(db, shortlink); err != nil {
			return SendError(c, fiber.ErrInternalServerError.Message)
		}

		return SendMessage(c, shortlink)
	}
}

func Redirect(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		link := c.Params("shortlink")

		shortlink, err := database.GetShortlink(db, link)
		if err != nil {
			return SendError(c, "Short link doesn't exist")
		}
		return c.Redirect(shortlink.OriginalURL, fiber.StatusMovedPermanently)
	}

}
