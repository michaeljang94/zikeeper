package service

import "github.com/michaeljang94/zikeeper/internal/repo"

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
