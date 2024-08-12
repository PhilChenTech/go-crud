package controller

import (
	response "app/user/api/controller/create"
	request "app/user/api/controller/delete"
	"app/user/api/controller/update"
	"app/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// aaahandleError 处理错误并返回 JSON 响应
func handleError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}

// Create 处理创建用户请求
func (u *UserController) Create(ctx *gin.Context) {
	var userRequest response.CreateUserRequest
	if err := ctx.BindJSON(&userRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := u.userService.Create(userRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	userResponse := response.CreateUserResponse{
		Email: userRequest.Email,
	}

	ctx.JSON(http.StatusOK, userResponse)
}

// Update 更新用户信息
func (u *UserController) Update(ctx *gin.Context) {
	var updateRequest update.UpdateUserRequest
	if err := ctx.BindJSON(&updateRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := u.userService.Update(updateRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "update ok"})
}

// Delete 删除用户
func (u *UserController) Delete(ctx *gin.Context) {
	var userRequest request.DeleteUserRequest
	if err := ctx.BindJSON(&userRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := u.userService.Delete(userRequest); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "delete ok"})
}
