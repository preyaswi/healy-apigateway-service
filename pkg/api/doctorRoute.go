package server

import (
	"healy-apigateway/pkg/api/handler"

	"github.com/gofiber/fiber/v2"
)

func DoctorRoutes(route *fiber.App, doctorHandler *handler.DoctorHandler) {
	doctor := route.Group("/doctor")
	doctor.Post("/signup",doctorHandler.DoctorSignUp)
	doctor.Post("/login",doctorHandler.DoctorLogin)

}
