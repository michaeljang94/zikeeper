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

type CreateUserRequest struct {
	Name string
}

type CreateUserResponse struct {
	User User
}

func CreateUser(request CreateUserRequest) CreateUserResponse {
	return CreateUserResponse{
		User: User{
			Id:    "1234",
			Name:  request.Name,
			Score: 0,
		},
	}
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
