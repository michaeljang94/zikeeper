package repo

import (
	"database/sql"
	"fmt"
)

type GameSessionsRepo struct {
	Db *sql.DB
}

type AddPlayerToGameSessionRequest struct {
	SessionId string
	TableName string
	Player    Player
}

type AddPlayerToGameSessionResponse struct {
}

func (repo *GameSessionsRepo) AddPlayerToGameSession(request AddPlayerToGameSessionRequest) (AddPlayerToGameSessionResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO game_sessions (session_id, table_name, username) VALUES (?, ?, ?)", request.SessionId, request.TableName, request.Player.Name)

	if err != nil {
		fmt.Println(err)
		return AddPlayerToGameSessionResponse{}, err
	}

	return AddPlayerToGameSessionResponse{}, nil
}
