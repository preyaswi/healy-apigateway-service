package server

import (
	"healy-apigateway/pkg/api/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ServerHTTP struct {
	engine *fiber.App
}

func NewServerHTTP(patientHandler *handler.PatientHandler) *ServerHTTP {
	route := fiber.New()
	patient:=route.Group("/patient")

	patient.Post("/signup", patientHandler.PatientSignup)
	patient.Post("/login",patientHandler.PatientLogin)
	return &ServerHTTP{
		engine: route,
	}
}
func (s *ServerHTTP) Start() {
    log.Printf("starting server on :5000")
    err := s.engine.Listen(":5000")
    if err != nil {
        log.Fatalf("error while starting the server: %v", err)
    }
}

