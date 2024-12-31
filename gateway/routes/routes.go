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

	router.GET(
		"/product/:id",
		controller.GetProduct(srv))

	router.GET(
		"/products",
		controller.GetProducts(srv))
}

func ProtectedRoutes(router *gin.Engine, srv *server.Server) {

	router.GET(
		"/account",
		middleware.AuthMiddleware(srv),
		controller.GetUser(srv))

	router.POST(
		"/cart/add",
		middleware.CartMiddleware(srv),
		controller.AddItemToCart(srv),
	)

	router.POST(
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

	router.POST(
		"/product/add",
		middleware.SellerMiddleware(srv),
		controller.PostProduct(srv))

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
		"/address/delete/:id",
		middleware.AuthMiddleware(srv),
		controller.DeleteAddress(srv))

}
