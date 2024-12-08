package main

import (
	"github.com/Sumitk99/ecom_microservices/account"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string
}

func main() {
	var cfg Config
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
	log.Println("Listening on port 8080")
	log.Println("database url : ", cfg.DatabaseURL)
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, "8080"))
}
