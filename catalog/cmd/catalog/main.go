package main

import (
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	CloudID string
	ApiKey  string
	PORT    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config
	cfg.CloudID = os.Getenv("ELASTIC_SEARCH_CLOUD_ID")
	cfg.ApiKey = os.Getenv("ELASTIC_SEARCH_API_KEY")
	cfg.PORT = os.Getenv("PORT")
	if len(cfg.PORT) == 0 {
		cfg.PORT = "8080"
	}

	if len(cfg.CloudID) == 0 || len(cfg.ApiKey) == 0 {
		log.Fatal("Elastic search cloud id and api key are required")
	}

	var r catalog.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.CloudID, cfg.ApiKey)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Printf("Catalog Service Listening on port %s", cfg.PORT)
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, cfg.PORT))
}
