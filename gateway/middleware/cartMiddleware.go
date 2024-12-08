package middleware

import (
	"context"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CartMiddleware(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		clientToken := c.Request.Header.Get("authorization")
		guestToken := c.Request.Header.Get("guestAuth")

		if len(clientToken) > 0 {
			ctx := context.WithValue(context.Background(), "authorization", clientToken)

			accountClient, err := srv.AccountClient.Authentication(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.Abort()
				return
			}

			c.Set("id", accountClient.ID)
			//c.Set("name", accountClient.Name)
			//c.Set("email", accountClient.Email)
			//c.Set("phone", accountClient.Phone)
			//c.Set("user_type", accountClient.UserType)
			//log.Println("finished authenticating")
		} else if len(guestToken) > 0 {
			ctx := context.WithValue(context.Background(), "guestAuth", guestToken)
			guestId, err := srv.ValidateGuestCartToken(ctx, guestToken)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.Abort()
				return
			}

			c.Set("GuestID", guestId)
			log.Println("finished cart authentication")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
