package http

import (
	"fmt"
	"os"

	"github.com/adnanbrq/slugify/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type server struct {
	app *fiber.App
}

func NewRestServer() *server {
	return &server{
		app: fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				msg := "Interval Error"

				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
					msg = e.Message
				}

				return c.Status(code).JSON(fiber.Map{"error": msg})
			},
		}),
	}
}

func (s *server) Init() {
	linkHandler := handler.LinkHandler{}
	linkGroup := s.app.Group("/link")
	linkGroup.Post("/", linkHandler.HandleCreateLink)
	linkGroup.Get("/:slug", linkHandler.FollowLink)
}

func (s *server) Boot() error {
	return s.app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
