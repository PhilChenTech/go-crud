package service

import (
	"app/user/controller/create"
	request "app/user/controller/delete"
	"app/user/controller/update"
	"app/user/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) UserService {
	return &UserServiceImpl{
		db: db,
	}
}

func (u *UserServiceImpl) Create(createUserRequest create.CreateUserRequest) error {
	user := entity.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	u.db.Save(&user)

	return nil
}

func (u *UserServiceImpl) Update(user update.UpdateUserRequest) error {
	u.db.
		Model(&entity.User{}).
		Where("id=?", user.Id).
		Updates(entity.User{Password: user.Password})

	return nil
}

func (u *UserServiceImpl) FindByPk(pk int) (*entity.User, error) {
	user := entity.User{}

	result := u.db.First(&user, pk)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(user)

	return &user, nil
}

func (u *UserServiceImpl) Delete(user request.DeleteUserRequest) error {

	u.db.Delete(&entity.User{}, user.Id)

	return nil
}
