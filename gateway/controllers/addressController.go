package controller

import (
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAddresses(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := CreateUserMetaData(c)
		addresses, err := srv.GetAddresses(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, addresses)

	}
}

func AddAddress(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := CreateUserMetaData(c)
		var form models.AddAddressRequest
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		address, err := srv.AddAddress(ctx, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, address)
	}
}

func DeleteAddress(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := CreateUserMetaData(c)
		addressId := c.Param("address_id")
		err := srv.DeleteAddress(ctx, addressId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
	}
}

func GetAddress(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := CreateUserMetaData(c)
		addressId := c.Param("id")
		address, err := srv.GetAddress(ctx, addressId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Address": address})
	}
}
