package service

import (
	"errors"

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
