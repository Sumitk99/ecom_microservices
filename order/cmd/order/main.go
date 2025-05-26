package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/order"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL  string `envconfig:"CATALOG_SERVICE_URL"`
	PORT        string `envconfig:"PORT"`
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
	cfg.PORT = os.Getenv("PORT")

	if len(cfg.PORT) == 0 {
		cfg.PORT = "8083"
	}
	var r order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("Postgres error : ", err)
		}
		return
	})

	defer r.Close()

	log.Println(fmt.Sprintf("Order Service Listening on port %s", cfg.PORT))
	s := order.NewService(r)

	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, cfg.PORT))
}
