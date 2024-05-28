package interfaces

import models "healy-apigateway/pkg/utils"

type AdminClient interface {
	AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	MakePaymentRazorpay(patient_id int,doctodetails models.DoctorPaymentDetail) (models.Payment, string, error)
	VerifyPayment(payment_id int)error

}
