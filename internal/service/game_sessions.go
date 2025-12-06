package service

import "github.com/michaeljang94/zikeeper/internal/repo"

type GameSessionsService struct {
	Repo *repo.GameSessionsRepo
}

type AddPlayerToGameSessionRequest struct {
	SessionId string
	TableName string `json:"table_name"`
	Player    Player `json:"player"`
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

func (service *GameSessionsService) GetPlayersForSessionId(request GetPlayersForSessionIdRequest) (GetPlayersForSessionIdResponse, error) {
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

func (service *GameSessionsService) AddPlayerToGameSession(request AddPlayerToGameSessionRequest) (AddPlayerToGameSessionResponse, error) {
	// Check that the session exists...

	// TODO: Check user exists

	// TODO: Check to ensure user is not already in the table

	req := repo.AddPlayerToGameSessionRequest{
		SessionId: request.SessionId,
		TableName: request.TableName,
		Player: repo.Player{
			Name: request.Player.Name,
		},
	}

	_, err := service.Repo.AddPlayerToGameSession(req)

	if err != nil {
		return AddPlayerToGameSessionResponse{}, err
	}

	return AddPlayerToGameSessionResponse{}, nil
}
