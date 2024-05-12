package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(adminClient interfaces.AdminClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: adminClient,
	}
}

func (ad *AdminHandler) LoginHandler(c *fiber.Ctx) error {
	var adminDetails models.AdminLogin
	if err := c.BodyParser(&adminDetails); err != nil {
		errs := response.ClientResponse("Details not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(adminDetails); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		errs := response.ClientResponse("Cannot authenticate user", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	success := response.ClientResponse("Admin authenticated successfully", admin, nil)
	return c.Status(http.StatusOK).JSON(success)
}

func (ad *AdminHandler) AdminSignUp(c *fiber.Ctx) error {
	var adminDetails models.AdminSignUp
	if err := c.BodyParser(&adminDetails); err != nil {
		errs := response.ClientResponse("Details not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	admin, err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil {
		errs := response.ClientResponse("Cannot authenticate user", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	success := response.ClientResponse("Admin created successfully", admin, nil)
	return c.Status(http.StatusOK).JSON(success)
}
