package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)
type PatientHandler struct{
	Grpc_client interfaces.PatientClient
}
func NewPatientHandler(PatientClient interfaces.PatientClient) *PatientHandler {
	return &PatientHandler{
		Grpc_client: PatientClient,
	}
}

func (p *PatientHandler)PatientSignup(c *fiber.Ctx)error  {
	var SignupDetail models.PatientSignUp
	if err:=c.BodyParser(&SignupDetail);err!=nil{
		errs:=response.ClientResponse("Details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	if err := validator.New().Struct(SignupDetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	Patient,err:=p.Grpc_client.PatientsSignUp(SignupDetail)
	if err!=nil{
		errs := response.ClientResponse("Details not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success:=response.ClientResponse("Patient created successfully",Patient,nil)
	return c.Status(201).JSON(success)
}
func(p *PatientHandler) PatientLogin(c *fiber.Ctx) error {
	var logindetail models.PatientLogin
	if err:=c.BodyParser(&logindetail);err!=nil{
		errs:=response.ClientResponse("logindetails are not in correct format",nil,err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	
	if err := validator.New().Struct(logindetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	patient,err:=p.Grpc_client.PatientLogin(logindetail)
	

	if err!=nil{
		errs:=response.ClientResponse("details are not in correct format",nil,err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success:=response.ClientResponse("patient logined succesfully",patient,nil)
	return c.Status(201).JSON(success)
}