package service

import (
	"github.com/michaeljang94/zikeeper/internal/repo"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	UserName string `json:"username"`
	Rank     int    `json:"rank"`
	Role     string `json:"role"`
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
			Role:     response.User.Role,
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

func (service *UserService) GetUsers(request GetUsersRequest) (GetUsersResponse, error) {
	repoRequest := repo.GetUsersRequest{}

	repoResponse, err := service.Repo.GetUsers(repoRequest)

	if err != nil {
		return GetUsersResponse{}, err
	}

	var users []User
	for i := range repoResponse.Users {
		user := User{
			Id:       repoResponse.Users[i].Id,
			Name:     repoResponse.Users[i].Name,
			Score:    repoResponse.Users[i].Score,
			UserName: repoResponse.Users[i].UserName,
			Role:     repoResponse.Users[i].Role,
		}

		users = append(users, user)
	}

	return GetUsersResponse{
		Users: users,
	}, nil
}

type ScoreboardUser struct {
	Username string
	Score    int
	Rank     int
}

type GetUsersScoreboardRequest struct {
	Limit int
}

type GetUsersScoreboardResponse struct {
	Users []ScoreboardUser `json:"users"`
}

func (service *UserService) GetUsersScoreboard(request GetUsersScoreboardRequest) (GetUsersScoreboardResponse, error) {
	req := repo.GetUsersScoreboardRequest{
		Limit: request.Limit,
	}

	res, err := service.Repo.GetUsersScoreboard(req)

	if err != nil {
		return GetUsersScoreboardResponse{}, err
	}

	var users []ScoreboardUser
	for i := range res.Users {
		user := ScoreboardUser{
			Username: res.Users[i].Username,
			Score:    res.Users[i].Score,
			Rank:     res.Users[i].Rank,
		}

		users = append(users, user)
	}

	return GetUsersScoreboardResponse{
		Users: users,
	}, nil
}

type UpdateUserByUsernameRequest struct {
	Username string
	Score    int    `json:"score"`
	Role     string `json:"role"`
}

type UpdateUserByUsernameResponse struct {
}

func (service *UserService) UpdateUserByUsername(request UpdateUserByUsernameRequest) (UpdateUserByUsernameResponse, error) {
	req := repo.UpdateUserByUsernameRequest{
		Username: request.Username,
		Score:    request.Score,
		Role:     request.Role,
	}

	_, err := service.Repo.UpdateUserByUsername(req)

	if err != nil {
		return UpdateUserByUsernameResponse{}, err
	}

	return UpdateUserByUsernameResponse{}, nil
}
