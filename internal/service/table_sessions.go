package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type TableSessionsService struct {
	Repo     *repo.TableSessionsRepo
	UserRepo *repo.UserRepo
}

type CreateTableSessionRequest struct {
	TableName string `json:"table_name"`
}

type CreateTableSessionResponse struct {
	SessionId string `json:"session_id"`
}

type GetTableSessionsRequest struct {
	TableName string `json:"table_name"`
}

type GetTableSessionsResponse struct {
	TableSessions []TableSession `json:"table_sessions"`
}

type TableSession struct {
	SessionId string `json:"session_id"`
	TableName string `json:"table_name"`
	Dealer    string `json:"dealer"`
	Status    string `json:"status"`
}

type GetTableSessionBySessionIdRequest struct {
	SessionId string
}

type GetTableSessionBySessionIdResponse struct {
	TableSession TableSession `json:"table_session"`
}

type AddDealerToTableSessionRequest struct {
	SessionId string
	Dealer    string `json:"dealer"`
}

type AddDealerToTableSessionResponse struct {
}

type RemoveDealerFromTableSessionRequest struct {
	SessionId string
}

type RemoveDealerFromTableSessionResponse struct {
}

func (service *TableSessionsService) RemoveDealerFromTableSession(request RemoveDealerFromTableSessionRequest) (RemoveDealerFromTableSessionResponse, error) {
	req := repo.RemoveDealerFromTableSessionRequest{
		SessionId: request.SessionId,
	}

	_, err := service.Repo.RemoveDealerFromTableSession(req)

	if err != nil {
		return RemoveDealerFromTableSessionResponse{}, err
	}

	return RemoveDealerFromTableSessionResponse{}, nil
}

func (service *TableSessionsService) AddDealerToTableSession(request AddDealerToTableSessionRequest) (AddDealerToTableSessionResponse, error) {
	userReq := repo.GetUserRequest{
		UserName: request.Dealer,
	}

	userRes, userErr := service.UserRepo.GetUserByUserName(userReq)

	if userErr != nil {
		return AddDealerToTableSessionResponse{}, userErr
	}

	if userRes.User.Role != "dealer" {
		return AddDealerToTableSessionResponse{}, errors.New("User is not a dealer role")
	}

	req := repo.AddDealerToTableSessionRequest{
		SessionId: request.SessionId,
		Dealer:    request.Dealer,
	}

	_, err := service.Repo.AddDealerToTableSession(req)

	if err != nil {
		return AddDealerToTableSessionResponse{}, err
	}

	return AddDealerToTableSessionResponse{}, nil
}

func (service *TableSessionsService) GetTableSessions(request GetTableSessionsRequest) (GetTableSessionsResponse, error) {
	req := repo.GetTableSessionsRequest{
		TableName: request.TableName,
	}

	res, err := service.Repo.GetTableSessions(req)

	if err != nil {
		return GetTableSessionsResponse{}, err
	}

	var tableSessions []TableSession
	for i := range res.TableSessions {
		tableSession := TableSession{
			SessionId: res.TableSessions[i].SessionId,
			TableName: res.TableSessions[i].TableName,
			Dealer:    res.TableSessions[i].Dealer.String,
			Status:    res.TableSessions[i].Status,
		}

		tableSessions = append(tableSessions, tableSession)
	}

	return GetTableSessionsResponse{
		TableSessions: tableSessions,
	}, nil
}

type DeleteTableSessionsByTableNameRequest struct {
	TableName string `json:"table_name"`
}

type DeleteTableSessionsByTableNameResponse struct {
	Status string `json:"status"`
}

type DeleteTableSessionBySessionIdRequest struct {
	SessionId string `json:"session_id"`
	TableName string
}

type DeleteTableSessionBySessionIdResponse struct {
}

func (service *TableSessionsService) DeleteTableSessionBySessionId(request DeleteTableSessionBySessionIdRequest) (DeleteTableSessionBySessionIdResponse, error) {
	// TODO Check that session belongs to the table.. Otherwise do not delete

	playerSessionService := PlayerSessionsService{
		Repo: &repo.PlayerSessionsRepo{
			Db: service.Repo.Db,
		},
	}

	r := GetPlayersForSessionIdRequest{
		SessionId: request.SessionId,
	}

	res, err := playerSessionService.GetPlayersForSessionId(r)
	if err != nil {
		return DeleteTableSessionBySessionIdResponse{}, err
	}

	if len(res.Players) > 0 {
		return DeleteTableSessionBySessionIdResponse{}, errors.New("Session still active: " + request.SessionId)
	}

	req := repo.DeleteTableSessionBySessionIdRequest{
		SessionId: request.SessionId,
	}

	_, err = service.Repo.DeleteTableSessionBySessionId(req)

	if err != nil {
		return DeleteTableSessionBySessionIdResponse{}, err
	}

	return DeleteTableSessionBySessionIdResponse{}, nil
}

func (service *TableSessionsService) DeleteTableSessionsByTableName(request DeleteTableSessionsByTableNameRequest) (DeleteTableSessionsByTableNameResponse, error) {
	getTableSessionRequest := GetTableSessionsRequest{
		TableName: request.TableName,
	}

	getTableSessionResponse, err := service.GetTableSessions(getTableSessionRequest)

	if err != nil {
		return DeleteTableSessionsByTableNameResponse{}, err
	}

	playerSessionService := PlayerSessionsService{
		Repo: &repo.PlayerSessionsRepo{
			Db: service.Repo.Db,
		},
	}

	for _, tableSession := range getTableSessionResponse.TableSessions {
		r := GetPlayersForSessionIdRequest{
			SessionId: tableSession.SessionId,
		}

		res, err := playerSessionService.GetPlayersForSessionId(r)
		if err != nil {
			return DeleteTableSessionsByTableNameResponse{}, err
		}

		if len(res.Players) > 0 {
			return DeleteTableSessionsByTableNameResponse{}, errors.New("Session still active: " + tableSession.SessionId)
		}
	}

	deleteRepoRequest := repo.DeleteTableSessionsByTableNameRequest{
		TableName: request.TableName,
	}

	_, err = service.Repo.DeleteTableSessionsByTableName(deleteRepoRequest)

	if err != nil {
		return DeleteTableSessionsByTableNameResponse{
			Status: "error",
		}, err
	}

	return DeleteTableSessionsByTableNameResponse{
		Status: "OK",
	}, nil
}

func (service *TableSessionsService) CreateTableSession(request CreateTableSessionRequest) (CreateTableSessionResponse, error) {
	// Check table exists
	tableService := TableService{
		TableRepo: &repo.TableRepo{
			Db: service.Repo.Db,
		},
	}

	getTableRequest := GetTableByNameRequest{
		TableName: request.TableName,
	}

	_, err := tableService.GetTableByName(getTableRequest)

	if err != nil {
		return CreateTableSessionResponse{}, err
	}

	// Create
	sessionId := uuid.New()

	req := repo.CreateTableSessionRequest{
		SessionId: sessionId.String(),
		TableName: request.TableName,
	}

	_, err = service.Repo.CreateTableSession(req)

	if err != nil {
		return CreateTableSessionResponse{}, err
	}

	return CreateTableSessionResponse{
		SessionId: sessionId.String(),
	}, nil
}
