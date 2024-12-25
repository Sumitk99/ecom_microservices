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

func ProtectedRoutes(incomingRoutes *gin.Engine, srv *server.Server) {
	incomingRoutes.GET("/account", middleware.AuthMiddleware(srv), controller.GetUser(srv))

	incomingRoutes.POST("/cart/add", middleware.CartMiddleware(srv), controller.AddItemToCart(srv))
	incomingRoutes.GET("/cart/get", middleware.CartMiddleware(srv), controller.GetCart(srv))
	incomingRoutes.PUT("/cart/remove", middleware.CartMiddleware(srv), controller.RemoveItemFromCart(srv))
	incomingRoutes.PUT("/cart/update", middleware.CartMiddleware(srv), controller.UpdateCart(srv))
	incomingRoutes.DELETE("/cart/delete", middleware.CartMiddleware(srv), controller.DeleteCart(srv))
	incomingRoutes.POST("/cart/checkout", middleware.AuthMiddleware(srv), controller.Checkout(srv))

	incomingRoutes.GET("/user/orders", middleware.AuthMiddleware(srv), controller.GetOrders(srv))
	incomingRoutes.GET("/user/order/:id", middleware.AuthMiddleware(srv), controller.GetOrder(srv))
}
