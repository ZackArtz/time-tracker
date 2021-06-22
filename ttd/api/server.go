package api

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zackartz/ttd/prisma/db"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	client *db.PrismaClient
	Router *fiber.App
}

func (s *Server) Initialize() {
	s.client = db.NewClient()
	if err := s.client.Prisma.Connect(); err != nil {
		log.Panicln(err)
		return
	}

	defer func() {
		if err := s.client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	c = context.Background()

	if err := s.client.Schema.Create(c); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s.initializeRoutes()

	log.Fatal(s.Router.Listen(":6969"))
}
