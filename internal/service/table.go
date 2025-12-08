package service

import (
	"github.com/google/uuid"
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type TableService struct {
	TableRepo *repo.TableRepo
}

type Table struct {
	Name string `json:"name"`
	Game string `json:"game"`
}

type GetTableByNameRequest struct {
	TableName string `json:"name"`
}

type GetTableByNameResponse struct {
	Table Table `json:"table"`
}

type GetTablesRequest struct {
}

type GetTablesResponse struct {
	Tables []Table `json:"tables"`
}

type CreateTableRequest struct {
	TableName string `json:"name"`
}

type CreateTableResponse struct {
	Table Table `json:"table"`
}

func (service *TableService) CreateTable(request CreateTableRequest) (CreateTableResponse, error) {
	id := uuid.New()

	repoRequest := repo.CreateTableRequest{
		Id:        id.String(),
		TableName: request.TableName,
	}

	response, err := service.TableRepo.CreateTable(repoRequest)

	if err != nil {
		return CreateTableResponse{}, err
	}

	return CreateTableResponse{
		Table: Table{
			Name: response.Table.Name,
		},
	}, nil
}

func (service *TableService) GetTables(request GetTablesRequest) (GetTablesResponse, error) {
	repoRequest := repo.GetTablesRequest{}

	repoResponse, err := service.TableRepo.GetTables(repoRequest)

	if err != nil {
		return GetTablesResponse{}, err
	}

	var tables []Table
	for i := range repoResponse.Tables {
		table := Table{
			Name: repoResponse.Tables[i].Name,
		}

		tables = append(tables, table)
	}

	return GetTablesResponse{
		Tables: tables,
	}, nil
}

func (service *TableService) GetTableByName(request GetTableByNameRequest) (GetTableByNameResponse, error) {
	getTableRepoRequest := repo.GetTableByNameRequest{
		TableName: request.TableName,
	}

	response, err := service.TableRepo.GetTableByName(getTableRepoRequest)

	if err != nil {
		return GetTableByNameResponse{}, err
	}

	return GetTableByNameResponse{
		Table: Table{
			Name: response.Table.Name,
			Game: response.Table.Game,
		},
	}, nil
}
