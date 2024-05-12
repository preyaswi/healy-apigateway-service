package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(route *fiber.App,patientHandler *handler.PatientHandler,doctorHandler *handler.DoctorHandler)  {
	patient:=route.Group("/patient")
	patient.Post("/signup", patientHandler.PatientSignup)
	patient.Post("/login",patientHandler.PatientLogin)
	patient.Use(middleware.UserAuthMiddleware())
	{
		patient.Get("/profile",patientHandler.PatientDetails)

		doctor:=patient.Group("/doctor")
		{
			doctor.Get("",doctorHandler.DoctorsDetails)
			doctor.Get("/:doctor_id",doctorHandler.IndividualDoctor)
		}
		
	}
}