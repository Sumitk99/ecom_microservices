package routes

import (
	controller "github.com/Sumitk99/ecom_microservices/gateway/controllers"
	"github.com/Sumitk99/ecom_microservices/gateway/middleware"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/Sumitk99/ecom_microservices/gateway/validator"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(router *gin.Engine, srv *server.Server) {
	router.POST(
		"/signup",
		validator.SignUpValidator(),
		controller.SignUp(srv),
	)

	router.POST(
		"/login",
		//validator.LoginValidator(),
		controller.Login(srv),
	)

	router.GET(
		"/cart/req",
		controller.RequestGuestId(srv))
}

func ProtectedRoutes(router *gin.Engine, srv *server.Server) {

	router.GET(
		"/account",
		middleware.AuthMiddleware(srv),
		controller.GetUser(srv))

	router.POST(
		"/cart/add",
		validator.ValidateCartOpsReq(),
		middleware.CartMiddleware(srv),
		controller.AddItemToCart(srv),
	)

	router.GET(
		"/cart/get",
		middleware.CartMiddleware(srv),
		controller.GetCart(srv))
	router.PUT(
		"/cart/remove",
		validator.ValidateRemoveFromCartReq(),
		middleware.CartMiddleware(srv),
		controller.RemoveItemFromCart(srv),
	)
	router.PUT(
		"/cart/update",
		validator.ValidateCartOpsReq(),
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

	router.GET("/user/orders",
		middleware.AuthMiddleware(srv),
		controller.GetOrders(srv))

	router.GET(
		"/user/order/:id",
		middleware.AuthMiddleware(srv),
		controller.GetOrder(srv))
}
