package service

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/michaeljang94/zikeeper/internal/repo"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type GetUserRequest struct {
	Id string `json:"id" binding:"required"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	User User `json:"user"`
}

type UserService struct {
	Repo *repo.UserRepo
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
			Id:    response.User.Id,
			Name:  response.User.Name,
			Score: response.User.Score,
		},
	}, nil
}

func (service *UserService) GetUsers(request GetUserRequest) GetUsersResponse {
	return GetUsersResponse{}
}

func (service *UserService) CreateUser(request CreateUserRequest) (CreateUserResponse, error) {
	id := uuid.New()

	repoRequest := repo.CreateUserRequest{
		Id:    id.String(),
		Name:  request.Name,
		Score: 0,
	}

	response, err := service.Repo.CreateUser(repoRequest)

	if err != nil {
		fmt.Println(err)
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		User: User{
			Id:    response.User.Id,
			Name:  response.User.Name,
			Score: response.User.Score,
		},
	}, nil
}
