package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type AuthService struct {
	UserRepo *repo.UserRepo
	AuthRepo *repo.AuthRepo
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Pincode  string `json:"pincode"`
}

type AuthenticateUserRequest struct {
	AuthUser AuthUser
}

type AuthenticateUserResponse struct {
	Status string `json:"status"`
}

type CreateAuthUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Pincode  string `json:"pincode"`
	Password string `json:"password"`
}

type CreateAuthUserResponse struct {
	AuthUser AuthUser `json:"user"`
}

func (service *AuthService) CreateNewUser(request CreateAuthUserRequest) (CreateAuthUserResponse, error) {
	id := uuid.New()

	repoRequest := repo.CreateUserRequest{
		Id:       id.String(),
		Name:     request.Name,
		Score:    0,
		UserName: request.Username,
		Password: request.Password,
		Pincode:  request.Pincode,
	}

	response, err := service.AuthRepo.CreateNewUser(repoRequest)

	if err != nil {
		fmt.Println(err)
		return CreateAuthUserResponse{}, err
	}

	return CreateAuthUserResponse{
		AuthUser: AuthUser{
			Username: response.User.UserName,
		},
	}, nil
}

func (service *AuthService) AuthenticateUser(request AuthenticateUserRequest) (AuthenticateUserResponse, error) {
	authUser := request.AuthUser

	// Get user
	getAuthUserRequest := repo.GetAuthUserbyUsernameRequest{
		Username: authUser.Username,
	}

	user, err := service.AuthRepo.GetUserByUsername(getAuthUserRequest)
	if err != nil {
		return AuthenticateUserResponse{
			Status: "FAILED",
		}, err
	}

	if user.AuthUser.Pincode != authUser.Pincode {
		return AuthenticateUserResponse{
			Status: "FAILED",
		}, errors.New("FAILED")
	}

	return AuthenticateUserResponse{
		Status: "OK",
	}, nil
}
