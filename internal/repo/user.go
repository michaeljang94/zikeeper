package repo

import (
	"database/sql"
)

type UserRepo struct {
	Db *sql.DB
}

type User struct {
	Id       string
	Name     string
	Score    int
	UserName string
	Password string
	Role     string
}

type GetUserRequest struct {
	Id       string
	UserName string
}

type GetUserResponse struct {
	User User
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	Users []User
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
	Users []ScoreboardUser
}

func (repo *UserRepo) GetUsers(request GetUsersRequest) (GetUsersResponse, error) {
	rows, err := repo.Db.Query("SELECT id, name, score, username, role FROM users")

	if err != nil {
		return GetUsersResponse{}, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Score, &user.UserName, &user.Role); err != nil {
			return GetUsersResponse{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return GetUsersResponse{}, err
	}

	return GetUsersResponse{
		Users: users,
	}, nil
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
	row := repo.Db.QueryRow("SELECT id, name, score, username, role FROM users WHERE username = ?", request.UserName)

	user := User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Score, &user.UserName, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return GetUserResponse{}, err
		}
	}

	return GetUserResponse{
		User: user,
	}, nil
}

func (repo *UserRepo) GetUsersScoreboard(request GetUsersScoreboardRequest) (GetUsersScoreboardResponse, error) {
	rows, err := repo.Db.Query("SELECT username, score, DENSE_RANK() OVER (ORDER BY score DESC) AS 'rank' FROM users WHERE role = 'user' LIMIT ?", request.Limit)

	if err != nil {
		return GetUsersScoreboardResponse{}, err
	}

	defer rows.Close()

	var users []ScoreboardUser

	for rows.Next() {
		var user ScoreboardUser
		if err := rows.Scan(&user.Username, &user.Score, &user.Rank); err != nil {
			return GetUsersScoreboardResponse{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return GetUsersScoreboardResponse{}, err
	}

	return GetUsersScoreboardResponse{
		Users: users,
	}, nil
}

type UpdateUserByUsernameRequest struct {
	Username string
	Score    int
	Role     string
}

type UpdateUserByUsernameResponse struct {
}

func (repo *UserRepo) UpdateUserByUsername(request UpdateUserByUsernameRequest) (UpdateUserByUsernameResponse, error) {
	_, err := repo.Db.Exec("UPDATE users SET score = ?, role = ? WHERE username = ?", request.Score, request.Role, request.Username)

	if err != nil {
		return UpdateUserByUsernameResponse{}, err
	}

	return UpdateUserByUsernameResponse{}, nil
}
