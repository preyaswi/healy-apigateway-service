package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"strconv"

	"github.com/gofiber/fiber/v2"
)
type BookingHandler struct{
	Grpc_Client interfaces.AdminClient
}
func NewBookingHandler(BookingClient interfaces.AdminClient)*BookingHandler  {
	return &BookingHandler{
		Grpc_Client: BookingClient,
	}
}
func (b *BookingHandler)AddToBookings(c *fiber.Ctx)error  {
	patientID := c.Locals("user_id").(int)
    if patientID == 0{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "missing patient ID in headers",
        })
    }

    doctorIDStr := c.Query("doctor_id")
    doctorID, err := strconv.Atoi(doctorIDStr)
    if err != nil {
		errs := response.ClientResponse("Doctor id is given in the wrong format", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
    }

  
	err=b.Grpc_Client.AddToBookings(patientID,doctorID)
	if err!=nil{
		errs := response.ClientResponse("couldn't book please try again", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Successfully booked the doctor", nil, nil)
	return c.Status(fiber.StatusOK).JSON(success)

}
func (b *BookingHandler)CancelBookings(c *fiber.Ctx)error  {
	patientID := c.Locals("user_id").(int)
    if patientID == 0{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "missing patient ID in headers",
        })
    }

	bookingIDStr := c.Query("booking_id")
    bookingID, err := strconv.Atoi(bookingIDStr)
    if err != nil {
		errs := response.ClientResponse("booking id is given in the wrong format", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
    }

  
	err=b.Grpc_Client.CancelBookings(patientID,bookingID)
	if err!=nil{
		errs := response.ClientResponse("couldn't cancel booking, please try again", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Successfully cancelled booking", nil, nil)
	return c.Status(fiber.StatusOK).JSON(success)


}
