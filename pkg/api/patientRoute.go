package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PatientRoutes(route *fiber.App, patientHandler *handler.PatientHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler, bookingHandler *handler.BookingHandler, paymentHandler *handler.PaymentHandler) {
    route.Get("/google/redirect", middleware.LoggingMiddleware(patientHandler.GoogleCallback))
    patient := route.Group("/patient")
   
    patient.Get("login", middleware.LoggingMiddleware(patientHandler.GoogleLogin))
    patient.Get("/payment", middleware.LoggingMiddleware(paymentHandler.MakePaymentRazorpay))
    patient.Get("/verifypayment", middleware.LoggingMiddleware(paymentHandler.VerifyPayment))

	patient.Use(middleware.UserAuthMiddleware())
    profile := patient.Group("/profile")
    profile.Get("", middleware.LoggingMiddleware(patientHandler.PatientDetails))
    profile.Put("", middleware.LoggingMiddleware(patientHandler.UpdatePatientDetails))

    doctor := patient.Group("/doctor")
    doctor.Get("availability",middleware.LoggingMiddleware(bookingHandler.GetDoctorSlotAvailability))
    doctor.Get("", middleware.LoggingMiddleware(doctorHandler.DoctorsDetails))
    doctor.Get("/:doctor_id", middleware.LoggingMiddleware(doctorHandler.IndividualDoctor))
    doctor.Post("/rate/:doctor_id", middleware.LoggingMiddleware(doctorHandler.RateDoctor))
    

    booking := patient.Group("/booking")
    booking.Post("", middleware.LoggingMiddleware(bookingHandler.AddToBookings))
    booking.Delete("", middleware.LoggingMiddleware(bookingHandler.CancelBookings))
    booking.Post("/bookslot",middleware.LoggingMiddleware(bookingHandler.BookSlot))
}
