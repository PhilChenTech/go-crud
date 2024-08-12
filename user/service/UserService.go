package service

import (
	request "app/user/api/controller/request"
	"app/user/entity"
)

type UserService interface {
	Create(createUserRequest request.CreateUserRequest) error
	Update(user request.UpdateUserRequest) error
	FindByPk(pk int) (*dao.User, error)
	Delete(user request.DeleteUserRequest) error
}
