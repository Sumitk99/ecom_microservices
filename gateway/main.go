package main

import (
	"github.com/Sumitk99/ecom_microservices/gateway/routes"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {

	var router *gin.Engine = gin.New()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	router.Use(gin.Logger())

	accountUrl := os.Getenv("ACCOUNT_SERVICE_URL")
	srv, err := server.NewGinServer(accountUrl)
	if err != nil {
		log.Println(err)
	}
	routes.PublicRoutes(router, srv)
	routes.ProtectedRoutes(router, srv)

	err = router.Run(":8000")
	if err != nil {
		log.Println(err)
	}
}