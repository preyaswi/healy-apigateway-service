package server

import (
	"healy-apigateway/pkg/api/handler"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(route *fiber.App, adminHandler *handler.AdminHandler) {
	doctor := route.Group("/admin")
	doctor.Post("/signup",adminHandler.AdminSignUp)
	doctor.Post("/login",adminHandler.LoginHandler)

}
