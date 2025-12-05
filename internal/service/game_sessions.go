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
