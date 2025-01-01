package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine, srv *server.Server) {
	router.GET(
		"/cart/req",
		controller.RequestGuestId(srv))

	router.POST(
		"/cart/add",
		middleware.CartMiddleware(srv),
		controller.AddItemToCart(srv),
	)

	router.GET(
		"/cart/get",
		middleware.CartMiddleware(srv),
		controller.GetCart(srv))

	router.PUT(
		"/cart/remove",
		//validator.ValidateRemoveFromCartReq(),
		middleware.CartMiddleware(srv),
		controller.RemoveItemFromCart(srv),
	)
	router.PUT(
		"/cart/update",
		//validator.ValidateCartOpsReq(),
		middleware.CartMiddleware(srv),
		controller.UpdateCart(srv),
	)

	router.DELETE(
		"/cart/delete",
		middleware.CartMiddleware(srv),
		controller.DeleteCart(srv))

	router.POST(
		"/cart/checkout",
		middleware.AuthMiddleware(srv),
		controller.Checkout(srv))

}
