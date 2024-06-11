package main

import (
	"healy-apigateway/pkg/config"
	"healy-apigateway/pkg/di"
	"healy-apigateway/pkg/logging"
)

func main() {
	logging.Init()
	logEntry := logging.Logger().WithField("context", "loadingconfig")
	cfg, cfgerr := config.LoadConfig()
	if cfgerr != nil {
		// log.Fatal("cannot load api gateway config", err)
		logEntry.WithError(cfgerr).Fatal("cannot load api gateway config")
	}
	logEntry = logEntry.WithField("context", "initializingapi")
	logEntry.Info("Initializing API server")
	server, diErr := di.InitializeApi(cfg)
	if diErr != nil {
		logEntry.WithError(diErr).Fatal("cannot start server")
		// log.Fatal("cannot start server: ", diErr)
	} else {
		logEntry = logEntry.WithField("context", "starting server")
		logEntry.Info("starting server")
		server.Start(cfg)
	}
}
