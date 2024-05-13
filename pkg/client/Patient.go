package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	pb "healy-apigateway/pkg/pb/patient"
	models "healy-apigateway/pkg/utils"

	"google.golang.org/grpc"
)

type patientClient struct {
	Client pb.PatientClient
}

func NewPatientClient(cfg config.Config) interfaces.PatientClient {
	grpcConnection, err := grpc.Dial(cfg.PatientSvc, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	grpcClient := pb.NewPatientClient(grpcConnection)
	return &patientClient{
		Client: grpcClient,
	}
}
func (p *patientClient) PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error) {

	res, err := p.Client.PatientSignUp(context.Background(), &pb.PatientSignUpRequest{
		Fullname:        patient.FullName,
		Email:           patient.Email,
		Password:        patient.Password,
		Confirmpassword: patient.ConfirmPassword,
		Gender:          patient.Gender,
		Contactnumber:   patient.ContactNumber,
	})
	if err != nil {
		return models.TokenPatient{}, err
	}
	patientDetails := models.SignupdetailResponse{
		Id:            uint(res.PatientDetails.Id),
		Fullname:      res.PatientDetails.Fullname,
		Email:         res.PatientDetails.Email,
		Gender:        res.PatientDetails.Gender,
		Contactnumber: res.PatientDetails.Contactnumber,
	}
	return models.TokenPatient{
		Patient:      patientDetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil

}
func (p *patientClient) PatientLogin(patient models.PatientLogin) (models.TokenPatient, error) {
	res, err := p.Client.PatientLogin(context.Background(), &pb.PatientLoginRequest{
		Email:    patient.Email,
		Password: patient.Password,
	})
	if err != nil {
		return models.TokenPatient{}, err
	}
	patientdetails := models.SignupdetailResponse{
		Id:            uint(res.PatientDetails.Id),
		Fullname:      res.PatientDetails.Fullname,
		Email:         res.PatientDetails.Email,
		Gender:        res.PatientDetails.Gender,
		Contactnumber: res.PatientDetails.Contactnumber,
	}
	return models.TokenPatient{
		Patient:      patientdetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
func (p *patientClient) PatientDetails(user_id int) (models.SignupdetailResponse, error) {
	res, err := p.Client.IndPatientDetails(context.Background(), &pb.Idreq{UserId: uint64(user_id)})
	if err != nil {
		return models.SignupdetailResponse{}, err
	}
	fmt.Println(res.Id, "resid")
	return models.SignupdetailResponse{
		Id:            uint(res.Id),
		Fullname:      res.Fullname,
		Email:         res.Email,
		Gender:        res.Gender,
		Contactnumber: res.Contactnumber,
	}, nil
}
func (p *patientClient) UpdatePatientDetails(pa models.PatientDetails, patient_id int) (models.PatientDetails, error) {
	patient := &pb.InPatientDetails{
		Fullname:      pa.Fullname,
		Email:         pa.Email,
		Gender:        pa.Gender,
		Contactnumber: pa.Contactnumber,
	}
	res, err := p.Client.UpdatePatientDetails(context.Background(), &pb.UpdateRequest{
		PatientId:        uint64(patient_id),
		InPatientDetails: patient,
	})
	if err != nil {
		return models.PatientDetails{}, err
	}
	return models.PatientDetails{
		Fullname:      res.Fullname,
		Email:         res.Email,
		Gender:        res.Gender,
		Contactnumber: res.Contactnumber,
	}, nil
}
func (p *patientClient) UpdatePassword(ctx context.Context,userId int, body models.UpdatePassword)error {

	res,err:=p.Client.UpdatePassword(context.Background(),&pb.UpdatePasswordRequest{
		PatientId: uint64(userId),
		OldPassword: body.OldPassword,
		NewPassword: body.NewPassword,
		ConfirmNewPassword: body.ConfirmNewPassword,
	})
	if err!=nil{
		return err
	}
	fmt.Println(res)
	return nil
}
