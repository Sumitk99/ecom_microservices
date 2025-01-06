package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/gateway/routes"
	"github.com/Sumitk99/ecom_microservices/gateway/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type Config struct {
	AccountURL string
	OrderURL   string
	CartURL    string
	CatalogURL string
	PORT       string
}

func main() {

	var router *gin.Engine = gin.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

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
	var cfg Config
	cfg.AccountURL = os.Getenv("ACCOUNT_SERVICE_URL")
	cfg.OrderURL = os.Getenv("ORDER_SERVICE_URL")
	cfg.CartURL = os.Getenv("CART_SERVICE_URL")
	cfg.CatalogURL = os.Getenv("CATALOG_SERVICE_URL")
	cfg.PORT = os.Getenv("PORT")
	fmt.Println(cfg)
	if len(cfg.PORT) == 0 {
		cfg.PORT = "8080"
	}
	srv, err := server.NewGinServer(cfg.AccountURL, cfg.CartURL, cfg.OrderURL, cfg.CatalogURL)
	if err != nil {
		log.Println(err)
	}

	routes.SetupRoutes(router, srv)

	fmt.Println(fmt.Sprintf("Gateway Listening on Port %s\n", cfg.PORT))
	err = router.Run(fmt.Sprintf("0.0.0.0:%s", cfg.PORT))
	if err != nil {
		log.Fatal(err)
	}
}
