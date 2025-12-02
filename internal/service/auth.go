package service

import (
	"errors"

	"github.com/michaeljang94/zikeeper/internal/repo"
)

type AuthService struct {
	UserRepo *repo.UserRepo
}

type AuthUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateUserRequest struct {
	User AuthUser
}

type AuthenticateUserResponse struct {
	Status string
}

func (service *AuthService) AuthenticateUser(request AuthenticateUserRequest) (AuthenticateUserResponse, error) {
	authUser := request.User

	// Get user
	getUserRequest := repo.GetUserRequest{
		UserName: authUser.UserName,
	}

	user, err := service.UserRepo.GetUserByUserName(getUserRequest)
	if err != nil {
		return AuthenticateUserResponse{
			Status: "FAILED",
		}, err
	}

	if user.User.Password != authUser.Password {
		return AuthenticateUserResponse{
			Status: "FAILED",
		}, errors.New("FAILED")
	}

	return AuthenticateUserResponse{
		Status: "OK",
	}, nil
}
