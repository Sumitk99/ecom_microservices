package main

import (
	"github.com/Sumitk99/ecom_microservices/order"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	DatabaseURL string
	AccountURL  string
	CatalogURL  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	cfg.CatalogURL = os.Getenv("CATALOG_SERVICE_URL")
	cfg.AccountURL = os.Getenv("ACCOUNT_SERVICE_URL")
	var r order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("Postgres error : ", err)
		}
		return
	})

	defer r.Close()

	log.Println("Listening on port 8085")
	s := order.NewService(r)

	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, 8085))
}
