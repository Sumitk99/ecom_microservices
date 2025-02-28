package controller

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strconv"
)

const AlreadyExists = "rpc error: code = Unknown desc = cannot add item to cart : Item already exists in selected cart"

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
		cartName := c.Param("cart_name")
		productId := c.Param("product_id")
		quantity := c.Param("quantity")
		if len(productId) == 0 || len(quantity) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID and Quantity are required"})
			return
		}
		if len(cartName) == 0 {
			cartName = c.GetString("id")
		}
		c.Set("CartID", cartName)
		ctx := GetCartContext(c)
		Quantity, err := strconv.Atoi(quantity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Quantity"})
			return
		}
		res, err := srv.AddItemToCart(ctx, productId, uint64(Quantity))
		if err != nil {
			state, _ := status.FromError(err)
			if state.Code() == codes.AlreadyExists {
				c.JSON(http.StatusConflict, gin.H{"error": AlreadyExists})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)
	}
}

func GetCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Get Cart")

		cartId := c.Param("cart_id")
		if len(cartId) == 0 {
			cartId = c.GetString("id")
		}
		log.Println(cartId)
		c.Set("CartID", cartId)
		ctx := GetCartContext(c)

		res, err := srv.GetCart(ctx)
		log.Println(res, err)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Response Received")
		c.JSON(http.StatusOK, res)
	}
}

func RemoveItemFromCart(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId := c.Param("cart_id")
		productId := c.Param("product_id")
		if len(cartId) == 0 {
			cartId = c.GetString("id")
		}
		if len(productId) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID is required"})
			return
		}
		c.Set("CartID", cartId)
		ctx := GetCartContext(c)
		res, err := srv.RemoveItemFromCart(ctx, productId)
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
		cartId := c.Param("cart_id")
		productId := c.Param("product_id")
		quantity := (c.Param("quantity"))
		if len(cartId) == 0 {
			cartId = c.GetString("id")
		}
		if len(productId) == 0 || len(quantity) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID and Quantity are required"})
			return
		}
		c.Set("CartID", cartId)
		quantityInt, err := strconv.Atoi(quantity)
		ctx := GetCartContext(c)
		res, err := srv.UpdateCart(ctx, productId, uint64(quantityInt))
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
		cartId := c.Param("cart_id")
		if len(cartId) == 0 {
			cartId = c.GetString("id")
		}
		c.Set("CartID", cartId)
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
		var form models.CheckoutRequest
		err := c.ShouldBindJSON(&form)
		cartId := c.Param("cart_id")
		if len(cartId) == 0 {
			cartId = c.GetString("id")
		}
		c.Set("CartID", cartId)
		ctx := GetCartContext(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("1.%s 2.%s 3.%s\n", cartId, form.MethodOfPayment, form.TransactionID)
		res, err := srv.Checkout(ctx, cartId, form.MethodOfPayment, form.TransactionID, form.AddressId)
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
