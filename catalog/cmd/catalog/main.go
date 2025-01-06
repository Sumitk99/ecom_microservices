package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	CloudID string `envconfig:"ELASTIC_SEARCH_CLOUD_ID"`
	ApiKey  string `envconfig:"ELASTIC_SEARCH_API_KEY"`
	PORT    string `envconfig:"PORT"`
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	//cfg.cloudID = os.Getenv("ELASTIC_SEARCH_CLOUD_ID")
	//cfg.apiKey = os.Getenv("ELASTIC_SEARCH_API_KEY")
	//cfg.PORT = os.Getenv("PORT")
	fmt.Printf("Error: %s\n", err)
	fmt.Printf("Cloud ID: %s\n", cfg.CloudID)
	fmt.Printf("API Key: %s\n", cfg.ApiKey)
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
