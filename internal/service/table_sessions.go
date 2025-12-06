package service

import (
	"github.com/google/uuid"
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type TableSessionsService struct {
	Repo *repo.TableSessionsRepo
}

type CreateTableSessionRequest struct {
	TableName string `json:"table_name"`
}

type CreateTableSessionResponse struct {
}

type GetTableSessionsRequest struct {
	TableName string `json:"table_name"`
}

type GetTableSessionsResponse struct {
	TableSessions []TableSession `json:"table_sessions"`
}

type TableSession struct {
	SessionId string `json:"session_id"`
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
		}

		tableSessions = append(tableSessions, tableSession)
	}

	return GetTableSessionsResponse{
		TableSessions: tableSessions,
	}, nil
}

func (service *TableSessionsService) CreateTableSession(request CreateTableSessionRequest) (CreateTableSessionResponse, error) {
	sessionId := uuid.New()

	req := repo.CreateTableSessionRequest{
		SessionId: sessionId.String(),
		TableName: request.TableName,
	}

	_, err := service.Repo.CreateTableSession(req)

	if err != nil {
		return CreateTableSessionResponse{}, err
	}

	return CreateTableSessionResponse{}, nil
}
