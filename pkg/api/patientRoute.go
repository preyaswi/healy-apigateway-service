package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(route *fiber.App, patientHandler *handler.PatientHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler) {
	patient := route.Group("/patient")
	patient.Post("/signup", patientHandler.PatientSignup)
	patient.Post("/login", patientHandler.PatientLogin)
	patient.Get("/payment/:user_id/:doctor_id", adminHandler.MakePaymentRazorpay)
	patient.Get("/verifypayment/:payment_id",adminHandler.VerifyPayment)

	patient.Use(middleware.UserAuthMiddleware())
	{
		profile := patient.Group("/profile")
		profile.Get("", patientHandler.PatientDetails)
		profile.Put("", patientHandler.UpdatePatientDetails)
		profile.Put("/update-password", patientHandler.UpdatePassword)

		doctor := patient.Group("/doctor")
		{
			doctor.Get("", doctorHandler.DoctorsDetails)
			doctor.Get("/:doctor_id", doctorHandler.IndividualDoctor)
			doctor.Post("/rate/:doctor_id", doctorHandler.RateDoctor)

		}

	}
}
