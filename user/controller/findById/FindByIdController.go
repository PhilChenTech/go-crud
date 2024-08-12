package findById

import (
	"app/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FindUserByIdController struct {
	userService service.UserService
}

func FindUserController(userService service.UserService) *FindUserByIdController {
	return &FindUserByIdController{
		userService: userService,
	}
}

// aaahandleError 处理错误并返回 JSON 响应
func aaahandleError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}

// FindByPk 根据主键查找用户
func (u *FindUserByIdController) FindByPk(ctx *gin.Context) {
	pk, err := strconv.Atoi(ctx.Param("pk"))
	if err != nil {
		aaahandleError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := u.userService.FindByPk(pk)
	if err != nil {
		aaahandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
