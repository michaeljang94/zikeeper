package service

import (
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type PlayerSessionsService struct {
	Repo              *repo.PlayerSessionsRepo
	UserRepo          *repo.UserRepo
	TableSessionsRepo *repo.TableSessionsRepo
}

type AddPlayerToPlayerSessionRequest struct {
	SessionId string
	TableName string
	Player    Player `json:"player"`
}

type AddPlayerToPlayerSessionResponse struct {
}

type GetPlayersForSessionIdRequest struct {
	SessionId string `json:"session_id"`
}

type GetPlayersForSessionIdResponse struct {
	Players []Player `json:"players"`
}

type Player struct {
	Name string `json:"name"`
}

type GetPlayerSessionByUsernameRequest struct {
	Username string
}

type GetPlayerSessionByUsernameResponse struct {
	PlayerSession PlayerSessionObject `json:"player_session"`
	TableSession  TableSession        `json:"table_session"`
}

type PlayerSessionObject struct {
	SessionId string `json:"session_id"`
	TableName string `json:"table_name"`
	Username  string `json:"username"`
}

func (service *PlayerSessionsService) GetPlayerSessionByUsername(request GetPlayerSessionByUsernameRequest) (GetPlayerSessionByUsernameResponse, error) {
	req := repo.GetPlayerSessionByUsernameRequest{
		Username: request.Username,
	}

	res, err := service.Repo.GetPlayerSessionByUsername(req)

	if err != nil {
		return GetPlayerSessionByUsernameResponse{}, err
	}

	// Get table session info
	tableSessionReq := repo.GetTableSessionBySessionIdRequest{
		SessionId: res.PlayerSession.SessionId,
	}

	tableSessionRes, err := service.TableSessionsRepo.GetTableSessionBySessionId(tableSessionReq)

	if err != nil {
		return GetPlayerSessionByUsernameResponse{}, err
	}

	return GetPlayerSessionByUsernameResponse{
		PlayerSession: PlayerSessionObject{
			SessionId: res.PlayerSession.SessionId,
			TableName: res.PlayerSession.TableName,
			Username:  res.PlayerSession.Username,
		},
		TableSession: TableSession{
			Dealer: tableSessionRes.TableSession.Dealer.String,
		},
	}, nil
}

func (service *PlayerSessionsService) GetPlayersForSessionId(request GetPlayersForSessionIdRequest) (GetPlayersForSessionIdResponse, error) {
	req := repo.GetPlayersForSessionIdRequest{
		SessionId: request.SessionId,
	}

	res, err := service.Repo.GetPlayersForSessionId(req)

	if err != nil {
		return GetPlayersForSessionIdResponse{}, err
	}

	var players []Player

	for i := range res.Players {
		player := Player{
			Name: res.Players[i].Name,
		}

		players = append(players, player)
	}

	return GetPlayersForSessionIdResponse{
		Players: players,
	}, nil
}

func (service *PlayerSessionsService) AddPlayerToPlayerSession(request AddPlayerToPlayerSessionRequest) (AddPlayerToPlayerSessionResponse, error) {
	// Check that the session exists...

	// TODO: Check user exists
	userRequest := repo.GetUserRequest{
		UserName: request.Player.Name,
	}

	_, userErr := service.UserRepo.GetUserByUserName(userRequest)

	if userErr != nil {
		return AddPlayerToPlayerSessionResponse{}, userErr
	}

	// TODO: Check to ensure user is not already in the table

	req := repo.AddPlayerToPlayerSessionRequest{
		SessionId: request.SessionId,
		TableName: request.TableName,
		Player: repo.Player{
			Name: request.Player.Name,
		},
	}

	_, err := service.Repo.AddPlayerToPlayerSession(req)

	if err != nil {
		return AddPlayerToPlayerSessionResponse{}, err
	}

	return AddPlayerToPlayerSessionResponse{}, nil
}

type DeletePlayerFromPlayerSessionRequest struct {
	SessionId string
	TableName string
	Username  string `json:"username"`
}

type DeletePlayerFromPlayerSessionResponse struct {
}

func (service *PlayerSessionsService) DeletePlayerFromPlayerSession(request DeletePlayerFromPlayerSessionRequest) (DeletePlayerFromPlayerSessionResponse, error) {
	req := repo.DeletePlayerFromPlayerSessionRequest{
		Username: request.Username,
	}

	_, err := service.Repo.DeletePlayerFromPlayerSession(req)

	if err != nil {
		return DeletePlayerFromPlayerSessionResponse{}, err
	}

	return DeletePlayerFromPlayerSessionResponse{}, nil
}
