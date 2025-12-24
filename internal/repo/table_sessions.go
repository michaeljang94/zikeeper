package repo

import (
	"database/sql"
)

type TableSessionsRepo struct {
	Db *sql.DB
}

type TableSession struct {
	SessionId string
	TableName string
	Dealer    string
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

func (repo *TableSessionsRepo) GetTableSessionBySessionId(request GetTableSessionBySessionIdRequest) (GetTableSessionBySessionIdResponse, error) {
	row := repo.Db.QueryRow("SELECT session_id, table_name, dealer FROM table_sessions WHERE session_id = ?", request.SessionId)

	tableSession := TableSession{}
	if err := row.Scan(&tableSession.SessionId, &tableSession.TableName, &tableSession.Dealer); err != nil {
		if err == sql.ErrNoRows {
			return GetTableSessionBySessionIdResponse{}, err
		}
	}

	return GetTableSessionBySessionIdResponse{
		TableSession: tableSession,
	}, nil
}

func (repo *TableSessionsRepo) GetTableSessions(request GetTableSessionsRequest) (GetTableSessionsResponse, error) {
	rows, err := repo.Db.Query("SELECT session_id FROM table_sessions WHERE table_name = ?", request.TableName)

	if err != nil {
		return GetTableSessionsResponse{}, err
	}

	defer rows.Close()

	var tableSessions []TableSession
	for rows.Next() {
		var tableSession TableSession

		if err := rows.Scan(&tableSession.SessionId); err != nil {
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
