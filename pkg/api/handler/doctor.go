package handler

import (
	"fmt"
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
func (d *DoctorHandler)DoctorsDetails(c *fiber.Ctx)error  {
	fmt.Println("hello")
	doctor,err:=d.Grpc_Client.DoctorsDetails()
	if err!=nil{
		errs := response.ClientResponse("couldnt fetch doctors data", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctors data fetched succesfully", doctor, nil)
	return c.Status(201).JSON(success)
}
func (d *DoctorHandler)IndividualDoctor(c *fiber.Ctx)error  {
	doctorID:=c.Params("doctor_id")
	fmt.Println(doctorID,"this is th doctor id ")
	doctor,err:=d.Grpc_Client.IndividualDoctor(doctorID)
	if err!=nil{
		errs:=response.ClientResponse("couldn't fetch octors data",nil,err.Error())
		return c.Status(400).JSON(errs)
	}
	success:=response.ClientResponse("returned individual doctor data",doctor,nil)
	return c.Status(201).JSON(success)
}