package repo

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

type TableRepo struct {
	Db *sql.DB
}

type Table struct {
	Id   string
	Name string
	Game string
}

type GetTableByNameRequest struct {
	TableName string
}

type GetTableByNameResponse struct {
	Table Table
}

type GetTablesRequest struct {
}

type GetTablesResponse struct {
	Tables []Table
}

// TODO auto generate a random ID instead of taking it in
type CreateTableRequest struct {
	Id        string
	TableName string
	Game      string
}

type CreateTableResponse struct {
	Table Table
}

type DeleteTableRequest struct {
	TableName string
}

type DeleteTableResponse struct {
}

func (repo *TableRepo) DeleteTable(request DeleteTableRequest) (DeleteTableResponse, error) {
	_, err := repo.Db.Exec("DELETE FROM tables WHERE name = ?", request.TableName)

	if err != nil {
		return DeleteTableResponse{}, err
	}

	return DeleteTableResponse{}, nil
}

func (repo *TableRepo) CreateTable(request CreateTableRequest) (CreateTableResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO tables (id, name, game) VALUES (?, ?, ?)", request.Id, request.TableName, request.Game)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return CreateTableResponse{}, errors.New("table already exists")
		}

		return CreateTableResponse{}, err
	}

	return CreateTableResponse{
		Table: Table{
			Id:   request.Id,
			Name: request.TableName,
			Game: request.Game,
		},
	}, nil
}

func (repo *TableRepo) GetTables(request GetTablesRequest) (GetTablesResponse, error) {
	rows, err := repo.Db.Query("SELECT id, name FROM tables")

	if err != nil {
		return GetTablesResponse{}, err
	}

	defer rows.Close()

	var tables []Table

	for rows.Next() {
		var table Table
		if err := rows.Scan(&table.Id, &table.Name); err != nil {
			return GetTablesResponse{}, err
		}
		tables = append(tables, table)
	}

	if err = rows.Err(); err != nil {
		return GetTablesResponse{}, err
	}

	return GetTablesResponse{
		Tables: tables,
	}, nil
}

func (repo *TableRepo) GetTableByName(request GetTableByNameRequest) (GetTableByNameResponse, error) {
	row := repo.Db.QueryRow("SELECT id, name, game FROM tables WHERE name = ?", request.TableName)

	table := Table{}
	if err := row.Scan(&table.Id, &table.Name, &table.Game); err != nil {
		if err == sql.ErrNoRows {
			return GetTableByNameResponse{}, err
		}
	}

	return GetTableByNameResponse{
		Table: table,
	}, nil
}
