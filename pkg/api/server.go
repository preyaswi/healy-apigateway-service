package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/config"
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
	PatientRoutes(route,patientHandler,doctorHandler)
	AdminRoutes(route,adminHandler,patientHandler,doctorHandler)

	
	// patient.Use(middleware.UserAuthMiddleware())
	// {
	// 	patient.Get("/doctor",)
	// }

	return &ServerHTTP{
		engine: route,
	}
}
func (s *ServerHTTP) Start(cfg config.Config) {
	
    log.Printf("starting server on :5080")
    err := s.engine.Listen(cfg.Port)
    if err != nil {
        log.Fatalf("error while starting the server: %v", err)
    }
}

