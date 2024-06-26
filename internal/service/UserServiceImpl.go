package service

import (
	request "app/internal/api/controller/request"
	"app/internal/entity"
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

func (u *UserServiceImpl) Create(createUserRequest request.CreateUserRequest) error {
	user := dao.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	u.db.Save(&user)

	return nil
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) error {
	u.db.
		Model(&dao.User{}).
		Where("id=?", user.Id).
		Updates(dao.User{Password: user.Password})

	return nil
}

func (u *UserServiceImpl) FindByPk(pk int) (*dao.User, error) {
	user := dao.User{}

	result := u.db.First(&user, pk)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(user)

	return &user, nil
}

func (u *UserServiceImpl) Delete(user request.DeleteUserRequest) error {

	u.db.Delete(&dao.User{}, user.Id)

	return nil
}
