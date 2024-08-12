package controller

import (
	"app/user/controller/findById"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router                 *gin.Engine
	userController         *UserController
	findUserByIdController *findById.FindUserByIdController
}

func NewHTTPServer(router *gin.Engine, userController *UserController, findUserByIdController *findById.FindUserByIdController) *HTTPServer {

	user := router.Group("/user")
	{
		user.POST("/v1", userController.Create)
		user.GET("/v1/:pk", findUserByIdController.FindByPk)
		user.PUT("/v1", userController.Update)
		user.DELETE("/v1", userController.Delete)
	}

	return &HTTPServer{
		Router:         router,
		userController: userController,
	}

}
