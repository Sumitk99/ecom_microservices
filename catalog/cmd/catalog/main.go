package main

import (
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	cloudID string `envconfig:"ELASTIC_SEARCH_CLOUD_ID"`
	apiKey  string `envconfig:"ELASTIC_SEARCH_API_KEY"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r catalog.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.cloudID, cfg.apiKey)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8080")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, "8080"))
}
