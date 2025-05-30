package middleware

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authenticating")
		clientToken := c.Request.Header.Get("authorization")
		if len(clientToken) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required, please login"})
			c.Abort()
			return
		}
		ctx := context.WithValue(context.Background(), "authorization", clientToken)

		accountClient, err := srv.Authentication(ctx)
		if err != nil {
			log.Println("Error authenticating: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", accountClient.Account.Id)
		c.Set("name", accountClient.Account.Name)
		c.Set("email", accountClient.Account.Email)
		c.Set("phone", accountClient.Account.Phone)
		c.Set("user_type", accountClient.Account.UserType)
		log.Println("finished authenticating")
		c.Next()
	}
}
