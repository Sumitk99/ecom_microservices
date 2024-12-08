package main

import (
	"github.com/Sumitk99/ecom_microservices/cart"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string
	CatalogURL  string
}

func main() {
	var cfg Config
	cfg.CatalogURL = "localhost:8082"
	if cfg.DatabaseURL == "" {
		log.Fatal("No DATABASE_URL set")
	}
	if cfg.CatalogURL == "" {
		log.Fatal("No CATALOG_SERVICE_URL set")
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
	log.Println("Listening on port 8083")
	log.Println("database url : ", cfg.DatabaseURL)
	s := cart.NewService(r)
	log.Fatal(cart.ListenGRPC(s, cfg.CatalogURL, "8083"))

}
