package repo

import "database/sql"

type TableRepo struct {
	Db *sql.DB
}

type Table struct {
	Id   string
	Name string
}

type GetTableByNameRequest struct {
	TableName string
}

type GetTableByNameResponse struct {
	TableName string
}

func (repo *TableRepo) GetTableByName(request GetTableByNameRequest) (GetTableByNameResponse, error) {
	row := repo.Db.QueryRow("SELECT id, name FROM tables WHERE name = ?", request.TableName)

	table := Table{}
	if err := row.Scan(&table.Id, &table.Name); err != nil {
		if err == sql.ErrNoRows {
			return GetTableByNameResponse{}, err
		}
	}

	return GetTableByNameResponse{
		TableName: table.Name,
	}, nil
}
