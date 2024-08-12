package service

import (
	"app/user/controller/create"
	request "app/user/controller/delete"
	"app/user/controller/update"
	"app/user/entity"
)

type UserService interface {
	Create(createUserRequest create.CreateUserRequest) error
	Update(user update.UpdateUserRequest) error
	FindByPk(pk int) (*entity.User, error)
	Delete(user request.DeleteUserRequest) error
}
