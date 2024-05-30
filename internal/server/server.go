package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gustafer/go-games-blueprint/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "Tumadre2",
			AppName:      "Tumadre2",
		}),

		db: database.New(),
	}

	return server
}
