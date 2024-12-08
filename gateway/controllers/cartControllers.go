package controller

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func RequestGuestId(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		guestId, err := srv.IssueGuestCartToken(ctx)
		fmt.Println("error is : ", err)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"guestId": guestId})
	}
}

//func AddItemToCart(srv *server.Server) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ctx := context.WithValue(context.Background(), "UserID", c.GetString("id"))
//		ctx = context.WithValue(ctx, "GuestID", c.GetString("GuestID"))
//		var form models.AddItemRequest
//		err := c.BindJSON(&form)
//		if err != nil {
//			c.JSON(400, gin.H{"error": err})
//			return
//		}
//
//
//	}
//}
