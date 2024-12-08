package main

import (
	"fmt"
	"github.com/Sumitk99/ecom_microservices/catalog"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	cloudID string
	apiKey  string
}

func main() {
	var cfg Config
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
