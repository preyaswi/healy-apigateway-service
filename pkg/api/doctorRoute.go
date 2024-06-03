package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func DoctorRoutes(route *fiber.App, doctorHandler *handler.DoctorHandler, patientHandler *handler.PatientHandler,bookinghandler *handler.BookingHandler) {
	doctor := route.Group("/doctor")
	doctor.Post("/signup", doctorHandler.DoctorSignUp)
	doctor.Post("/login", doctorHandler.DoctorLogin)
	doctor.Use(middleware.UserAuthMiddleware())
	{
		profile := doctor.Group("/profile")
		profile.Get("", doctorHandler.DoctorProfile)
		profile.Put("", doctorHandler.UpdateDoctorProfile)

	}
	patient := doctor.Group("/patient")
	{
		patient.Get("",bookinghandler.GetBookedPatients)
		patient.Post("/prescription/:patient_id", patientHandler.CreatePrescription)
	}

}
