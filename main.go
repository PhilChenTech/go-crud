package main

import (
	"app/common"
	"app/common/config"
	"app/common/database"
	"app/user/controller"
	"app/user/controller/findById"
	"log"

	"app/user/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// init config
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// init database
	database, err := database.NewMysql(
		config.Connections.MySQL.User,
		config.Connections.MySQL.Password,
		config.Connections.MySQL.Host,
		config.Connections.MySQL.DB,
	)
	if err != nil {
		log.Fatal(err)
	}

	// init controller / service
	userService := service.NewUserServiceImpl(database)
	userController := controller.NewUserController(userService)
	findUserByIdController := findById.FindUserController(userService)
	// init server
	server := common.NewHTTPServer(gin.Default(), userController, findUserByIdController)

	// run server
	err = server.Router.Run(":" + config.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}
