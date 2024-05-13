package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	pb "healy-apigateway/pkg/pb/doctor"
	models "healy-apigateway/pkg/utils"

	"google.golang.org/grpc"
)

type doctorClient struct {
	Client pb.DoctorClient
}

func NewDoctorClient(cfg config.Config) interfaces.DoctorClient {
	grpcConnection, err := grpc.Dial(cfg.DoctorSvc, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	grpcClient := pb.NewDoctorClient(grpcConnection)
	return &doctorClient{
		Client: grpcClient,
	}
}
func (d *doctorClient) DoctorSignUp(signup models.DoctorSignUp) (models.DoctorSignUpResponse, error) {
	signupRes, err := d.Client.DoctorSignUp(context.Background(), &pb.DoctorSignUpRequest{
		FullName:          signup.FullName,
		Email:             signup.Email,
		PhoneNumber:       signup.PhoneNumber,
		Password:          signup.Password,
		ConfirmPassword:   signup.ConfirmPassword,
		Specialization:    signup.Specialization,
		YearsOfExperience: signup.YearsOfExperience,
		LicenseNumber:     signup.LicenseNumber,
	})
	if err != nil {
		return models.DoctorSignUpResponse{}, err
	}
	signupdetail := models.DoctorDetail{
		Id:                uint(signupRes.DoctorDetail.Id),
		FullName:          signupRes.DoctorDetail.FullName,
		Email:             signupRes.DoctorDetail.Email,
		PhoneNumber:       signupRes.DoctorDetail.PhoneNumber,
		Specialization:    signupRes.DoctorDetail.Specialization,
		YearsOfExperience: signupRes.DoctorDetail.YearsOfExperience,
		LicenseNumber:     signupRes.DoctorDetail.LicenseNumber,
	}
	return models.DoctorSignUpResponse{
		DoctorDetail: signupdetail,
		AccessToken:  signupRes.AccessToken,
		RefreshToken: signupRes.RefreshToken,
	}, nil
}
func (d *doctorClient) DoctorLogin(login models.DoctorLogin) (models.DoctorSignUpResponse, error) {
	loginRes, err := d.Client.DoctorLogin(context.Background(), &pb.DoctorLoginRequest{
		Email:    login.Email,
		Password: login.Password,
	})
	if err != nil {
		return models.DoctorSignUpResponse{}, err
	}
	loginDetail := models.DoctorDetail{
		Id:                uint(loginRes.DoctorDetail.Id),
		FullName:          loginRes.DoctorDetail.FullName,
		Email:             loginRes.DoctorDetail.Email,
		PhoneNumber:       loginRes.DoctorDetail.PhoneNumber,
		Specialization:    loginRes.DoctorDetail.Specialization,
		YearsOfExperience: loginRes.DoctorDetail.YearsOfExperience,
		LicenseNumber:     loginRes.DoctorDetail.LicenseNumber,
	}
	return models.DoctorSignUpResponse{
		DoctorDetail: loginDetail,
		AccessToken:  loginRes.AccessToken,
		RefreshToken: loginRes.RefreshToken,
	}, nil

}
func (d *doctorClient)DoctorsDetails() ([]models.DoctorsDetails,error) {
	response,err:=d.Client.DoctorsDetail(context.Background(),&pb.Doreq{})
	if err!=nil{
		return []models.DoctorsDetails{},err
	}
	doctorsDetails := make([]models.DoctorsDetails, len(response.DoctorsDetailr))
    for i, detail := range response.DoctorsDetailr {
        doctorDetail := models.DoctorDetail{
            Id:                uint(detail.Id),
            FullName:          detail.FullName,
            Email:             detail.Email,
            PhoneNumber:       detail.PhoneNumber,
            Specialization:    detail.Specialization,
            YearsOfExperience: detail.YearsOfExperience,
            LicenseNumber:     detail.LicenseNumber,
        }
        doctorsDetails[i] = models.DoctorsDetails{
            DoctorDetail: doctorDetail,
            Rating:       detail.Rating,
        }
    }

    return doctorsDetails, nil
}
func (d *doctorClient)IndividualDoctor(doctorId string)(models.IndDoctorDetail,error)  {
	doctor,err:=d.Client.IndividualDoctor(context.Background(),&pb.Doid{DoctorId: doctorId})
	if err!=nil{
		return models.IndDoctorDetail{},err
	}
	return models.IndDoctorDetail{
		Id: uint(doctor.Id),
		FullName: doctor.FullName,
		Email: doctor.Email,
		PhoneNumber: doctor.PhoneNumber,
		Specialization: doctor.Specialization,
		YearsOfExperience: doctor.YearsOfExperience,
		LicenseNumber: doctor.LicenseNumber,
		Rating: doctor.Rating,
	},nil
}
func (d *doctorClient)DoctorProfile(id int)(models.IndDoctorDetail,error)  {
	res,err:=d.Client.DoctorProfile(context.Background(),&pb.DoId{Id: uint64(id)})
	if err!=nil{
		return models.IndDoctorDetail{},err
	}
	return models.IndDoctorDetail{
		Id: uint(id),
		FullName: res.FullName,
		Email: res.Email,
		PhoneNumber: res.PhoneNumber,
		Specialization: res.Specialization,
		YearsOfExperience: res.YearsOfExperience,
		LicenseNumber: res.LicenseNumber,
	},nil
}