package main

import (
	"healy-apigateway/pkg/config"
	"healy-apigateway/pkg/di"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load api gateway config", err)
	}
	
	server, diErr := di.InitializeApi(cfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start(cfg)
	}
}
