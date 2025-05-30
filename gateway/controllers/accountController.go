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
		res, err := srv.Login(ctx, form.Email, form.Phone, form.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie(
			"jwt_token",            // Name of the cookie
			res.JWT_Token,          // Value of the cookie
			36000,                  // MaxAge in seconds (10 hour)
			"/",                    // Path
			"192.168.205.240:3000", // Domain (empty means any domain)
			false,                  // Secure (true if HTTPS)
			true,                   // HttpOnly (true if not accessible via JavaScript)
		)
		c.SetSameSite(http.SameSiteNoneMode)

		c.JSON(http.StatusOK, res)
	}
}
