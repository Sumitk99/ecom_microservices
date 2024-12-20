package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
	"log"
	"os"
	"time"
)

type Config struct {
	cloudID string
	apiKey  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config

	cfg.cloudID = os.Getenv("ELASTIC_SEARCH_CLOUD_ID")
	cfg.apiKey = os.Getenv("ELASTIC_SEARCH_API_KEY")
	if len(cfg.cloudID) == 0 || len(cfg.apiKey) == 0 {
		log.Fatal("Elastic search cloud id and api key are required")
	}
	fmt.Println("cfg.apiKey : ", cfg.apiKey)
	fmt.Println("cfg.cloudID : ", cfg.cloudID)

	var r catalog.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.cloudID, cfg.apiKey)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8082")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, "8082"))
}
