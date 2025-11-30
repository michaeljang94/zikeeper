package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	Id    string
	Name  string
	Score int
}

type GetUserRequest struct {
	Id string
}

type GetUserResponse struct {
	User User
}

type CreateUserRequest struct {
	Id    string
	Name  string
	Score int
}

type CreateUserResponse struct {
	User User
}

type UserRepo struct {
	Db *sql.DB
}

func (repo *UserRepo) CreateUser(request CreateUserRequest) (CreateUserResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO users (id, name, score) VALUES (?, ?, ?)",
		request.Id, request.Name, request.Score)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return CreateUserResponse{}, errors.New("duplicate entry")
		}

		fmt.Println(err)
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		User: User{
			Id:    request.Id,
			Name:  request.Name,
			Score: request.Score,
		},
	}, nil
}

func (repo *UserRepo) GetUser(request GetUserRequest) (GetUserResponse, error) {
	row := repo.Db.QueryRow("SELECT * FROM users WHERE id = ?", request.Id)

	user := User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Score); err != nil {
		if err == sql.ErrNoRows {
			return GetUserResponse{}, err
		}
	}

	return GetUserResponse{
		User: user,
	}, nil
}
