package request

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


type UpdateUserRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}
