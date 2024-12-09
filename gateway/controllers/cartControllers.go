package controller

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"net/http"
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

func AddItemToCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		err := c.BindJSON(&form)

		ctx := context.WithValue(context.Background(), "UserID", c.GetString("id"))
		ctx = context.WithValue(ctx, "CartID", form.CartName)
		ctx = context.WithValue(ctx, "GuestID", c.GetString("GuestID"))
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		fmt.Printf("%s %s %s\n", form.CartName, form.Quantity, form.ProductID)
		res, err := srv.AddItemToCart(ctx, form.ProductID, form.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(200, res)
	}
}

func GetCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), "UserID", c.GetString("id"))
		ctx = context.WithValue(ctx, "CartID", c.GetString("CartID"))
		ctx = context.WithValue(ctx, "GuestID", c.GetString("GuestID"))
		fmt.Printf("Triplets : %s %s %s\n", ctx.Value("UserID"), ctx.Value("CartID"), ctx.Value("GuestID"))
		res, err := srv.GetCart(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(200, res)

	}
}

func RemoveItemFromCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		err := c.BindJSON(&form)
		ctx := context.WithValue(context.Background(), "UserID", c.GetString("id"))
		ctx = context.WithValue(ctx, "CartID", c.GetString("CartID"))
		ctx = context.WithValue(ctx, "GuestID", c.GetString("GuestID"))
		fmt.Printf("Triplets : %s %s %s\n", ctx.Value("UserID"), ctx.Value("CartID"), ctx.Value("GuestID"))
		res, err := srv.RemoveItemFromCart(ctx, form.ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(200, res)
	}
}

func UpdateCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		err := c.BindJSON(&form)
		ctx := context.WithValue(context.Background(), "UserID", c.GetString("id"))
		ctx = context.WithValue(ctx, "CartID", c.GetString("CartID"))
		ctx = context.WithValue(ctx, "GuestID", c.GetString("GuestID"))
		fmt.Printf("Triplets : %s %s %s\n", ctx.Value("UserID"), ctx.Value("CartID"), ctx.Value("GuestID"))
		res, err := srv.UpdateCart(ctx, form.ProductID, form.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(200, res)
	}
}

func DeleteCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), "UserID", c.GetString("id"))
		ctx = context.WithValue(ctx, "CartID", c.GetString("CartID"))
		ctx = context.WithValue(ctx, "GuestID", c.GetString("GuestID"))
		fmt.Printf("Triplets : %s %s %s\n", ctx.Value("UserID"), ctx.Value("CartID"), ctx.Value("GuestID"))
		err := srv.DeleteCart(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(200, gin.H{"Message": "Cart Succesfully Deleted"})
	}
}
