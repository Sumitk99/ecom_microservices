package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine, srv *server.Server) {
	router.GET(
		"/cart/req", // issue a guest token to a user who is not logged in
		controller.RequestGuestId(srv))

	router.POST(
		"/cart/add/:product_id/:quantity",
		middleware.CartMiddleware(srv),
		controller.AddItemToCart(srv),
	)
	router.GET(
		"/cart/get",
		middleware.CartMiddleware(srv),
		controller.GetCart(srv))

	router.GET(
		"/cart/get/:cart_id",
		middleware.CartMiddleware(srv),
		controller.GetCart(srv))

	router.DELETE(
		"/cart/remove/:product_id",
		middleware.CartMiddleware(srv),
		controller.RemoveItemFromCart(srv),
	)
	router.PUT(
		"/cart/update/:product_id/:quantity",
		//validator.ValidateCartOpsReq(),
		middleware.CartMiddleware(srv),
		controller.UpdateCart(srv),
	)

	router.DELETE(
		"/cart/delete/{cart_id}",
		middleware.CartMiddleware(srv),
		controller.DeleteCart(srv))

	router.POST(
		"/cart/checkout",
		middleware.AuthMiddleware(srv),
		controller.Checkout(srv))

}
