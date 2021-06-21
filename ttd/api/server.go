package api

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zackartz/ttd/ent"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	client *ent.Client
	Router *fiber.App
}

func (s *Server) Initialize() {
	var err error
	s.client, err = ent.Open("sqlite3", "file:dev.db?cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	c = context.Background()

	if err := s.client.Schema.Create(c); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s.initializeRoutes()

	log.Fatal(s.Router.Listen(":6969"))
}
