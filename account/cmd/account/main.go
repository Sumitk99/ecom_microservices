package main

import (
	"github.com/Sumitk99/ecom_microservices/account"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	DatabaseURL string
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}
	var cfg Config
	port := os.Getenv("PORT")
	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseURL == "" {
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
	log.Printf("Account Service Listening on port %s", port)
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, port))
}
