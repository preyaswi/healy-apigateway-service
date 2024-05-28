package handler

import (
	"fmt"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	GRPC_Client  interfaces.AdminClient
	DoctorClient interfaces.DoctorClient
}

func NewAdminHandler(adminClient interfaces.AdminClient,doctorClient interfaces.DoctorClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: adminClient,
		DoctorClient: doctorClient ,
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
func (ad *AdminHandler) MakePaymentRazorpay(c *fiber.Ctx) error {
	patientId := c.Params("user_id")
	doctorid := c.Params("doctor_id")
	doctor_id, _ := strconv.Atoi(doctorid)
	patient_id, err := strconv.Atoi(patientId)
	if err != nil {
		errs := response.ClientResponse("cannot convert id string to int", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}
	doctorDetail, err := ad.DoctorClient.DoctorDetailforPayment(doctor_id)
	if err != nil {
		errs := response.ClientResponse("couldn't fetch doctor Details", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	paymentDetails, razorId, err := ad.GRPC_Client.MakePaymentRazorpay(patient_id, doctorDetail)
	if err != nil {
		errs := response.ClientResponse("couldn't make payment Details", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}
	fmt.Println(paymentDetails)
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"final_price": paymentDetails.Fees * 100,
		"razor_id":    razorId,
		"user_id":     paymentDetails.PatientId,
		"order_id":    paymentDetails.PaymentId,
		"user_name":   paymentDetails.DoctorName,
		"total":       int(paymentDetails.Fees),
	})
	
	
}
func (ad *AdminHandler) VerifyPayment(c *fiber.Ctx)error  {
	payment_id:=c.Params("payment_id")
	paymentId,err:=strconv.Atoi(payment_id)
	if err!=nil{
		errs := response.ClientResponse("cannot convert id string to int", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	err=ad.GRPC_Client.VerifyPayment(paymentId)
	if err!=nil{
		errs := response.ClientResponse("couldn't update payment Details", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	success := response.ClientResponse("Successfully updated payment details", nil, nil)
	return c.Status(http.StatusOK).JSON(success)
}