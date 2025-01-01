package routes

import (
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, srv *server.Server) {
	AccountRoutes(router, srv)
	CatalogRoutes(router, srv)
	CartRoutes(router, srv)
	AddressRoutes(router, srv)
	OrderRoutes(router, srv)
}
