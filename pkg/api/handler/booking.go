package handler

import (
	"fmt"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/logging"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BookingHandler struct {
	Grpc_Client interfaces.AdminClient
}

func NewBookingHandler(BookingClient interfaces.AdminClient) *BookingHandler {
	return &BookingHandler{
		Grpc_Client: BookingClient,
	}
}
func (b *BookingHandler) AddToBookings(c *fiber.Ctx) error {
	patientID := c.Locals("user_id").(string)
	if patientID == "" {
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

	err = b.Grpc_Client.AddToBookings(patientID, doctorID)
	if err != nil {
		errs := response.ClientResponse("couldn't book please try again", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Successfully booked the doctor", nil, nil)
	return c.Status(fiber.StatusOK).JSON(success)

}
func (b *BookingHandler) CancelBookings(c *fiber.Ctx) error {
	patientID := c.Locals("user_id").(string)
	if patientID == "" {
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

	err = b.Grpc_Client.CancelBookings(patientID, bookingID)
	if err != nil {
		errs := response.ClientResponse("couldn't cancel booking, please try again", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Successfully cancelled booking", nil, nil)
	return c.Status(fiber.StatusOK).JSON(success)

}
func (b *BookingHandler) GetBookedPatients(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	patientdetails, err := b.Grpc_Client.GetPaidPatients(doctorId)
	if err != nil {
		errs := response.ClientResponse("couldn't fetch booked patients, please try again", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Successfully fetched booked patients", patientdetails, nil)
	return c.Status(fiber.StatusOK).JSON(success)
}
func (b *BookingHandler) CreatePrescription(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	patientID := c.Query("patient_id")
	bookingIDStr := c.Query("booking_id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert booking ID string to int", nil, err.Error()))
	}

	var prescription models.PrescriptionRequest
	if err := c.BodyParser(&prescription); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Details are not in correct format", nil, err.Error()))
	}

	prescription.DoctorID = doctorId
	prescription.PatientID = patientID
	prescription.BookingID = bookingID

	if err := validator.New().Struct(prescription); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Constraints not satisfied", nil, err.Error()))
	}

	createdPrescription, err := b.Grpc_Client.CreatePrescription(prescription)
	if err != nil {
		errorRes := response.ClientResponse("failed to create prescription", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("prescription created", createdPrescription, nil)
	return c.Status(200).JSON(successRes)

}
func (b *BookingHandler) SetDoctorAvailability(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errorRes := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	var setreq models.SetAvailability
	if err := c.BodyParser(&setreq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Details are not in correct format", nil, err.Error()))
	}
	res, err := b.Grpc_Client.SetDoctorAvailability(setreq, doctorId)
	if err != nil {
		errorRes := response.ClientResponse("failed to set availability", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("slot availability created created", res, nil)
	return c.Status(200).JSON(successRes)

}
func (b *BookingHandler) GetDoctorSlotAvailability(c *fiber.Ctx) error {
	doctorId, err := strconv.Atoi(c.Query("doctor_id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert doctor ID string to int", nil, err.Error()))
	}
	date := c.Query("date")
	res, err := b.Grpc_Client.GetDoctorAvailability(doctorId, date)
	if err != nil {
		errorRes := response.ClientResponse("failed to get doctor's availability", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("listed doctor's availabile slots", res, nil)
	return c.Status(200).JSON(successRes)
}
func (b *BookingHandler) BookSlot(c *fiber.Ctx) error {
	patientid := c.Locals("user_id").(string)
	bookingIDStr := c.Query("booking_id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert doctor ID string to int", nil, err.Error()))
	}
	slotIDStr := c.Query("slot_id")
	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert doctor ID string to int", nil, err.Error()))
	}
	err = b.Grpc_Client.BookSlot(patientid, bookingID, slotID)
	if err != nil {
		errorRes := response.ClientResponse("couldn't book the slot", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("slot booked for the bookingid", nil, nil)
	return c.Status(200).JSON(successRes)
}
func (b *BookingHandler) BookDoctor(c *fiber.Ctx) error {
	slotIdStr := c.Query("slot_id")
	slotID, err := strconv.Atoi(slotIdStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert slot ID string to int", nil, err.Error()))
	}

	patientId := c.Query("patient_id")
	bookingdetails, razorId, err := b.Grpc_Client.BookDoctor(patientId, slotID)
	if err != nil {
		errorRes := response.ClientResponse("Couldn't book Doctor", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	fmt.Println(bookingdetails, "paymentdetails")
	fmt.Println(razorId, "razorid")
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"final_price": bookingdetails.Fees * 100,
		"razor_id":    razorId,
		"user_id":     bookingdetails.PatientId,
		"order_id":    bookingdetails.BookingId,
		"user_email":  bookingdetails.DoctorEmail,
		"total":       int(bookingdetails.Fees),
	})
}

func (b *BookingHandler) VerifyandCalenderCreation(c *fiber.Ctx) error {
	logger := logging.Logger().WithField("function", "VerifyandCalenderCreation")
	logger.Info("Payment success endpoint hit")

	bookingIdStr := c.Query("booking_id")
	paymentId := c.Query("payment_id")
	razorId := c.Query("razor_id")

	logger = logger.WithFields(logrus.Fields{
		"booking_id": bookingIdStr,
		"payment_id": paymentId,
		"razor_id":   razorId,
	})
	logger.Info("Received payment verification request")

	bookingId, err := strconv.Atoi(bookingIdStr)
	if err != nil {
		logger.WithError(err).Error("Failed to convert booking ID to int")
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Invalid booking ID", nil, "Booking ID must be a valid integer"))
	}

	logger.Info("Calling gRPC VerifyandCalenderCreation")
	err = b.Grpc_Client.VerifyandCalenderCreation(bookingId, paymentId, razorId)
	if err != nil {
		logger.WithError(err).Error("gRPC VerifyandCalenderCreation failed")
		fmt.Println("gRPC error:", err.Error()) // Print the full error message

		// Determine if this is a client error or a server error
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "invalid") {
			return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Verification failed", nil, err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Internal server error", nil, "An unexpected error occurred"))
	}

	logger.Info("Slot booked successfully")
	return c.Status(http.StatusOK).JSON(response.ClientResponse("Slot booked on the calendar", nil, nil))
}
