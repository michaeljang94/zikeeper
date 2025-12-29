package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type AuthRepo struct {
	Db *sql.DB
}

type GetAuthUserbyUsernameRequest struct {
	Username string
}

type GetAuthUserbyUsernameResponse struct {
	AuthUser AuthUser
}

type AuthUser struct {
	Id       string
	Username string
	Pincode  string
	Role     string
}

type CreateUserRequest struct {
	Id            string
	Name          string
	Score         int
	UserName      string
	Password      string
	Pincode       string
	StudentNumber StudentNumber
}

type CreateUserResponse struct {
	User User
}

func (repo *AuthRepo) GetUserByUsername(request GetAuthUserbyUsernameRequest) (GetAuthUserbyUsernameResponse, error) {
	row := repo.Db.QueryRow("SELECT username, pincode, role FROM users WHERE username = ?", request.Username)

	user := AuthUser{}
	if err := row.Scan(&user.Username, &user.Pincode, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return GetAuthUserbyUsernameResponse{}, err
		}
	}

	return GetAuthUserbyUsernameResponse{
		AuthUser: user,
	}, nil
}

func (repo *AuthRepo) CreateNewUser(request CreateUserRequest) (CreateUserResponse, error) {
	_, err := repo.Db.Exec("INSERT INTO users (id, name, score, username, password, pincode, role, student_year, student_class, student_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		request.Id, request.Name, request.Score, request.UserName, request.Password, request.Pincode, "user",
		request.StudentNumber.Year, request.StudentNumber.Class, request.StudentNumber.Number)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return CreateUserResponse{}, errors.New("duplicate entry")
		}

		fmt.Println(err)
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		User: User{
			Id:       request.Id,
			Name:     request.Name,
			Score:    request.Score,
			UserName: request.UserName,
		},
	}, nil
}
