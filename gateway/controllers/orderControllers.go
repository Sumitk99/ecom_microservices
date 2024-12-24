package controller

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostOrder(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := GetCartContext(c)
		err := srv.DeleteCart(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(200, gin.H{"Message": "Cart Successfully Deleted"})
	}
}
