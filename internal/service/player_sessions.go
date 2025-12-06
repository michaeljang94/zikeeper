package service

import "github.com/michaeljang94/zikeeper/internal/repo"

type PlayerSessionsService struct {
	Repo *repo.PlayerSessionsRepo
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
