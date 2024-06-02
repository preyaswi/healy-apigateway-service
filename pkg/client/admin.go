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
func (ad *adminClient)AddToBookings(patientid,doctorid int)error  {
	_,err:=ad.Client.AddTobookings(context.Background(),&pb.Bookingreq{
		PatientId: int32(patientid),
		DoctorId: int32(doctorid),
	})
	if err!=nil{
		return err
	}
	return nil
}
func (ad *adminClient)CancelBookings(patientid,bookingid int)error {
	_,err:=ad.Client.CancelBookings(context.Background(),&pb.Canbookingreq{PatientId: int32(patientid),BookingId: int32(bookingid)})
	if err!=nil{
		return err
	}
	return nil
}
func (ad *adminClient)MakePaymentRazorpay(bookingid int) (models.CombinedBookingDetails, string, error){
	res, err := ad.Client.MakePaymentRazorpay(context.Background(),&pb.PaymentReq{BookingId: uint32(bookingid)})
	if err != nil {
		return models.CombinedBookingDetails{}, "", err
	}
	paymentdetail:=models.CombinedBookingDetails{
		BookingId: uint(res.PaymentDetails.BookingId),
		PatientId: uint(res.PaymentDetails.PatientId),
		DoctorId: uint(res.PaymentDetails.DoctorId),
		DoctorName: res.PaymentDetails.DoctorName,
		DoctorEmail: res.PaymentDetails.DoctorEmail,
		Fees: res.PaymentDetails.Fees,
		PaymentStatus: res.PaymentDetails.PaymentStatus,
	}
	razor_id:=res.Razorid
	return paymentdetail,razor_id,nil

}
func (ad *adminClient)VerifyPayment(payment_id int) error {
	_,err:=ad.Client.VerifyPayment(context.Background(),&pb.Verifyreq{PaymentId: uint32(payment_id)})
	if err!=nil{
		return err
	}
	return nil
}
