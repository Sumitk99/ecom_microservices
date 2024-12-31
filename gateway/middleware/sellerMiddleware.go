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
		if account.Account.UserType != "SELLER" && account.Account.UserType != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Access"})
			c.Abort()
			return
		}
		c.Set("id", account.Account.Id)
		c.Set("name", account.Account.Name)
		c.Set("email", account.Account.Email)
		c.Set("phone", account.Account.Phone)
		c.Set("user_type", account.Account.UserType)
		log.Println("finished authenticating")
		c.Next()
	}
}
