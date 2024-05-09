package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type DoctorHandler struct {
	Grpc_Client interfaces.DoctorClient
}

func NewDoctorHandler(DoctorClient interfaces.DoctorClient) *DoctorHandler {
	return &DoctorHandler{
		Grpc_Client: DoctorClient,
	}
}

func (d *DoctorHandler) DoctorSignUp(c *fiber.Ctx) error {
	var SignupDetail models.DoctorSignUp
	if err := c.BodyParser(&SignupDetail); err != nil {
		errs := response.ClientResponse("Details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	if err := validator.New().Struct(SignupDetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	doctor, err := d.Grpc_Client.DoctorSignUp(SignupDetail)
	if err != nil {
		errs := response.ClientResponse("Details not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctor created successfully", doctor, nil)
	return c.Status(201).JSON(success)
}
func (d *DoctorHandler) DoctorLogin(c *fiber.Ctx) error {
	var logindetail models.DoctorLogin
	if err := c.BodyParser(&logindetail); err != nil {
		errs := response.ClientResponse("logindetails are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	if err := validator.New().Struct(logindetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	doctor, err := d.Grpc_Client.DoctorLogin(logindetail)

	if err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctor logined succesfully", doctor, nil)
	return c.Status(201).JSON(success)
}