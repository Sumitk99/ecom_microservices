package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/routes"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {

	var router *gin.Engine = gin.New()
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file %s", err)
	//}

	corsPolicy := cors.Config{
		AllowOrigins:     []string{"http://192.168.205.239:4200", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsPolicy))
	router.Use(gin.Logger())

	accountUrl := os.Getenv("ACCOUNT_SERVICE_URL")
	orderUrl := os.Getenv("ORDER_SERVICE_URL")
	cartUrl := os.Getenv("CART_SERVICE_URL")
	catalogUrl := os.Getenv("CATALOG_SERVICE_URL")
	PORT := os.Getenv("PORT")
	srv, err := server.NewGinServer(accountUrl, cartUrl, orderUrl, catalogUrl)
	if err != nil {
		log.Println(err)
	}

	routes.SetupRoutes(router, srv)

	fmt.Println("Gateway Listening on Port 8000")
	err = router.Run(fmt.Sprintf("0.0.0.0:%s", PORT))
	if err != nil {
		log.Println(err)
	}
}
