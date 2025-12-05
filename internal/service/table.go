package service

import (
	"strings"

	"github.com/michaeljang94/zikeeper/internal/repo"
)

type TableService struct {
	TableRepo *repo.TableRepo
}

type Player struct {
	Name string `json:"name"`
}

type Table struct {
	Name    string   `json:"name"`
	Game    string   `json:"game"`
	Players []string `json:"players"`
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
	Id        string `json:"id"`
	TableName string `json:"name"`
}

type CreateTableResponse struct {
	Table Table `json:"table"`
}

type AddPlayerToTableRequest struct {
	TableName string `json:"table_name"`
	Player    Player `json:"player"`
}

type AddPlayerToTableResponse struct {
}

func (service *TableService) AddPlayerToTable(request AddPlayerToTableRequest) (AddPlayerToTableResponse, error) {
	// TODO: Check to ensure user is not already in the table

	req := repo.AddPlayerToTableRequest{
		TableName: request.TableName,
		Player: repo.Player{
			Name: request.Player.Name,
		},
	}

	_, err := service.TableRepo.AddPlayerToTable(req)

	if err != nil {
		return AddPlayerToTableResponse{}, err
	}

	return AddPlayerToTableResponse{}, nil
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

	players := strings.Split(response.Table.Players, ",")

	return GetTableByNameResponse{
		Table: Table{
			Name:    response.Table.Name,
			Game:    response.Table.Game,
			Players: players,
		},
	}, nil
}
