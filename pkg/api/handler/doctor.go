package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"strconv"

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
func (d *DoctorHandler) DoctorsDetails(c *fiber.Ctx) error {
	doctor, err := d.Grpc_Client.DoctorsDetails()
	if err != nil {
		errs := response.ClientResponse("couldnt fetch doctors data", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctors data fetched succesfully", doctor, nil)
	return c.Status(201).JSON(success)
}
func (d *DoctorHandler) IndividualDoctor(c *fiber.Ctx) error {
	doctorID := c.Params("doctor_id")
	doctor, err := d.Grpc_Client.IndividualDoctor(doctorID)
	if err != nil {
        errorRes := response.ClientResponse("couldn't fetch doctors data", nil, err.Error())
        return c.Status(http.StatusBadRequest).JSON(errorRes)
    }
	success := response.ClientResponse("returned individual doctor data", doctor, nil)
	return c.Status(201).JSON(success)
}
func (d *DoctorHandler) DoctorProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	userID,err:=strconv.Atoi(userId)
	if err!=nil{
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	doctorDetail, err := d.Grpc_Client.DoctorProfile(userID)
	if err != nil {
		errorRes := response.ClientResponse("failed to retrieve doctor details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)

	}
	successRes := response.ClientResponse("doctor Details", doctorDetail, nil)
	return c.Status(200).JSON(successRes)

}
func (d *DoctorHandler) RateDoctor(c *fiber.Ctx) error {
	patientid := c.Locals("user_id").(string)
	doctor_id := c.Params("doctor_id")
	var rate models.Rate
	if err := c.BodyParser(&rate); err != nil {
		errs := response.ClientResponse("Details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(rate); err != nil {
		errs := response.ClientResponse("give rate between 1 to 5", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	rated, err := d.Grpc_Client.RateDoctor(patientid, doctor_id, rate)
	if err != nil {
		errs := response.ClientResponse("couldn't add the rating", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	successRes := response.ClientResponse("doctor Details", rated, nil)
	return c.Status(202).JSON(successRes)

}
func (d *DoctorHandler) UpdateDoctorProfile(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId,err:=strconv.Atoi(doctorid)
	if err!=nil{
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	var doctor models.DoctorDetails

	if err := c.BodyParser(&doctor); err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(doctor); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	res, err := d.Grpc_Client.UpdateDoctorProfile(doctorId, doctor)

	if err != nil {
		errorRes := response.ClientResponse("failed update doctor", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("Updated doctor Details", res, nil)
	return c.Status(200).JSON(successRes)
}
