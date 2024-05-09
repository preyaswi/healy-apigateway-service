package server

import (
	"healy-apigateway/pkg/api/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ServerHTTP struct {
	engine *fiber.App
}

func NewServerHTTP(patientHandler *handler.PatientHandler,doctorHandler *handler.DoctorHandler,adminHandler *handler.AdminHandler) *ServerHTTP {
	route := fiber.New()
	route.Use(logger.New())
	DoctorRoutes(route,doctorHandler)
	PatientRoutes(route,patientHandler)
	AdminRoutes(route,adminHandler)

	
	// patient.Use(middleware.UserAuthMiddleware())
	// {
	// 	patient.Get("/doctor",)
	// }

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

