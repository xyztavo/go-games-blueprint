package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gustafer/go-games-blueprint/internal/database"
	"github.com/gustafer/go-games-blueprint/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.New()

	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if err := database.New().AutoMigrate(); err != nil {
		panic(err)
	}
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
