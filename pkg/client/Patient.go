package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	models "healy-apigateway/pkg/utils"
	pb "healy-apigateway/pkg/pb/patient"

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
		Fullname:        patient.Fullname,
		Email:           patient.Email,
		Password:        patient.Password,
		Confirmpassword: patient.Confirmpassword,
		Gender:          patient.Gender,
		Contactnumber:   patient.Contactnumber,
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
	fmt.Println(patient, "patient details")
	res, err := p.Client.PatientLogin(context.Background(), &pb.PatientLoginRequest{
		Email:    patient.Email,
		Password: patient.Password,
	})
	if err != nil {
		return models.TokenPatient{}, err
	}
	fmt.Println(res, "REsponse")
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
