package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/routes"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {

	var router *gin.Engine = gin.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	config := cors.Config{
		AllowOrigins:     []string{"http://192.168.196.240:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	accountUrl := os.Getenv("ACCOUNT_SERVICE_URL")
	orderUrl := os.Getenv("ORDER_SERVICE_URL")
	cartUrl := os.Getenv("CART_SERVICE_URL")
	srv, err := server.NewGinServer(accountUrl, cartUrl, orderUrl)
	if err != nil {
		log.Println(err)
	}
	routes.PublicRoutes(router, srv)
	routes.ProtectedRoutes(router, srv)
	fmt.Println("Gateway Listening on Port 8000")
	err = router.Run(":8000")
	if err != nil {
		log.Println(err)
	}
}
