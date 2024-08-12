package service

import (
	"app/user/api/controller/create"
	request "app/user/api/controller/delete"
	"app/user/api/controller/update"
	"app/user/entity"
)

type UserService interface {
	Create(createUserRequest create.CreateUserRequest) error
	Update(user update.UpdateUserRequest) error
	FindByPk(pk int) (*entity.User, error)
	Delete(user request.DeleteUserRequest) error
}
