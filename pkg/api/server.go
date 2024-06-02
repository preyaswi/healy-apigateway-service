package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type ServerHTTP struct {
	engine *fiber.App
}

func NewServerHTTP(patientHandler *handler.PatientHandler,doctorHandler *handler.DoctorHandler,adminHandler *handler.AdminHandler,bookingHandler *handler.BookingHandler,paymentHandler *handler.PaymentHandler) *ServerHTTP {
	engine := html.New("./template", ".html")
	route := fiber.New(fiber.Config{
		Views: engine,
	})
	route.Use(logger.New())
	DoctorRoutes(route,doctorHandler,patientHandler)
	PatientRoutes(route,patientHandler,doctorHandler,adminHandler,bookingHandler,paymentHandler)
	AdminRoutes(route,adminHandler,patientHandler,doctorHandler)
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

