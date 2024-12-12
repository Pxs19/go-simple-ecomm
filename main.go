package main

import (
	"log"

	"ecom2.5/config"
	"ecom2.5/internal/api"
)

func main() {

	cfg, err := config.SetUpEnv()

	if err != nil {
		log.Fatal("config file is not loaded properly %v\n", err)
	}

	// start server
	api.StartServer(cfg)

}
