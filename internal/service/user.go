package service

import "github.com/michaeljang94/zikeeper/internal/repo"

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

func GetUser(request GetUserRequest) GetUserResponse {
	repoRequest := repo.GetUserRequest{
		Id: request.Id,
	}

	response := repo.GetUser(repoRequest)

	return GetUserResponse{
		User: User{
			Id:    response.User.Id,
			Name:  response.User.Name,
			Score: response.User.Score,
		},
	}
}

func GetUsers(request GetUserRequest) GetUsersResponse {
	return GetUsersResponse{}
}

func CreateUser(request CreateUserRequest) CreateUserResponse {
	repoRequest := repo.CreateUserRequest{
		Name: request.Name,
	}

	response := repo.CreateUser(repoRequest)

	return CreateUserResponse{
		User: User{
			Id:    response.User.Id,
			Name:  response.User.Name,
			Score: response.User.Score,
		},
	}
}
