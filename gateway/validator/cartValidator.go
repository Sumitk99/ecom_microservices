package validator

import (
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func ValidateCartOpsReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		_ = c.BindJSON(&form)
		validationErr := validator.New().Struct(form)
		if validationErr != nil {
			log.Println(validationErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ValidateRemoveFromCartReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.CartOpsReq
		_ = c.BindJSON(&form)
		if len(form.ProductID) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
