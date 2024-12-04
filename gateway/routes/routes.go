package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(incomingRoutes *gin.Engine, srv *server.Server) {
	incomingRoutes.POST("/signup", controller.SignUp(srv))
	incomingRoutes.GET("/login", controller.Login(srv))
}
func ProtectedRoutes(incomingRoutes *gin.Engine, srv *server.Server) {
	incomingRoutes.Use(middleware.Authenticate(srv))
	incomingRoutes.GET("/account", controller.GetUser(srv))
}
