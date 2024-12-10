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
		ctx := context.WithValue(context.Background(), "authorization", clientToken)

		accountClient, err := srv.Authentication(ctx)
		if err != nil {
			log.Println("Error authenticating: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("id", accountClient.ID)
		c.Set("name", accountClient.Name)
		c.Set("email", accountClient.Email)
		c.Set("phone", accountClient.Phone)
		c.Set("user_type", accountClient.UserType)
		log.Println("finished authenticating")
		c.Next()
	}
}
