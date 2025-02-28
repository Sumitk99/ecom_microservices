package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func AddressRoutes(router *gin.Engine, srv *server.Server) {
	router.GET(
		"/address/get",
		middleware.AuthMiddleware(srv),
		controller.GetAddresses(srv))

	router.POST(
		"/address/add",
		middleware.AuthMiddleware(srv),
		controller.AddAddress(srv))

	router.GET(
		"/address/get/:id",
		middleware.AuthMiddleware(srv),
		controller.GetAddress(srv))

	router.DELETE(
		"/address/delete/:address_id",
		middleware.AuthMiddleware(srv),
		controller.DeleteAddress(srv))

}
