package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type PutItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) PutItem(c *fiber.Ctx) error {
	var req PutItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}
	err := s.Repository.Update(c.Context(), tracker.Item{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
