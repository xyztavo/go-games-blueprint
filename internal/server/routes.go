package server

import "github.com/gustafer/go-games-blueprint/internal/handlers"

func (s *FiberServer) RegisterFiberRoutes() {
	// Hello World
	s.App.Get("/", handlers.HelloWorldHandler)
	// db health info
	s.App.Get("/health", handlers.HealthHandler)

	// // Games routes:
	// Create game:
	s.App.Post("/game", handlers.CreateGameHandler)
	// Get Games
	s.App.Get("/games", handlers.GetGames)
	// Get Game by ID
	s.App.Get("/game/:id", handlers.GetGameByIdHandler)
	// Update Game by ID
	s.App.Put("/game/:id", handlers.UpdateGame)
	// Update Game by ID
	s.App.Delete("/game/:id", handlers.DeleteGame)
}
