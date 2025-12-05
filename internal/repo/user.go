package repo

import (
	"database/sql"
)

type User struct {
	Id       string
	Name     string
	Score    int
	UserName string
	Password string
}

type GetUserRequest struct {
	Id       string
	UserName string
}

type GetUserResponse struct {
	User User
}

type UserRepo struct {
	Db *sql.DB
}

func (repo *UserRepo) GetUser(request GetUserRequest) (GetUserResponse, error) {
	row := repo.Db.QueryRow("SELECT * FROM users WHERE id = ?", request.Id)

	user := User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Score, &user.UserName); err != nil {
		if err == sql.ErrNoRows {
			return GetUserResponse{}, err
		}
	}

	return GetUserResponse{
		User: user,
	}, nil
}

func (repo *UserRepo) GetUserByUserName(request GetUserRequest) (GetUserResponse, error) {
	row := repo.Db.QueryRow("SELECT id, name, score, username FROM users WHERE username = ?", request.UserName)

	user := User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Score, &user.UserName); err != nil {
		if err == sql.ErrNoRows {
			return GetUserResponse{}, err
		}
	}

	return GetUserResponse{
		User: user,
	}, nil
}
