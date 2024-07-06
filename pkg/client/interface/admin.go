package interfaces

import models "healy-apigateway/pkg/utils"

type AdminClient interface {
	AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)

	AddToBookings(patientid string, doctorid int) error
	CancelBookings(patientid string, bookingid int) error
	MakePaymentRazorpay(bookingid int) (models.CombinedBookingDetails, string, error)
	VerifyPayment(booking_id int) error

	GetPaidPatients(doctor_id int) ([]models.Patient, error)
	CreatePrescription(prescription models.PrescriptionRequest) (models.CreatedPrescription, error)

	SetDoctorAvailability(setreq models.SetAvailability, doctorId int) (string, error)
	GetDoctorAvailability(doctorid int, date string) ([]models.GetAvailability, error)
	BookSlot(patientid string, bookingid int, slotid int) error

	BookDoctor(patientid string, slotid int) (models.CombinedBookingDetails, string, error)
	VerifyandCalenderCreation(bookingId int, paymentid, razorid string) error
}
