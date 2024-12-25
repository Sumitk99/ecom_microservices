package controller

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func GetOrder(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("id")
		ctx := CreateUserMetaData(c)
		res, err := srv.GetOrder(ctx, orderId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func GetOrders(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := CreateUserMetaData(c)
		res, err := srv.GetOrders(ctx, c.GetString("id"))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func CreateUserMetaData(c *gin.Context) context.Context {
	md := metadata.New(map[string]string{
		"UserID": c.GetString("id"),
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}
