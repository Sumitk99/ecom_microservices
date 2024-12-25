package controller

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func RequestGuestId(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		guestId, err := srv.IssueGuestCartToken(ctx)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"guestId": guestId})
	}
}

func AddItemToCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		err := c.BindJSON(&form)
		ctx := GetCartContext(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		res, err := srv.AddItemToCart(ctx, form.ProductID, form.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)
	}
}

func GetCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := GetCartContext(c)
		res, err := srv.GetCart(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)

	}
}

func RemoveItemFromCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		err := c.BindJSON(&form)
		ctx := GetCartContext(c)
		res, err := srv.RemoveItemFromCart(ctx, form.ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)
	}
}

func UpdateCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		err := c.BindJSON(&form)
		ctx := GetCartContext(c)
		res, err := srv.UpdateCart(ctx, form.ProductID, form.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)
	}
}

func DeleteCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := GetCartContext(c)
		err := srv.DeleteCart(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, gin.H{"Message": "Cart Successfully Deleted"})
	}
}

func Checkout(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("CartID", c.Request.Header.Get("CartID"))
		ctx := GetCartContext(c)

		var form models.CheckoutRequest
		form.CartID = c.Request.Header.Get("CartID")
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		res, err := srv.Checkout(ctx, form.CartID, form.MethodOfPayment, form.TransactionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)
	}
}

func GetCartContext(c *gin.Context) context.Context {
	md := metadata.New(map[string]string{
		"UserID":  c.GetString("id"),
		"CartID":  c.GetString("CartID"),
		"GuestID": c.GetString("GuestID"),
	})
	// NewOutgoingContext uses mdOutgoingKey{} as  Key. It creates a new context with outgoing md attached. If used
	// in conjunction with AppendToOutgoingContext, NewOutgoingContext will
	// overwrite any previously-appended metadata. md must not be modified after
	// calling this function.
	// To avoid this, use NewIncomingContext which uses mdIncomingKey{} as key
	// if md is possibly modified later
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}
