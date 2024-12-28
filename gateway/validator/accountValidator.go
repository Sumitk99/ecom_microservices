package validator

import (
	"github.com/Sumitk99/ecom_microservices/gateway/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.LoginRequest
		_ = c.BindJSON(&form)
		if len(form.Phone) == 0 && len(form.Email) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or phone is required"})
			c.Abort()
			return
		}
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

func SignUpValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form models.SignUpRequest
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
