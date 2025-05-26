package controller

import (
	"log"
	"net/http"

	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
)

func AddCard(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.AddCardRequest
		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		form.UserID = c.GetString("id")
		log.Println("userId: ", form.UserID)
		res, err := srv.AddCard(&form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func RemoveCard(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.RemoveCardRequest
		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		form.UserID = c.GetString("id")
		log.Println("userId: ", form.UserID)

		res, err := srv.RemoveCard(&form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func ListCards(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Listing Cards")
		userID := c.GetString("id")
		res, err := srv.ListCards(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func ProcessPayment(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.ProcessPaymentRequest
		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		form.UserID = c.GetString("id")
		log.Println("userId: ", form.UserID)

		res, err := srv.ProcessPayment(&form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
