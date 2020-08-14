package api

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func (s *Server) initializeRoutes() {
	s.Router = fiber.New()

	s.Router.Use(middleware.Logger())

	s.Router.Get("/", func(ctx *fiber.Ctx) {
		ctx.JSON(struct {
			Message string `json:"message"`
		}{
			Message: "the ttd daemon is working",
		})
	})

	s.Router.Post("/api/v1/create", s.CreateTimestamp)
	s.Router.Get("/api/v1/end/:uuid", s.EndTimestampByUUID)

	s.Router.Get("/api/v1/timestamps", s.GetAllTimestamps)
	s.Router.Get("/api/v1/active-timestamps", s.GetAllActiveTimestamps)
	s.Router.Get("/api/v1/timestamps/:project", s.GetTimestampsByProject)
	s.Router.Get("/api/v1/timestamps/:uuid", s.GetTimestampByUUID)
	s.Router.Delete("/api/v1/timestamps/:id", s.DeleteTimestamp)
}
