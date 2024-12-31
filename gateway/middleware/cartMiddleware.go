package middleware

import (
	"context"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CartMiddleware(srv *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("cart middleware")
		clientToken := c.Request.Header.Get("authorization")
		guestToken := c.Request.Header.Get("guestAuth")
		fmt.Println(clientToken)

		if len(clientToken) > 0 {
			ctx := context.WithValue(context.Background(), "authorization", clientToken)
			account, err := srv.Authentication(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			fmt.Println(account)
			c.Set("id", account.Account.Id)
			c.Set("CartID", c.Request.Header.Get("CartID"))
			//c.Set("name", accountClient.Name)
			//c.Set("email", accountClient.Email)
			//c.Set("phone", accountClient.Phone)
			//c.Set("user_type", accountClient.UserType)
			//log.Println("finished authenticating")
		} else if len(guestToken) > 0 {
			ctx := context.WithValue(context.Background(), "guestAuth", guestToken)
			guestId, err := srv.ValidateGuestCartToken(ctx, guestToken)
			fmt.Printf("guestId : %s\n", guestId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.Set("GuestID", guestId)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		log.Println("finished cart authentication")
		log.Printf("cartuser %s\n", c.GetString("id"))
		c.Next()
	}
}
