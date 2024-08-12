package request

type UpdateUserRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
