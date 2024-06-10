package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	Grpc_Client interfaces.AdminClient
}

func NewPaymentHandler(paymentClient interfaces.AdminClient) *PaymentHandler {
	return &PaymentHandler{
		Grpc_Client: paymentClient,
	}
}
func (pa *PaymentHandler) MakePaymentRazorpay(c *fiber.Ctx) error {
	bookingid := c.Query("booking_id")
	booking_id, err := strconv.Atoi(bookingid)
	if err != nil {
		errs := response.ClientResponse("cannot convert booking id string to int", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	paymentDetails, razorId, err := pa.Grpc_Client.MakePaymentRazorpay(booking_id)
	if err != nil {
		errs := response.ClientResponse("couldn't make payment Details", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"final_price": paymentDetails.Fees * 100,
		"razor_id":    razorId,
		"user_id":     paymentDetails.PatientId,
		"order_id":    paymentDetails.BookingId,
		"user_email":   paymentDetails.DoctorEmail,
		"total":       int(paymentDetails.Fees),
	})

}
func (pa *PaymentHandler) VerifyPayment(c *fiber.Ctx)error  {
	bookingId,err:=strconv.Atoi(c.Query("booking_id"))
	if err!=nil{
		errs := response.ClientResponse("cannot convert id string to int", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	err=pa.Grpc_Client.VerifyPayment(bookingId)
	if err!=nil{
		errs := response.ClientResponse("couldn't update payment Details", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errs)
	}

	success := response.ClientResponse("Successfully updated payment details", nil, nil)
	return c.Status(http.StatusOK).JSON(success)
}