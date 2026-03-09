package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/item/", s.CreateItem)
	route.Get("/items/", s.GetItems)
	route.Delete("/item/:id", s.DeleteItem)
	route.Put("/item/:id", s.PutItem)
	route.Get("/items/search", s.FindByNameItem)
}
