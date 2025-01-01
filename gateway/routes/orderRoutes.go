package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, srv *server.Server) {
	router.GET("/user/orders",
		middleware.AuthMiddleware(srv),
		controller.GetOrders(srv))

	router.GET(
		"/user/order/:id",
		middleware.AuthMiddleware(srv),
		controller.GetOrder(srv))
}
