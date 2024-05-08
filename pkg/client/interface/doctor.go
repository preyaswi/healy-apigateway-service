package interfaces

import models "healy-apigateway/pkg/utils"

type DoctorClient interface {
	DoctorSignUp(models.DoctorSignUp) (models.DoctorSignUpResponse, error)
	DoctorLogin(models.DoctorLogin) (models.DoctorSignUpResponse, error)
}
