package repo

import (
	"database/sql"
	"fmt"
)

type TableSessionsRepo struct {
	Db *sql.DB
}

type TableSession struct {
	SessionId string
	TableName string
	Dealer    sql.NullString
	Status    string
	Pool      string
}

type CreateTableSessionRequest struct {
	SessionId string
	TableName string
}

type CreateTableSessionResponse struct {
}

type GetTableSessionsRequest struct {
	TableName string
}

type GetTableSessionsResponse struct {
	TableSessions []TableSession
}

type DeleteTableSessionsByTableNameRequest struct {
	TableName string
}

type DeleteTableSessionsByTableNameResponse struct {
}

type DeleteTableSessionBySessionIdRequest struct {
	SessionId string
}

type DeleteTableSessionBySessionIdResponse struct {
}

type GetTableSessionBySessionIdRequest struct {
	SessionId string
}

type GetTableSessionBySessionIdResponse struct {
	TableSession TableSession
}

type AddDealerToTableSessionRequest struct {
	SessionId string
	Dealer    string
}

type AddDealerToTableSessionResponse struct {
}

type RemoveDealerFromTableSessionRequest struct {
	SessionId string
}

type RemoveDealerFromTableSessionResponse struct {
}

func (repo *TableSessionsRepo) RemoveDealerFromTableSession(request RemoveDealerFromTableSessionRequest) (RemoveDealerFromTableSessionResponse, error) {
	_, err := repo.Db.Exec("UPDATE table_sessions SET dealer = ? WHERE session_id = ?", sql.NullString{}, request.SessionId)

	if err != nil {
		return RemoveDealerFromTableSessionResponse{}, err
	}

	return RemoveDealerFromTableSessionResponse{}, nil
}

func (repo *TableSessionsRepo) AddDealerToTableSession(request AddDealerToTableSessionRequest) (AddDealerToTableSessionResponse, error) {
	_, err := repo.Db.Exec("UPDATE table_sessions SET dealer = ? WHERE session_id = ?", request.Dealer, request.SessionId)

	if err != nil {
		return AddDealerToTableSessionResponse{}, err
	}

	return AddDealerToTableSessionResponse{}, nil
}

func (repo *TableSessionsRepo) GetTableSessionBySessionId(request GetTableSessionBySessionIdRequest) (GetTableSessionBySessionIdResponse, error) {
	row := repo.Db.QueryRow("SELECT session_id, table_name, dealer, status, money_pool FROM table_sessions WHERE session_id = ?", request.SessionId)

	tableSession := TableSession{}
	if err := row.Scan(&tableSession.SessionId, &tableSession.TableName, &tableSession.Dealer, &tableSession.Status, &tableSession.Pool); err != nil {
		if err == sql.ErrNoRows {
			return GetTableSessionBySessionIdResponse{}, err
		}
	}

	return GetTableSessionBySessionIdResponse{
		TableSession: tableSession,
	}, nil
}

func (repo *TableSessionsRepo) GetTableSessions(request GetTableSessionsRequest) (GetTableSessionsResponse, error) {
	rows, err := repo.Db.Query("SELECT session_id, table_name, dealer, status, money_pool FROM table_sessions WHERE table_name = ?", request.TableName)

	if err != nil {
		return GetTableSessionsResponse{}, err
	}

	defer rows.Close()

	var tableSessions []TableSession
	for rows.Next() {
		var tableSession TableSession

		if err := rows.Scan(&tableSession.SessionId, &tableSession.TableName, &tableSession.Dealer, &tableSession.Status, &tableSession.Pool); err != nil {
			fmt.Println(err)
			return GetTableSessionsResponse{}, err
		}

		tableSessions = append(tableSessions, tableSession)
	}

	return GetTableSessionsResponse{
		TableSessions: tableSessions,
	}, nil
}

func (repo *TableSessionsRepo) CreateTableSession(request CreateTableSessionRequest) (CreateTableResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO table_sessions (session_id, table_name) VALUES (?, ?)", request.SessionId, request.TableName)

	if err != nil {
		return CreateTableResponse{}, err
	}

	return CreateTableResponse{}, nil
}

func (repo *TableSessionsRepo) DeleteTableSessionsByTableName(request DeleteTableSessionsByTableNameRequest) (DeleteTableSessionsByTableNameResponse, error) {
	_, err := repo.Db.Exec("DELETE FROM table_sessions WHERE table_name = ?", request.TableName)

	if err != nil {
		return DeleteTableSessionsByTableNameResponse{}, err
	}

	return DeleteTableSessionsByTableNameResponse{}, nil
}

func (repo *TableSessionsRepo) DeleteTableSessionBySessionId(request DeleteTableSessionBySessionIdRequest) (DeleteTableSessionBySessionIdResponse, error) {
	_, err := repo.Db.Exec("DELETE FROM table_sessions WHERE session_id = ?", request.SessionId)

	if err != nil {
		return DeleteTableSessionBySessionIdResponse{}, err
	}

	return DeleteTableSessionBySessionIdResponse{}, nil
}

type GetTableSessionByDealerRequest struct {
	Dealer string
}

type GetTableSessionByDealerResponse struct {
	TableSession TableSession
}

func (repo *TableSessionsRepo) GetTableSessionByDealer(request GetTableSessionByDealerRequest) (GetTableSessionByDealerResponse, error) {
	row := repo.Db.QueryRow("SELECT session_id, table_name, dealer, status, money_pool FROM table_sessions WHERE dealer = ?", request.Dealer)

	tableSession := TableSession{}
	if err := row.Scan(&tableSession.SessionId, &tableSession.TableName, &tableSession.Dealer, &tableSession.Status, &tableSession.Pool); err != nil {
		if err == sql.ErrNoRows {
			return GetTableSessionByDealerResponse{}, err
		}

		return GetTableSessionByDealerResponse{}, err
	}

	return GetTableSessionByDealerResponse{
		TableSession: tableSession,
	}, nil
}

type UpdateTableSessionStatusBySessionIdRequest struct {
	SessionId string
	Status    string
}

type UpdateTableSessionStatusBySessionIdResponse struct {
}

func (repo *TableSessionsRepo) UpdateTableSessionStatusBySessionId(request UpdateTableSessionStatusBySessionIdRequest) (UpdateTableSessionStatusBySessionIdResponse, error) {
	_, err := repo.Db.Exec("UPDATE table_sessions SET status = ? WHERE session_id = ?", request.Status, request.SessionId)

	if err != nil {
		return UpdateTableSessionStatusBySessionIdResponse{}, err
	}

	return UpdateTableSessionStatusBySessionIdResponse{}, nil
}
