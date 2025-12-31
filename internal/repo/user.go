package repo

import (
	"context"
	"database/sql"
)

type UserRepo struct {
	Db *sql.DB
}

type StudentNumber struct {
	Year   int
	Class  int
	Number int
}

type User struct {
	Id            string
	Name          string
	Score         int
	UserName      string
	Password      string
	Role          string
	StudentNumber StudentNumber
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
	rows, err := repo.Db.Query("SELECT id, name, score, username, role, student_year, student_class, student_number FROM users")

	if err != nil {
		return GetUsersResponse{}, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		user := User{
			StudentNumber: StudentNumber{},
		}
		if err := rows.Scan(&user.Id, &user.Name, &user.Score, &user.UserName, &user.Role,
			&user.StudentNumber.Year, &user.StudentNumber.Class, &user.StudentNumber.Number); err != nil {
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
	row := repo.Db.QueryRow("SELECT id, name, score, username, role, student_year, student_class, student_number FROM users WHERE username = ?", request.UserName)

	user := User{
		StudentNumber: StudentNumber{},
	}

	if err := row.Scan(&user.Id, &user.Name, &user.Score, &user.UserName, &user.Role,
		&user.StudentNumber.Year, &user.StudentNumber.Class, &user.StudentNumber.Number); err != nil {
		if err == sql.ErrNoRows {
			return GetUserResponse{}, err
		}
	}

	return GetUserResponse{
		User: user,
	}, nil
}

func (repo *UserRepo) GetUsersScoreboard(request GetUsersScoreboardRequest) (GetUsersScoreboardResponse, error) {
	rows, err := repo.Db.Query("SELECT username, score, DENSE_RANK() OVER (ORDER BY score DESC) AS user_rank FROM users WHERE role = 'user' LIMIT ?", request.Limit)

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

type GetPlayerRankingByUsernameRequest struct {
	Username string
}

type GetPlayerRankingByUsernameResponse struct {
	User ScoreboardUser
}

func (repo *UserRepo) GetPlayerRankingByUsername(request GetPlayerRankingByUsernameRequest) (GetPlayerRankingByUsernameResponse, error) {
	row := repo.Db.QueryRow("SELECT username, score, user_rank FROM (SELECT username, score, DENSE_RANK() OVER (ORDER BY score DESC) AS user_rank FROM users WHERE role = 'user') AS rank_table WHERE username = ?", request.Username)

	user := ScoreboardUser{}
	if err := row.Scan(&user.Username, &user.Score, &user.Rank); err != nil {
		if err == sql.ErrNoRows {
			return GetPlayerRankingByUsernameResponse{}, err
		}
	}

	return GetPlayerRankingByUsernameResponse{
		User: user,
	}, nil
}

type TransferScoreByUsernameRequest struct {
	From   string
	To     string
	Amount int
}

type TransferScoreByUsernameResponse struct {
}

func (repo *UserRepo) TransferScoreByUsername(request TransferScoreByUsernameRequest) (TransferScoreByUsernameResponse, error) {
	ctx := context.Background()

	tx, err := repo.Db.BeginTx(ctx, nil)

	if err != nil {
		return TransferScoreByUsernameResponse{}, err
	}

	// Get score value of A
	// row := repo.Db.QueryRow("SELECT * FROM users WHERE id = ?", request.Id)
	fromRes := tx.QueryRow("SELECT score FROM users WHERE username = ?", request.From)
	fromUser := User{}
	if err := fromRes.Scan(&fromUser.Score); err != nil {
		if err == sql.ErrNoRows {
			return TransferScoreByUsernameResponse{}, err
		}
	}

	if err != nil {
		return TransferScoreByUsernameResponse{}, err
	}

	// Get score value of B
	toRes := tx.QueryRow("SELECT score FROM users WHERE username = ?", request.To)

	if err != nil {
		return TransferScoreByUsernameResponse{}, err
	}

	toUser := User{}
	if err := toRes.Scan(&toUser.Score); err != nil {
		if err == sql.ErrNoRows {
			return TransferScoreByUsernameResponse{}, err
		}
	}

	// TODO check that these dont overflow to negative
	newFromScore := fromUser.Score - request.Amount
	newToScore := toUser.Score + request.Amount

	// Remove from A
	_, err = tx.ExecContext(ctx, "UPDATE users SET score = ? WHERE username = ?", newFromScore, request.From)

	if err != nil {
		tx.Rollback()
		return TransferScoreByUsernameResponse{}, err
	}

	// Add to B
	_, err = tx.ExecContext(ctx, "UPDATE users SET score = ? WHERE username = ?", newToScore, request.To)

	if err != nil {
		tx.Rollback()
		return TransferScoreByUsernameResponse{}, err
	}

	txErr := tx.Commit()
	if txErr != nil {
		return TransferScoreByUsernameResponse{}, txErr
	}

	return TransferScoreByUsernameResponse{}, nil
}
