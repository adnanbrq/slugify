package handler

import (
	"github.com/adnanbrq/slugify/internal/dto"
	"github.com/adnanbrq/slugify/internal/entity"
	"github.com/adnanbrq/slugify/internal/errs"
	"github.com/adnanbrq/slugify/internal/repo"
	"github.com/adnanbrq/validation"
	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct{}

func (LinkHandler) HandleCreateLink(c *fiber.Ctx) error {
	input := new(dto.CreateLinkDTO)
	if e := c.BodyParser(input); e != nil {
		return errs.ErrInvalidInput
	}

	if err := validation.Validate(*input); len(err) > 0 {
		return c.Status(403).JSON(fiber.Map{"errors": err})
	}

	linkRepo := repo.LinkRepo{}
	if exists := linkRepo.LinkExists(input.Slug); exists != false {
		return errs.ErrSlugTaken
	}

	link := entity.Link{
		URL:  input.Url,
		Slug: input.Slug,
	}

	if err := linkRepo.Create(&link); err != nil {
		return err
	}

	return c.JSON(link)
}

func (LinkHandler) FollowLink(c *fiber.Ctx) error {
	linkRepo := repo.LinkRepo{}
	slug := c.Params("slug", "-")
	if slug == "-" {
		return errs.ErrSlugTaken
	}

	link, err := linkRepo.GetBySlug(slug)
	if err != nil {
		return err
	}

	return c.Redirect(link.URL)
}
