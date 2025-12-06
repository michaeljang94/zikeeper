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

type GetPlayersForSessionIdRequest struct {
	SessionId string
}

type GetPlayersForSessionIdResponse struct {
	Players []Player
}

type Player struct {
	Name string `json:"name"`
}

func (repo *GameSessionsRepo) GetPlayersForSessionId(request GetPlayersForSessionIdRequest) (GetPlayersForSessionIdResponse, error) {

	fmt.Println(request.SessionId)

	rows, err := repo.Db.Query("SELECT username FROM game_sessions WHERE session_id = ?", request.SessionId)

	if err != nil {
		return GetPlayersForSessionIdResponse{}, err
	}

	defer rows.Close()

	var players []Player

	for rows.Next() {
		var player Player

		if err := rows.Scan(&player.Name); err != nil {
			return GetPlayersForSessionIdResponse{}, err
		}

		players = append(players, player)
	}

	return GetPlayersForSessionIdResponse{
		Players: players,
	}, nil
}

func (repo *GameSessionsRepo) AddPlayerToGameSession(request AddPlayerToGameSessionRequest) (AddPlayerToGameSessionResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO game_sessions (session_id, table_name, username) VALUES (?, ?, ?)", request.SessionId, request.TableName, request.Player.Name)

	if err != nil {
		fmt.Println(err)
		return AddPlayerToGameSessionResponse{}, err
	}

	return AddPlayerToGameSessionResponse{}, nil
}
