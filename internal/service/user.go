package service

import (
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	UserName string `json:"username"`
}

type GetUserRequest struct {
	Id       string `json:"id" binding:"required"`
	UserName string `json:"username"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type UserService struct {
	Repo *repo.UserRepo
}

func (service *UserService) GetUserByUserName(request GetUserRequest) (GetUserResponse, error) {
	repoRequest := repo.GetUserRequest{
		UserName: request.UserName,
	}

	response, err := service.Repo.GetUserByUserName(repoRequest)

	if err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		User: User{
			Id:       response.User.Id,
			Name:     response.User.Name,
			Score:    response.User.Score,
			UserName: response.User.UserName,
		},
	}, nil
}

func (service *UserService) GetUser(request GetUserRequest) (GetUserResponse, error) {
	repoRequest := repo.GetUserRequest{
		Id: request.Id,
	}

	response, err := service.Repo.GetUser(repoRequest)

	if err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		User: User{
			Id:       response.User.Id,
			Name:     response.User.Name,
			Score:    response.User.Score,
			UserName: response.User.UserName,
		},
	}, nil
}

func (service *UserService) GetUsers(request GetUserRequest) GetUsersResponse {
	return GetUsersResponse{}
}
