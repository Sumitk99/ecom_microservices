package middleware

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SellerMiddleware(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authenticating Seller")
		clientToken := c.Request.Header.Get("authorization")
		if len(clientToken) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required, please login"})
			c.Abort()
			return
		}
		ctx := context.WithValue(context.Background(), "authorization", clientToken)

		account, err := srv.Authentication(ctx)
		if err != nil {
			log.Println("Error authenticating: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if account.UserType != "SELLER" && account.UserType != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Access"})
			c.Abort()
			return
		}
		c.Set("id", account.ID)
		c.Set("name", account.Name)
		c.Set("email", account.Email)
		c.Set("phone", account.Phone)
		c.Set("user_type", account.UserType)
		log.Println("finished authenticating")
		c.Next()
	}
}
