package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	err := s.Repository.Delete(c.Context(), id)
	if err != nil {
		log.Errorw("s.Repository.Delete", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
	//return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}
