package interfaces

import models "healy-apigateway/pkg/utils"

type AdminClient interface {
	AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)

	AddToBookings(patientid,doctorid int)error
	CancelBookings(patientid,bookingid int)error
	MakePaymentRazorpay(bookingid int) (models.CombinedBookingDetails, string, error)
	VerifyPayment(payment_id int)error

}
