package service

import (
	request "app/internal/api/controller/request"
	"app/internal/model/dao"
)

type UserService interface {
	Create(createUserRequest request.CreateUserRequest) error
	Update(user request.UpdateUserRequest) error
	FindByPk(pk int) (*dao.User, error)
	Delete(user request.DeleteUserRequest) error
}
