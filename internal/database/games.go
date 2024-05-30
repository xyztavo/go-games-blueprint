package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Service represents a service that interacts with a database.
type QueryService interface {
	// Create a Game
	CreateGame(game *Game) (id int64, err error)
	// Get a Game by ID
	GetGameById(id string) (game Game, err error)
	// Gets all Games
	GetGames() (games []Game, err error)
	// Update a Game by ID
	UpdateGame(id string, game *Game) (rowsAffected int64, err error)
	// Delete game by ID
	DeleteGame(id string) (rowsAffected int64, err error)
}

func NewQuery() QueryService {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) CreateGame(game *Game) (id int64, err error) {
	// insert the data and get id
	err = s.db.QueryRow(`
		INSERT INTO games (name,description) 
		VALUES ($1, $2) RETURNING id
		`, game.Name, game.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) GetGameById(gameId string) (game Game, err error) {
	// insert the data and get id
	err = s.db.QueryRow(`SELECT * FROM games WHERE id = $1;`, gameId).Scan(&game.Id, &game.Name, &game.Description)
	if err != nil {
		return game, err
	}

	return game, nil
}

func (s *service) GetGames() (games []Game, err error) {
	// Select everything from games
	rows, err := s.db.Query(`SELECT * FROM games ORDER BY id`)
	if err != nil {
		return nil, err
	}
	// Map each row
	for rows.Next() {
		// Instanciante a var game
		var game Game
		// Scan returned rows to the var
		if err := rows.Scan(&game.Id, &game.Name, &game.Description); err != nil {
			return nil, err
		}
		// Create a game array and append the current game
		games = append(games, game)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

func (s *service) UpdateGame(id string, game *Game) (updatedRows int64, err error) {
	// Update game name and description
	res, err := s.db.Exec(`UPDATE games SET name = $2, description = $3 WHERE id = $1;`, id, game.Name, game.Description)
	if err != nil {
		return 0, err
	}
	// Get rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s *service) DeleteGame(id string) (rowsAffected int64, err error) {
	res, err := s.db.Exec(`DELETE FROM games WHERE id = $1;`, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
