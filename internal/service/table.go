package service

import (
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type TableService struct {
	TableRepo *repo.TableRepo
}

type Table struct {
	Name string
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

type CreateTableRequest struct {
	Id        string `json:"id"`
	TableName string `json:"name"`
}

type CreateTableResponse struct {
	Table Table
}

func (service *TableService) CreateTable(request CreateTableRequest) (CreateTableResponse, error) {
	serviceRequest := repo.CreateTableRequest{
		Id:        request.Id,
		TableName: request.TableName,
	}

	response, err := service.TableRepo.CreateTable(serviceRequest)

	if err != nil {
		return CreateTableResponse{}, nil
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
			Name: response.TableName,
		},
	}, nil
}
