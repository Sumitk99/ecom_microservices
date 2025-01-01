package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func CatalogRoutes(router *gin.Engine, srv *server.Server) {
	router.GET(
		"/product/:id",
		controller.GetProduct(srv))

	router.GET(
		"/products",
		controller.GetProducts(srv))
	router.POST(
		"/product/add",
		middleware.SellerMiddleware(srv),
		controller.PostProduct(srv))

}
