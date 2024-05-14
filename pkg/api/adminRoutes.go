package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(route *fiber.App, adminHandler *handler.AdminHandler,patientHandler *handler.PatientHandler,doctorHandler *handler.DoctorHandler) {
	doctor := route.Group("/admin")
	doctor.Post("/signup",adminHandler.AdminSignUp)
	doctor.Post("/login",adminHandler.LoginHandler)
	doctor.Use(middleware.AdminAuthMiddleware())
	{
		dashboard:=doctor.Group("/dashboard")
		patient:=dashboard.Group("/patients")
		{
			patient.Get("",patientHandler.ListPatients)
		}
		doctor:=dashboard.Group("/doctors")
		{
			doctor.Get("",doctorHandler.DoctorsDetails)
		}
	}

}
