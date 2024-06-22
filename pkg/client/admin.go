package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	models "healy-apigateway/pkg/utils"

	pb "healy-apigateway/pkg/pb/admin"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnection, err := grpc.Dial(cfg.AdminSvc, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}

}
func (ad *adminClient) AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminSignup(context.Background(), &pb.AdminSignupRequest{
		Firstname: admindeatils.Firstname,
		Lastname:  admindeatils.Lastname,
		Email:     admindeatils.Email,
		Password:  admindeatils.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}

func (ad *adminClient) AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
	})

	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}
func (ad *adminClient) AddToBookings(patientid string, doctorid int) error {
	_, err := ad.Client.AddTobookings(context.Background(), &pb.Bookingreq{
		PatientId: patientid,
		DoctorId:  int32(doctorid),
	})
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminClient) CancelBookings(patientid string, bookingid int) error {
	_, err := ad.Client.CancelBookings(context.Background(), &pb.Canbookingreq{PatientId:patientid, BookingId: int32(bookingid)})
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminClient) MakePaymentRazorpay(bookingid int) (models.CombinedBookingDetails, string, error) {
	res, err := ad.Client.MakePaymentRazorpay(context.Background(), &pb.PaymentReq{BookingId: uint32(bookingid)})
	if err != nil {
		return models.CombinedBookingDetails{}, "", err
	}
	paymentdetail := models.CombinedBookingDetails{
		BookingId:     uint(res.PaymentDetails.BookingId),
		PatientId:     uint(res.PaymentDetails.PatientId),
		DoctorId:      uint(res.PaymentDetails.DoctorId),
		DoctorName:    res.PaymentDetails.DoctorName,
		DoctorEmail:   res.PaymentDetails.DoctorEmail,
		Fees:          res.PaymentDetails.Fees,
		PaymentStatus: res.PaymentDetails.PaymentStatus,
	}
	razor_id := res.Razorid
	return paymentdetail, razor_id, nil

}
func (ad *adminClient) VerifyPayment(booking_id int) error {
	_, err := ad.Client.VerifyPayment(context.Background(), &pb.PaymentReq{BookingId: uint32(booking_id)})
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminClient) GetPaidPatients(doctor_id int) ([]models.Patient, error) {
	res, err := ad.Client.GetPaidPatients(context.Background(), &pb.GetPaidPatientsRequest{DoctorId: int32(doctor_id)})
	if err != nil {
		return []models.Patient{}, err
	}
	patientsDetails := make([]models.Patient, len(res.BookedPatients))
	for i, bookedPatient := range res.BookedPatients {
		patientsDetails[i] = models.Patient{
			BookingId: uint(bookedPatient.BookingId),
			PaymentStatus: bookedPatient.PaymentStatus,
			PatientId: uint(bookedPatient.PatientDetail.Id),
			Fullname:  bookedPatient.PatientDetail.Fullname,
			Email:     bookedPatient.PatientDetail.Email,
			Gender:    bookedPatient.PatientDetail.Gender,
			Contactnumber: bookedPatient.PatientDetail.Contactnumber,
		}
	}
	return patientsDetails, nil
}
func (ad *adminClient) CreatePrescription(prescription models.PrescriptionRequest) (models.CreatedPrescription, error) {
	res, err := ad.Client.CreatePrescription(context.Background(), &pb.CreatePrescriptionRequest{
		BookingId: uint32(prescription.BookingID),
		DoctorId: uint32(prescription.DoctorID),
		PatientId: prescription.PatientID,
		Medicine:   prescription.Medicine,
		Dosage:     prescription.Dosage,
		Notes:      prescription.Notes,
	})
	if err != nil {
		return models.CreatedPrescription{}, err
	}
	return models.CreatedPrescription{
		Id: int(res.Id),
		DoctorID:   int(res.DoctorId),
		PatientID:  res.PatientId,
		BookingID: int(res.BookingId),
		Medicine:   res.Medicine,
		Dosage:     res.Dosage,
		Notes:      res.Notes,
	}, nil

}
