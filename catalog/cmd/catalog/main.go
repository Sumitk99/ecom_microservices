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
	AwsEndpoint string
	AwsRegion   string
	PORT        string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var cfg Config
	cfg.AwsEndpoint = os.Getenv("AWS_ENDPOINT_URL")
	cfg.AwsRegion = os.Getenv("AWS_REGION")
	cfg.PORT = os.Getenv("PORT")
	if len(cfg.PORT) == 0 {
		cfg.PORT = "8080"
	}

	if len(cfg.AwsEndpoint) == 0 || len(cfg.AwsRegion) == 0 {
		log.Fatal("Elastic search cloud id and api key are required")
	}

	var r catalog.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = catalog.NewOpenSearchRepository(cfg.AwsEndpoint, cfg.AwsRegion)
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
