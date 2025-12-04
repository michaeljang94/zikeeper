package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/michaeljang94/zikeeper/internal/repo"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	UserName string `json:"username"`
	Password string `json:"password"`
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

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	User User `json:"user"`
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
			Password: response.User.Password,
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
			Password: response.User.Password,
		},
	}, nil
}

func (service *UserService) GetUsers(request GetUserRequest) GetUsersResponse {
	return GetUsersResponse{}
}

func generateRandomPass() string {
	passLength := 5
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := make([]byte, passLength)

	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	return string(result)
}

func (service *UserService) CreateUser(request CreateUserRequest) (CreateUserResponse, error) {
	id := uuid.New()

	repoRequest := repo.CreateUserRequest{
		Id:       id.String(),
		Name:     request.Name,
		Score:    0,
		UserName: request.Name,
		Password: generateRandomPass(),
	}

	response, err := service.Repo.CreateUser(repoRequest)

	if err != nil {
		fmt.Println(err)
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		User: User{
			Id:       response.User.Id,
			Name:     response.User.Name,
			Score:    response.User.Score,
			UserName: response.User.UserName,
			Password: response.User.Password,
		},
	}, nil
}
