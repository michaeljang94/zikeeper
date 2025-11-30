package repo

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

func GetUser(request GetUserRequest) GetUserResponse {
	// TODO fetch from an actual DB

	return GetUserResponse{
		User: User{
			Id:    "1234",
			Name:  "Muone",
			Score: 1000,
		},
	}
}
