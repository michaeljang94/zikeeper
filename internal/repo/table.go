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
}

type CreateTableResponse struct {
	Table Table
}

func (repo *TableRepo) CreateTable(request CreateTableRequest) (CreateTableResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO tables (id, name) VALUES (?, ?)", request.Id, request.TableName)

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
