package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine, srv *server.Server) {
	router.POST(
		"/card/add",
		middleware.AuthMiddleware(srv),
		controller.AddCard(srv),
	)
	router.DELETE(
		"/card/remove",
		middleware.AuthMiddleware(srv),
		controller.RemoveCard(srv),
	)
	router.GET(
		"/cards",
		middleware.AuthMiddleware(srv),
		controller.ListCards(srv),
	)

	router.POST(
		"/payment/process",
		middleware.AuthMiddleware(srv),
		controller.ProcessPayment(srv),
	)
}
