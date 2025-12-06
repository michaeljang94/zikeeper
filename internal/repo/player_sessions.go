package repo

import (
	"database/sql"
	"fmt"
)

type PlayerSessionsRepo struct {
	Db *sql.DB
}

type AddPlayerToPlayerSessionRequest struct {
	SessionId string
	TableName string
	Player    Player
}

type AddPlayerToPlayerSessionResponse struct {
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

func (repo *PlayerSessionsRepo) GetPlayersForSessionId(request GetPlayersForSessionIdRequest) (GetPlayersForSessionIdResponse, error) {
	rows, err := repo.Db.Query("SELECT username FROM player_sessions WHERE session_id = ?", request.SessionId)

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

func (repo *PlayerSessionsRepo) AddPlayerToPlayerSession(request AddPlayerToPlayerSessionRequest) (AddPlayerToPlayerSessionResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO player_sessions (session_id, table_name, username) VALUES (?, ?, ?)", request.SessionId, request.TableName, request.Player.Name)

	if err != nil {
		fmt.Println(err)
		return AddPlayerToPlayerSessionResponse{}, err
	}

	return AddPlayerToPlayerSessionResponse{}, nil
}
