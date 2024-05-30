package handlers

import (
	"fmt"
	"net/http"

	"github.com/gustafer/go-games-blueprint/internal/database"

	"github.com/gofiber/fiber/v2"
)

func HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
		"desc":    "this project uses go blueprint with fiber and postgres.",
	}

	return c.JSON(resp)
}

func HealthHandler(c *fiber.Ctx) error {
	return c.JSON(database.New().Health())
}

func CreateGameHandler(c *fiber.Ctx) error {
	// instanciate new variable with game type
	game := new(database.Game)
	// bind the body with the game type
	if err := c.BodyParser(game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "malformed req body"})
	}
	// Make the query to the database
	createdGameId, err := database.NewQuery().CreateGame(game)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create game"})
	}

	return c.JSON(fiber.Map{"createdGameId": createdGameId})
}

func GetGameByIdHandler(c *fiber.Ctx) error {
	// get id from params
	id := c.Params("id")
	// make the DB query using the ID from params
	game, err := database.NewQuery().GetGameById(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("could not find game with id %s", id)})
	}

	return c.JSON(fiber.Map{"game": game})
}

func GetGames(c *fiber.Ctx) error {
	// make the DB query
	games, err := database.NewQuery().GetGames()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not get games."})
	}

	return c.JSON(fiber.Map{"game": games})
}

func UpdateGame(c *fiber.Ctx) error {
	// get param and instanciate a new game var with type database.Game
	id := c.Params("id")
	game := new(database.Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "malformed req body"})
	}
	// make the DB query
	rowsAffected, err := database.NewQuery().UpdateGame(id, game)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "something happened on server side"})
	}
	// check if it affected any rows
	if rowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("cannot find game with id %s", id), "rowsAffected": rowsAffected})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "updated successfully", "rowsAffected": rowsAffected})
}

func DeleteGame(c *fiber.Ctx) error {
	// get the id by params
	id := c.Params("id")
	// make the DB query using id by params
	rowsAffected, err := database.NewQuery().DeleteGame(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "something happened on server side"})
	}
	// check if it affected any rows
	if rowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("cannot find game by id: %s", id)})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "game deleted with ease!"})
}
