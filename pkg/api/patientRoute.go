package server

import (
	"healy-apigateway/pkg/api/handler"

	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(route *fiber.App,patientHandler *handler.PatientHandler,doctorHandler *handler.DoctorHandler)  {
	patient:=route.Group("/patient")
	patient.Post("/signup", patientHandler.PatientSignup)
	patient.Post("/login",patientHandler.PatientLogin)
	patient.Get("/doctors",doctorHandler.DoctorsDetails)
}