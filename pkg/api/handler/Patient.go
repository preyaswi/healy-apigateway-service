package handler

import (
	"context"
	"fmt"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type PatientHandler struct {
	Grpc_client interfaces.PatientClient
}

func NewPatientHandler(PatientClient interfaces.PatientClient) *PatientHandler {
	return &PatientHandler{
		Grpc_client: PatientClient,
	}
}

func (p *PatientHandler) PatientSignup(c *fiber.Ctx) error {
	var SignupDetail models.PatientSignUp
	if err := c.BodyParser(&SignupDetail); err != nil {
		errs := response.ClientResponse("Details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(SignupDetail); err != nil {
		fmt.Println(err, "validation error")
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	Patient, err := p.Grpc_client.PatientsSignUp(SignupDetail)
	if err != nil {
		errs := response.ClientResponse("Details not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Patient created successfully", Patient, nil)
	return c.Status(201).JSON(success)
}
func (p *PatientHandler) PatientLogin(c *fiber.Ctx) error {
	var logindetail models.PatientLogin
	if err := c.BodyParser(&logindetail); err != nil {
		errs := response.ClientResponse("logindetails are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	if err := validator.New().Struct(logindetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	patient, err := p.Grpc_client.PatientLogin(logindetail)

	if err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("patient logined succesfully", patient, nil)
	return c.Status(201).JSON(success)
}
func (p *PatientHandler) PatientDetails(c *fiber.Ctx) error {

	patientID := c.Locals("user_id").(int)

	fmt.Println(patientID, "patient id")
	patientDetails, err := p.Grpc_client.PatientDetails(patientID)
	if err != nil {
		errorRes := response.ClientResponse("failed to retrieve patient details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)

	}

	successRes := response.ClientResponse("patient Details", patientDetails, nil)
	return c.Status(200).JSON(successRes)

}
func (p *PatientHandler) UpdatePatientDetails(c *fiber.Ctx) error {

	user_id := c.Locals("user_id").(int)

	var patient models.PatientDetails

	if err := c.BodyParser(&patient); err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(patient); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	updatedDetails, err := p.Grpc_client.UpdatePatientDetails(patient, user_id)
	if err != nil {
		errorRes := response.ClientResponse("failed update user", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}

	successRes := response.ClientResponse("Updated User Details", updatedDetails, nil)
	return c.Status(200).JSON(successRes)

}

type contextKey string

const (
	userIDKey contextKey = "userID"
)

func (p *PatientHandler) UpdatePassword(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, userID)
	var body models.UpdatePassword

	if err := c.BodyParser(&body); err != nil {
		errorRes := response.ClientResponse("fields provided are in wrong format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}

	err := p.Grpc_client.UpdatePassword(ctx, userID, body)

	if err != nil {
		errorRes := response.ClientResponse("failed updating password", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorRes)
	}

	successRes := response.ClientResponse("Password updated successfully", nil, nil)
	return c.Status(http.StatusCreated).JSON(successRes)
}
func (p *PatientHandler) ListPatients(c *fiber.Ctx) error {
	listedPatients, err := p.Grpc_client.ListPatients()
	if err != nil {
		errorRes := response.ClientResponse("failed retreiving list of patients", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorRes)
	}
	successRes := response.ClientResponse("retrived list of patients", listedPatients, nil)
	return c.Status(200).JSON(successRes)
}
