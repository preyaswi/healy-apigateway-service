package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(route *fiber.App, patientHandler *handler.PatientHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler,bookingHandler *handler.BookingHandler,paymentHandler *handler.PaymentHandler) {
	patient := route.Group("/patient")
	patient.Post("/signup", patientHandler.PatientSignup)
	patient.Post("/login", patientHandler.PatientLogin)
	patient.Get("/payment", paymentHandler.MakePaymentRazorpay)
	patient.Get("/verifypayment",paymentHandler.VerifyPayment)

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
		booking:=patient.Group("/booking")
		{
			booking.Post("",bookingHandler.AddToBookings)
			booking.Delete("",bookingHandler.CancelBookings)
		}
		

	}
}
