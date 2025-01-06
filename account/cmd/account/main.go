package main

import (
	"github.com/Sumitk99/ecom_microservices/account"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	DatabaseURL string
	PORT        string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config
	cfg.PORT = os.Getenv("PORT")
	if len(cfg.PORT) == 0 {
		cfg.PORT = "8080"
	}
	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if len(cfg.DatabaseURL) == 0 {
		log.Fatal("No DATABASE_URL set")
	}
	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Printf("Account Service Listening on port %s", cfg.PORT)
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, cfg.PORT))
}
