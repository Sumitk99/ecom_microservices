package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/cart"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	DatabaseURL string
	CatalogURL  string
	PORT        string
	OrderURL    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	cfg.CatalogURL = os.Getenv("CATALOG_SERVICE_URL")
	cfg.OrderURL = os.Getenv("ORDER_SERVICE_URL")
	cfg.PORT = os.Getenv("PORT")

	if len(cfg.DatabaseURL) == 0 {
		log.Fatal("No DATABASE_URL set")
	}
	if len(cfg.CatalogURL) == 0 {
		log.Fatal("No CATALOG_SERVICE_URL set")
	}
	if len(cfg.OrderURL) == 0 {
		log.Fatal("No ORDER_SERVICE_URL set")
	}
	if len(cfg.PORT) == 0 {
		cfg.PORT = "8081"
	}

	var r cart.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = cart.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println(fmt.Sprintf("Cart Service Listening on port %s\n", cfg.PORT))
	s := cart.NewService(r)
	log.Fatal(cart.ListenGRPC(s, cfg.CatalogURL, cfg.OrderURL, cfg.PORT))
}
