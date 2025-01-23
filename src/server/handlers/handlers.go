package handlers

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// TODO create custom error types

func SendError(c *fiber.Ctx, errMsg string) error {
	return c.Status(500).JSON(Response{
		Err: &errMsg,
	})
}

func SendMessage(c *fiber.Ctx, msg any) error {
	return c.Status(200).JSON(Response{
		Data: msg,
	})
}

func ParseData[T any](c *fiber.Ctx, data *T) error {
	if err := c.BodyParser(data); err != nil {
		slog.Error("ParseData", "error", err)
		return errors.New("Invalid data sent")
	}
	return nil
}
