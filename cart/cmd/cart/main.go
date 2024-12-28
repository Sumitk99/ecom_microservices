package main

import (
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

	if len(cfg.DatabaseURL) == 0 {
		log.Fatal("No DATABASE_URL set")
	}
	if len(cfg.CatalogURL) == 0 {
		log.Fatal("No CATALOG_SERVICE_URL set")
	}
	if len(cfg.OrderURL) == 0 {
		log.Fatal("No ORDER_SERVICE_URL set")
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
	log.Println("Cart Service Listening on port 8083")
	s := cart.NewService(r)
	log.Fatal(cart.ListenGRPC(s, cfg.CatalogURL, cfg.OrderURL, "8083"))

}
