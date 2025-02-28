package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func AccountRoutes(router *gin.Engine, srv *server.Server) {
	router.POST(
		"/signup",
		//validator.SignUpValidator(),
		controller.SignUp(srv),
	)

	router.POST(
		"/login",
		//validator.LoginValidator(),
		controller.Login(srv),
	)
	router.GET(
		"/account",
		middleware.AuthMiddleware(srv),
		controller.GetUser(srv))

}
