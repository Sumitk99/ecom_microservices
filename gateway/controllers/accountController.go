package controller

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUser(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("id")
		log.Println("User ID: ", userId)
		ctx := context.WithValue(context.Background(), "UserID", userId)
		user, err := srv.GetAccount(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(user)
		c.JSON(http.StatusOK, user)
	}
}

func SignUp(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var form models.SignUpRequest
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := srv.SignUp(ctx, form.Name, form.Password, form.Email, form.Phone, form.UserType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func Login(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var form models.LoginRequest
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, token, refreshToken, err := srv.Login(ctx, form.Email, form.Phone, form.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := models.LoginResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			Phone:         user.Phone,
			UserType:      user.UserType,
			JWT_Token:     token,
			Refresh_Token: refreshToken,
		}
		c.JSON(http.StatusOK, res)
	}
}
