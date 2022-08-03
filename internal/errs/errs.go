package errs

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrDbNotConnected = fiber.NewError(500, "Server error.")
	ErrLinkNotFound   = fiber.NewError(404, "Link not found.")
	ErrDbError        = fiber.NewError(500, "Server error.")
	ErrSlugTaken      = fiber.NewError(403, "Slug already taken.")
	ErrInvalidInput   = fiber.NewError(403, "Invalid input received.")
)
