package repo

import "database/sql"

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
}

func (repo *AuthRepo) GetUserByUsername(request GetAuthUserbyUsernameRequest) (GetAuthUserbyUsernameResponse, error) {
	row := repo.Db.QueryRow("SELECT username, pincode FROM users WHERE username = ?", request.Username)

	user := AuthUser{}
	if err := row.Scan(&user.Username, &user.Pincode); err != nil {
		if err == sql.ErrNoRows {
			return GetAuthUserbyUsernameResponse{}, err
		}
	}

	return GetAuthUserbyUsernameResponse{
		AuthUser: user,
	}, nil
}
