package api

import (
	"app/user/api/controller/user"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router         *gin.Engine
	userController *controller.UserController
}

func NewHTTPServer(router *gin.Engine, userController *controller.UserController) *HTTPServer {

	user := router.Group("/user")
	{
		user.POST("/v1", userController.Create)
		user.GET("/v1/:pk", userController.FindByPk)
		user.PUT("/v1", userController.Update)
		user.DELETE("/v1", userController.Delete)
	}

	return &HTTPServer{
		Router:         router,
		userController: userController,
	}

}
