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

type DeletePlayerSessionsByTableNameRequest struct {
	TableName string
}

type DeletePlayerSessionsByTableNameResponse struct {
}

type GetPlayerSessionByUsernameRequest struct {
	Username string
}

type GetPlayerSessionByUsernameResponse struct {
	PlayerSession PlayerSessionObject
}

type PlayerSessionObject struct {
	SessionId string
	TableName string
	Username  string
}

func (repo *PlayerSessionsRepo) GetPlayerSessionByUsername(request GetPlayerSessionByUsernameRequest) (GetPlayerSessionByUsernameResponse, error) {
	row := repo.Db.QueryRow("SELECT session_id, table_name, username FROM player_sessions WHERE username = ?", request.Username)

	playerSession := PlayerSessionObject{}
	if err := row.Scan(&playerSession.SessionId, &playerSession.TableName, &playerSession.Username); err != nil {
		if err == sql.ErrNoRows {
			return GetPlayerSessionByUsernameResponse{}, err
		}
	}

	return GetPlayerSessionByUsernameResponse{
		PlayerSession: playerSession,
	}, nil
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

func (repo *PlayerSessionsRepo) DeletePlayerSessionsByTableName(request DeletePlayerSessionsByTableNameRequest) (DeletePlayerSessionsByTableNameResponse, error) {
	_, err := repo.Db.Exec("DELETE FROM player_sessions WHERE table_name = ?", request.TableName)

	if err != nil {
		fmt.Println(err)
		return DeletePlayerSessionsByTableNameResponse{}, err
	}

	return DeletePlayerSessionsByTableNameResponse{}, nil
}

type DeletePlayerFromPlayerSessionRequest struct {
	SessionId string
	TableName string
	Username  string
}

type DeletePlayerFromPlayerSessionResponse struct {
}

func (repo *PlayerSessionsRepo) DeletePlayerFromPlayerSession(request DeletePlayerFromPlayerSessionRequest) (DeletePlayerFromPlayerSessionResponse, error) {
	_, err := repo.Db.Exec("DELETE FROM player_sessions WHERE username = ?", request.Username)

	if err != nil {
		fmt.Println(err)

		return DeletePlayerFromPlayerSessionResponse{}, err
	}

	return DeletePlayerFromPlayerSessionResponse{}, nil
}
