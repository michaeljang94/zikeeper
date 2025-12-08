package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	iat := time.Now().UTC()
	exp := iat.Add(time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      iat.Unix(),
		"exp":      exp.Unix(),
		"username": user.AuthUser.Username,
		"role":     user.AuthUser.Role,
	})

	// TODO change this
	key := "SUPERSECRETKEY"
	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		fmt.Println(err)
		return AuthenticateUserResponse{
			Status: "FAILED",
		}, err
	}

	fmt.Println(tokenString)

	return AuthenticateUserResponse{
		Status: "OK",
	}, nil
}
