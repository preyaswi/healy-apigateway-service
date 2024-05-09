package di

import (
	server "healy-apigateway/pkg/api"
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/client"
	"healy-apigateway/pkg/config"
)

func InitializeApi(cfg config.Config) (*server.ServerHTTP, error) {
	patientClient := client.NewPatientClient(cfg)
	patientHandler := handler.NewPatientHandler(patientClient)

	doctorClient:=client.NewDoctorClient(cfg)
	doctorHandler:=handler.NewDoctorHandler(doctorClient)

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	serverHTTP := server.NewServerHTTP(patientHandler,doctorHandler,adminHandler)
	return serverHTTP, nil
}
