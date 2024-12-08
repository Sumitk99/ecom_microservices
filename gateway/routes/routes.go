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
	incomingRoutes.GET("/cart/req", controller.RequestGuestId(srv))
}

// func CartRoutes(incomingRoutes *gin.Engine, srv *server.Server) {
//
//		cartGroup := incomingRoutes.Group("/cart")
//		cartGroup.Use(middleware.CartMiddleware(srv))
//		{
//			//incomingRoutes.GET("/cart", controller.GetCart(srv))
//			incomingRoutes.POST("/add", controller.Add(srv))
//			//incomingRoutes.DELETE("/cart", controller.RemoveFromCart(srv))
//
//		}
//		incomingRoutes.Use(middleware.CartMiddleware(srv))
//	}
func ProtectedRoutes(incomingRoutes *gin.Engine, srv *server.Server) {
	incomingRoutes.Use(middleware.AuthMiddleware(srv))
	incomingRoutes.GET("/account", controller.GetUser(srv))
}
